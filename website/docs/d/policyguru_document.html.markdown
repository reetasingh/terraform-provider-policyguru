---
subcategory: "IAM"
layout: "policyguru"
page_title: "Policy Guru: policyguru_document"
description: |-
  Generates an IAM policy document in JSON format
---

# Data Source: policyguru_document

Generates an IAM policy document in JSON format.

This is a data source which can be used to construct a JSON representation of
an IAM policy document, for use with resources which expect policy documents,
such as the `aws_iam_policy` resource.

-> For more information about building AWS IAM policy with Terraform, see the [AWS IAM Policy Resource](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_policy).

```hcl
data "policyguru_document" "example" {
  actions_for_resources_at_access_level {
    write = list("arn:aws:kms:us-east-1:123456789012:key/aaaa-bbbb-cccc")
    read = list("arn:aws:s3:::mybucket")
  }

  actions_for_service_without_resource_constraint_support {
    list = list("s3")
    include_single_actions = ["ssm:GetParameter"]
  }

  overrides {
    skip_resource_constraints_for_actions = []
  }

  exclude_actions = list("kms:Decrypt*", "kms:Delete*", "kms:Schedule*")
}

resource "aws_iam_policy" "policy" {
  name        = "sample"
  path        = "/"
  description = "this uses policyguru document"
  policy      = data.policyguru_document.example.json
}
```

## Argument Reference

The following arguments are supported:

* `exclude_actions` (Optional) - A list of actions that should not be included in the resulting policy.
* `actions_for_resources_at_access_level` (Optional) - Provide Information about list of Amazon Resource Names (ARNs) that your role needs access to. `actions_for_resources_at_access_level` block is documented below.
* `actions_for_service_without_resource_constraint_support` (Optional) - Provide Information about AWS service actions that do not support resource ARN constraints. `actions_for_service_without_resource_constraint_support` block is documented below.
* `overrides` (Optional) - handle exceptions/overrides for policies. `overrides` block is documented below.


`actions_for_resources_at_access_level` supports the following

* `read` (Optional) - Provide a list of Amazon Resource Names (ARNs) that your role needs `READ` access to.
* `write` (Optional) - Provide a list of Amazon Resource Names (ARNs) that your role needs `WRITE` access to.
* `tagging` (Optional) - Provide a list of Amazon Resource Names (ARNs) that your role needs `TAGGING` access to.
* `list` (Optional) - Provide a list of Amazon Resource Names (ARNs) that your role needs `LIST` access to.
* `permissions_management` (Optional) - Provide a list of Amazon Resource Names (ARNs) that your role needs `PERMISSIONS MANAGEMENT` access to.


`actions_for_service_without_resource_constraint_support` supports the following

* `read` (Optional) - Provide a List of service prefix to generate a list of AWS service actions that (1) are at the `READ` access level and (2) do not support resource constraints.
* `write` (Optional) - Provide a List of service prefix to generate a list of AWS service actions that (1) are at the `WRITE` access level and (2) do not support resource constraints.
* `tagging` (Optional) - Provide a List of service prefix to generate a list of AWS service actions that (1) are at the `TAGGING` access level and (2) do not support resource constraints.
* `list` (Optional) - Provide a List of service prefix to generate a list of AWS service actions that (1) are at the `LIST` access level and (2) do not support resource constraints.
* `permissions_management` (Optional) - Provide a List of service prefix to generate a list of AWS service actions that (1) are at the `PERMISSIONS MANAGEMENT` access level and (2) do not support resource constraints.
* `include_single_actions` (Optional) - Provide a List of individual actions that do not support resource constraints. For example, s3:ListAllMyBuckets


`overrides` supports the following

* `skip_resource_constraints_for_actions` (Optional) - Provide a list of Actions that do not support resource constraints. For more details on how skip_resource_constraints_for_actions works - https://policy-sentry.readthedocs.io/en/latest/writing-policies/skipping-resource-constraints/

## Attributes Reference

The following attribute is exported:

* `json` - policy document in a JSON formatted string.