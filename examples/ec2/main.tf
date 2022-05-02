module "cluster" {
  source = "../.."

  name = var.name

  ec2_backed = true

  // This is just to make sure the destroy goes right
  protect_from_scale_in = false

  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo
}
