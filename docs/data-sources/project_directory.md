---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vercel_project_directory Data Source - terraform-provider-vercel"
subcategory: ""
description: |-
  Provides information about files within a directory on disk.
  This will recursively read files, providing metadata for use with a vercel_deployment.
  -> If you want to prevent files from being included, this can be done with a vercelignore file https://vercel.com/guides/prevent-uploading-sourcepaths-with-vercelignore.
---

# vercel_project_directory (Data Source)

Provides information about files within a directory on disk.

This will recursively read files, providing metadata for use with a `vercel_deployment`.

-> If you want to prevent files from being included, this can be done with a [vercelignore file](https://vercel.com/guides/prevent-uploading-sourcepaths-with-vercelignore).

## Example Usage

```terraform
data "vercel_project_directory" "example" {
  path = "../ui"
}

data "vercel_project" "example" {
  name = "my-project"
}

resource "vercel_deployment" "example" {
  project_id = data.vercel_project.example.id
  files      = data.vercel_project_directory.example.files
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **path** (String) The path to the directory on your filesystem. Note that the path is relative to the root of the terraform files.

### Read-Only

- **files** (Map of String) A map of filename to metadata about the file. The metadata contains the file size and hash, and allows a deployment to be created if the file changes.
- **id** (String) The ID of this resource.


