package provider

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/nuonco/nuon-go"
	"github.com/nuonco/nuon-go/models"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ContainerImageComponentResource{}
var _ resource.ResourceWithImportState = &ContainerImageComponentResource{}

func NewContainerImageComponentResource() resource.Resource {
	return &ContainerImageComponentResource{}
}

// ContainerImageComponentResource defines the resource implementation.
type ContainerImageComponentResource struct {
	baseResource
}

type AwsEcr struct {
	Region     types.String `tfsdk:"region"`
	Tag        types.String `tfsdk:"tag"`
	ImageURL   types.String `tfsdk:"image_url"`
	IAMRoleARN types.String `tfsdk:"iam_role_arn"`
}

type Public struct {
	ImageURL types.String `tfsdk:"image_url"`
	Tag      types.String `tfsdk:"tag"`
}

// ContainerImageComponentResourceModel describes the resource data model.
type ContainerImageComponentResourceModel struct {
	ID types.String `tfsdk:"id"`

	Name         types.String `tfsdk:"name"`
	VarName      types.String `tfsdk:"var_name"`
	Dependencies types.List   `tfsdk:"dependencies"`
	AppID        types.String `tfsdk:"app_id"`

	AwsEcr *AwsEcr `tfsdk:"aws_ecr"`
	Public *Public `tfsdk:"public"`

	EnvVar []EnvVar `tfsdk:"env_var"`
}

func (r *ContainerImageComponentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_container_image_component"
}

func (r *ContainerImageComponentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Use a Docker, ECR or OCI compatible image as a component.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The unique ID of the component.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The human-readable name of the component.",
				Optional:    false,
				Required:    true,
			},
			"var_name": schema.StringAttribute{
				Description: "The optional var name to be used when referencing this component.",
				Optional:    true,
				Required:    false,
			},
			"app_id": schema.StringAttribute{
				Description: "The unique ID of the app this component belongs too.",
				Optional:    false,
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"dependencies": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Component dependencies",
				Optional:    true,
				Required:    false,
			},
			"public": schema.SingleNestedAttribute{
				Description: "Use a publically-accessible image.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"image_url": schema.StringAttribute{
						Description: "The full image URL or docker hub alias (e.g. kennethreitz/httpbin).",
						Required:    true,
					},
					"tag": schema.StringAttribute{
						Description: "The image tag.",
						Required:    true,
					},
				},
			},
			"aws_ecr": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Use an image stored in AWS ECR.",
				Attributes: map[string]schema.Attribute{
					"image_url": schema.StringAttribute{
						Description: "The full URL of your ECR repo.",
						Required:    true,
					},
					"tag": schema.StringAttribute{
						Description: "The image tag.",
						Required:    true,
					},
					"region": schema.StringAttribute{
						Description: "The region of your ECR repo.",
						Required:    true,
					},
					"iam_role_arn": schema.StringAttribute{
						Description: "The AWS IAM role needed to access your ECR repo.",
						Required:    true,
					},
				},
			},
		},
		Blocks: map[string]schema.Block{
			"env_var": envVarSharedBlock(),
		},
	}
}

