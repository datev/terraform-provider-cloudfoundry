---
page_title: "cloudfoundry_org_role Data Source - terraform-provider-cloudfoundry"
subcategory: ""
description: |-
  Gets information on a Cloud Foundry org role with a given role ID.
---

# cloudfoundry_org_role (Data Source)

Gets information on a Cloud Foundry org role with a given role ID.

## Example Usage

```terraform
data "cloudfoundry_org_role" "my_role" {
  id = "233959b9-c1fe-428d-8b53-5c903e5bd66b"
}

output "role_object" {
  value = data.cloudfoundry_space_role.my_role
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The guid for the role

### Read-Only

- `created_at` (String) The date and time when the resource was created in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) format.
- `org` (String) The guid of the organization the role is assigned to
- `type` (String) Role type; see [Valid role types](https://v3-apidocs.cloudfoundry.org/version/3.154.0/index.html#valid-role-types)
- `updated_at` (String) The date and time when the resource was updated in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) format.
- `user` (String) The guid of the cloudfoundry user the role is assigned to