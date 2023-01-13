terraform {
  required_providers {
    fly = {
      source  = "fly-apps/fly"
      version = "0.0.20"
    }
  }
}

provider "fly" {
  # Configuration options
  fly_api_token = var.fly_api_token
  useinternaltunnel = true
  internaltunnelorg = var.fly_org_name
  internaltunnelregion = "fra"
}