func (r *ContainerImageComponentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ContainerImageComponentResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "creating component")

	dependencies := make([]string, 0)
	resp.Diagnostics.Append(data.Dependencies.ElementsAs(ctx, &dependencies, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	compResp, err := r.restClient.CreateComponent(ctx, data.AppID.ValueString(), &models.ServiceCreateComponentRequest{
		Name:         data.Name.ValueStringPointer(),
		VarName:      data.VarName.ValueString(),
		Dependencies: dependencies,
	})
	if err != nil {
		writeDiagnosticsErr(ctx, &resp.Diagnostics, err, "create component")
		return
	}
	tflog.Trace(ctx, "got ID -- "+compResp.ID)
	data.ID = types.StringValue(compResp.ID)
	data.Name = types.StringValue(compResp.Name)
	//data.VarName = types.StringValue(compResp.VarName)

	configRequest := &models.ServiceCreateExternalImageComponentConfigRequest{}
	if data.AwsEcr != nil {
		configRequest.ImageURL = data.AwsEcr.ImageURL.ValueStringPointer()
		configRequest.Tag = data.AwsEcr.Tag.ValueStringPointer()
		configRequest.AwsEcrImageConfig = &models.ServiceAwsECRImageConfigRequest{
			AwsRegion:  data.AwsEcr.Region.ValueString(),
			IamRoleArn: data.AwsEcr.IAMRoleARN.ValueString(),
		}
	} else if data.Public != nil {
		configRequest.ImageURL = data.Public.ImageURL.ValueStringPointer()
		configRequest.Tag = data.Public.Tag.ValueStringPointer()
	}

	_, err = r.restClient.CreateExternalImageComponentConfig(ctx, compResp.ID, configRequest)
	if err != nil {
		// attempt to cleanup component, that is in broken state and has no config
		_, cleanupErr := r.restClient.DeleteComponent(ctx, compResp.ID)
		if cleanupErr != nil {
			tflog.Trace(ctx, fmt.Sprintf("unable to cleanup component: %s", cleanupErr))
		}

		writeDiagnosticsErr(ctx, &resp.Diagnostics, err, "create component config")
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "successfully created component")
}

func (r *ContainerImageComponentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ContainerImageComponentResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	compResp, err := r.restClient.GetComponent(ctx, data.ID.ValueString())
	if nuon.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		writeDiagnosticsErr(ctx, &resp.Diagnostics, err, "get component")
		return
	}
	data.Name = types.StringValue(compResp.Name)
	data.VarName = types.StringValue(compResp.VarName)
	data.AppID = types.StringValue(compResp.AppID)

	configResp, err := r.restClient.GetComponentLatestConfig(ctx, data.ID.ValueString())
	if err != nil {
		writeDiagnosticsErr(ctx, &resp.Diagnostics, err, "get component config")
		return
	}
	if configResp.ExternalImage == nil {
		writeDiagnosticsErr(ctx, &resp.Diagnostics, errors.New("did not get external image config"), "get component config")
		return
	}
	externalImage := configResp.ExternalImage
	if externalImage.AwsEcrImageConfig != nil {
		data.AwsEcr = &AwsEcr{
			ImageURL:   types.StringValue(externalImage.ImageURL),
			Tag:        types.StringValue(externalImage.Tag),
			Region:     types.StringValue(externalImage.AwsEcrImageConfig.AwsRegion),
			IAMRoleARN: types.StringValue(externalImage.AwsEcrImageConfig.IamRoleArn),
		}
	} else {
		data.Public = &Public{
			ImageURL: types.StringValue(externalImage.ImageURL),
			Tag:      types.StringValue(externalImage.Tag),
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "successfully read component")
}

func (r *ContainerImageComponentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ContainerImageComponentResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	deleted, err := r.restClient.DeleteComponent(ctx, data.ID.ValueString())
	if err != nil {
		writeDiagnosticsErr(ctx, &resp.Diagnostics, err, "delete component")
		return
	}
	if !deleted {
		writeDiagnosticsErr(ctx, &resp.Diagnostics, err, "delete component")
		return
	}

	stateConf := &retry.StateChangeConf{
		Pending: []string{statusActive, statusDeleteQueued, statusDeprovisioning, statusTemporarilyUnavailable},
		Target:  []string{statusNotFound},
		Refresh: func() (interface{}, string, error) {
			tflog.Trace(ctx, "refreshing component status")
			cmp, err := r.restClient.GetComponent(ctx, data.ID.ValueString())
			if err == nil {
				return cmp.Status, cmp.Status, nil
			}
			if nuon.IsNotFound(err) {
				return "", statusNotFound, nil
			}

			logErr(ctx, err, "delete component")
			return statusTemporarilyUnavailable, statusTemporarilyUnavailable, nil
		},
		Timeout:    time.Minute * 20,
		Delay:      time.Second * 10,
		MinTimeout: 3 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		writeDiagnosticsErr(ctx, &resp.Diagnostics, err, "unable to delete component")
		return
	}

	tflog.Trace(ctx, "successfully deleted component")
}

func (r *ContainerImageComponentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ContainerImageComponentResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "updating component "+data.ID.ValueString())

	dependencies := make([]string, 0)
	resp.Diagnostics.Append(data.Dependencies.ElementsAs(ctx, &dependencies, false)...)
	if resp.Diagnostics.HasError() {
		return
	}
	compResp, err := r.restClient.UpdateComponent(ctx, data.ID.ValueString(), &models.ServiceUpdateComponentRequest{
		Name:         data.Name.ValueStringPointer(),
		VarName:      data.VarName.ValueString(),
		Dependencies: dependencies,
	})
	if err != nil {
		writeDiagnosticsErr(ctx, &resp.Diagnostics, err, "update component")
		return
	}
	data.Name = types.StringValue(compResp.Name)
	//data.VarName = types.StringValue(compResp.VarName)

	configRequest := &models.ServiceCreateExternalImageComponentConfigRequest{}
	if data.AwsEcr != nil {
		configRequest.ImageURL = data.AwsEcr.ImageURL.ValueStringPointer()
		configRequest.Tag = data.AwsEcr.Tag.ValueStringPointer()
		configRequest.AwsEcrImageConfig = &models.ServiceAwsECRImageConfigRequest{
			AwsRegion:  data.AwsEcr.Region.ValueString(),
			IamRoleArn: data.AwsEcr.IAMRoleARN.ValueString(),
		}
	} else {
		configRequest.ImageURL = data.Public.ImageURL.ValueStringPointer()
		configRequest.Tag = data.Public.Tag.ValueStringPointer()
	}
	_, err = r.restClient.CreateExternalImageComponentConfig(ctx, compResp.ID, configRequest)
	if err != nil {
		writeDiagnosticsErr(ctx, &resp.Diagnostics, err, "create component config")
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "successfully updated component")
}

func (r *ContainerImageComponentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
