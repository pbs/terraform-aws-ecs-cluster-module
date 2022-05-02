locals {
  name   = var.name != null ? var.name : var.product
  vpc_id = var.vpc_id != null ? var.vpc_id : data.aws_vpc.vpc[0].id

  subnets = var.subnets != null ? var.subnets : concat(
    tolist(data.aws_subnets.public_subnets[0].ids),
    tolist(data.aws_subnets.private_subnets[0].ids)
  )

  vpc_zone_identifier = local.subnets
  role_policy_json    = var.role_policy_json != null ? var.role_policy_json : data.aws_iam_policy_document.policy_doc[0].json

  creator = "terraform"

  tags = merge(
    var.tags,
    {
      Name                                      = local.name
      "${var.organization}:billing:product"     = var.product
      "${var.organization}:billing:environment" = var.environment
      creator                                   = local.creator
      provisioner                               = data.aws_caller_identity.current.user_id
      repo                                      = var.repo
    }
  )
}
