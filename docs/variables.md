---
page_title: "Provider: Resource Variables"
description: |-
    Resource variables allow you to configure Nuon resources with outputs from other resources.
---

# Using Variables in Nuon Resources

## Overview

The components of your application can get custom configuration settings applied by using variables. Each deployed component instance can make use of a set of name/value variables to handle configuration settings that need to be set differently on each customer's instance. Nuon brings in a tree of information from the sources detailed below, which can then be accessed by the component using a templating syntax inspired by Go templates. This variable replacement templating syntax is used by several other infrastructure management tools.

You can set these variables as vars in your Terraform configuration.

**Example Helm chart component**
```
resource "nuon_app" "my_app" {
  name = "My App"
}

data "nuon_connected_repo" "my_repo" {
  name = "my-github-org/my-repo"
}

resource "nuon_helm_chart_component" "my_component" {
  name   = "My Component"
  app_id = nuon_app.my_app.id

  chart_name = "my-helm-chart"

  connected_repo = {
    directory = "example/chart"
    repo      = nuon_connected_repo.my_repo.name
    branch    = nuon_connected_repo.my_repo.default_branch
  }

  # manually set a variable
  value {
    name  = "some-var-name"
    value = "some-var-value"
  }

  # reference another component
  value {
    name  = "reference-to-other-component"
    value = "{{.nuon.components.some_other_component.outputs.s3_bucket_name}}"
  }

  # reference the install
  value {
    name  = "reference-to-install-attribute"
    value = "{{.nuon.install.public_domain}}"
  }
}
```

Example directly-set value would be

| Name | Value |
|-|-|
| `port` | `3000` |

Example dynamic variable reference would be

| Name | Value |
|-|-|
| `bucket_name` | `{{ .nuon.components.storage.bucket_name }}` |

**Note** After changing variables, you must trigger a new build and release of that component before your changes will take affect.

## Same Variable Value on Every Install

Settings common to all your installs can be set directly on the component and it will be the same across all installs.

Example: `cache_size` with value `4096`.

## Unique Variable Value For Each Installs

Settings that need to change for each specific install should be set as follows:

1. On the component using a template reference to a named secret
    - Example: `{{ .nuon.secrets.api_key }}`
1. For each install set a variable with the matching name
    - Example `api_key` with value `ak_123456`

Each deployed instance of this component will get a different value corresponding to these settings.

## Variable Data Sources

### Nuon Information

- **Nuon Organization ID** `{{ .nuon.org.id }}`
  * Unique identifier for the vendor organization
- **Nuon Application ID** `{{ .nuon.app.id }}`
  * Unique identifier for the vendor application
- **Nuon Install ID** `{{ .nuon.install.id }}`
- **Nuon Install Public Domain** `{{ .nuon.install.public_domain }}`
- **Nuon Install Internal Domain** `{{ .nuon.install.internal_domain }}`
- **Nuon Install Sandbox Type** `{{ .nuon.install.sandbox.type }}`
  * Example: `aws-eks`
- **Nuon Install Sandbox Version** `{{ .nuon.install.sandbox.version }}`
  * Example: `0.11.1`

### Nuon Install Sandbox Outputs

- `{{ .nuon.install.sandbox.outputs.cluster_arn }}`
  - ARN for the EKS Cluster running this application
- `{{ .nuon.install.sandbox.outputs.cluster_certificate_authority_data }}`
  - CA data for this EKS cluster, encoded as base64
- `{{ .nuon.install.sandbox.outputs.cluster_endpoint }}`
  - URL for the EKS cluster
- `{{ .nuon.install.sandbox.outputs.cluster_name }}`
  - Name of the EKS cluster. Equal to the Nuon Install ID.
- `{{ .nuon.install.sandbox.outputs.cluster_platform_version }}`
  - EKS version running. Example: `eks.7`
- `{{ .nuon.install.sandbox.outputs.cluster_status }}`
  - Typically `ACTIVE`
- `{{ .nuon.install.sandbox.outputs.ecr_registry_arn }}`
  - ARN for the ECR Registry for this application
- `{{ .nuon.install.sandbox.outputs.ecr_registry_id }}`
  - ID for the EKS Cluster running this application
- `{{ .nuon.install.sandbox.outputs.ecr_registry_url }}`
  - URL for the EKS Cluster running this application
- `{{ .nuon.install.sandbox.outputs.odr_iam_role_arn }}`
  - ARN for the AWS IAM Role used by the On-Demand Runner for deployments


### Output from Other Components in an Application

Some component types including terraform components will provide output values that will be required as input variables to other components of the application. These outputs are available grouped under the slug version of the component name.

- Template Syntax: `{{ .nuon.components.<component_name_slug>.<output_name> }}`
- Example: `{{ .nuon.components.rds_db.db_url }}`

### Component Image Data from Other Components in an Application

- `{{ .nuon.components.<component_name_slug>.image.tag }}`
  - Docker/OCI Image tag used for the component deployment
- `{{ .nuon.components.<component_name_slug>.image.repository }}`
  - ECR Repository containing this component's images
- `{{ .nuon.components.<component_name_slug>.image.registry }}`
  - ECR Registry containing this component's images
