---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nuon_app_installer Resource - terraform-provider-nuon"
subcategory: ""
description: |-
  A public installer page for a nuon app.
---

# nuon_app_installer (Resource)

A public installer page for a nuon app.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) The application ID.
- `community_url` (String) Community url to a slack or discord, etc.
- `description` (String) Short description of app.
- `documentation_url` (String) Documentation url
- `github_url` (String) GitHub url, link to application.
- `homepage_url` (String) Homepage url
- `logo_url` (String) Logo image to display on page.
- `name` (String) App name to render on install page.
- `post_install_markdown` (String) Markdown that will be shown to users after a successful install. Supports interpolation.
- `slug` (String) URL slug to access app from.

### Optional

- `demo_url` (String) Demo url to show

### Read-Only

- `id` (String) The unique ID of the app.
