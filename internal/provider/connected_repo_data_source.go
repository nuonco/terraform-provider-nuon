package provider

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &ConnectedRepoDataSource{}

func NewConnectedRepoDataSource() datasource.DataSource {
	return &ConnectedRepoDataSource{}
}

// ConnectedRepoDataSource defines the data source implementation.
type ConnectedRepoDataSource struct {
	baseDataSource
}

// ConnectedRepoDataSourceModel describes the data source data model.
type ConnectedRepoDataSourceModel struct {
	// inputs
	Name types.String `tfsdk:"name"`

	// computed
	DefaultBranch types.String `tfsdk:"default_branch"`
	FullName      types.String `tfsdk:"full_name"`
	Repo          types.String `tfsdk:"repo"`
	Owner         types.String `tfsdk:"owner"`
	URL           types.String `tfsdk:"url"`
}

func (d *ConnectedRepoDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connected_repo"
}

func (d *ConnectedRepoDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Get a connected repo tied to your org.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description: "The name or URL of the connected repo",
				Required:    true,
			},
			"default_branch": schema.StringAttribute{
				Description: "The default branch of the repo.",
				Computed:    true,
			},
			"full_name": schema.StringAttribute{
				Description: "The full name of the repo.",
				Computed:    true,
			},
			"repo": schema.StringAttribute{
				Description: "The name of the repo.",
				Computed:    true,
			},
			"owner": schema.StringAttribute{
				Description: "The owner of the repo.",
				Computed:    true,
			},
			"url": schema.StringAttribute{
				Description: "The URL of the repo.",
				Computed:    true,
			},
		},
	}
}

func (d *ConnectedRepoDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ConnectedRepoDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "fetching connected repo")

	repos, err := d.restClient.GetAllVCSConnectedRepos(ctx)
	if err != nil {
		writeDiagnosticsErr(ctx, &resp.Diagnostics, err, "get connected repo")
		return
	}

	for _, repo := range repos {
		if repo.FullName == nil || *repo.FullName != data.Name.ValueString() {
			continue
		}

		data.DefaultBranch = types.StringValue(*repo.DefaultBranch)
		data.FullName = types.StringValue(*repo.FullName)
		data.Repo = types.StringValue(*repo.Name)
		data.Owner = types.StringValue(*repo.UserName)
		data.URL = types.StringValue(*repo.CloneURL)

		// Save data into Terraform state
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	writeDiagnosticsErr(ctx, &resp.Diagnostics, errors.New("repo not found"), "get connected repo")
}
