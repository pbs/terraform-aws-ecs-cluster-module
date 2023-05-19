# PBS TF ECS Cluster Module

## Installation

### Using the Repo Source

Use this URL for the source of the module. See the usage examples below for more details.

```hcl
github.com/pbs/terraform-aws-ecs-cluster-module?ref=x.y.z
```

### Alternative Installation Methods

More information can be found on these install methods and more in [the documentation here](./docs/general/install).

## Usage

This module creates an ECS cluster.

Normally, the ECS cluster that is created does not have any EC2 instances backing it. This is ideal for scenarios where only Fargate tasks will be provisioned within the cluster. To enable EC2 backing, use the `ec2_backed` parameter to provision an autoscaled fleet of EC2 instances to support your cluster.

The instances in this cluster will also automatically rotate out instances to ensure that all tasks running in the cluster can handle unexpected hardware failure. To control the rate at which instances are rotated out, utilize the `max_instance_lifetime` parameter.

Integrate this module like so:

```hcl
module "cluster" {
  source = "github.com/pbs/terraform-aws-ecs-cluster-module?ref=x.y.z"

  # Tagging Parameters
  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo

  # Optional Parameters
}
```

## Adding This Version of the Module

If this repo is added as a subtree, then the version of the module should be close to the version shown here:

`x.y.z`

Note, however that subtrees can be altered as desired within repositories.

Further documentation on usage can be found [here](./docs).

Below is automatically generated documentation on this Terraform module using [terraform-docs][terraform-docs]

---

[terraform-docs]: https://github.com/terraform-docs/terraform-docs
