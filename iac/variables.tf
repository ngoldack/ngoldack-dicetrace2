# Fly General
variable "fly_org_name" {
  type    = string
  default = "dicetrace"
}

variable "fly_api_token" {
  type = string
  sensitive = true
}

# Fly Database
variable "fly_database_regions" {
  type = set(string)
  default = [ "fra" ]
}

variable "fly_database_name" {
  type = string
  default = "dicetrace-database"
}

variable "fly_database_image" {
  type = string
  default = "mongo:latest"
}

# Fly Database volume
variable "fly_database_volume_regions" {
    type = set(string)
    default = [ "fra" ]
}

variable "fly_database_volume_name" {
    type = string
    default = "dicertace_database_volume"
}

variable "fly_database_volume_size" {
    type = number
    default = 1
}

# Fly Backend
variable "fly_backend_regions" {
    type = set(string)
    default = [ "fra" ]
}

variable "fly_backend_name" {
    type = string
    default = "dicetrace-backend"
}

variable "fly_backend_image" {
    type = string
    default = "traefik/whoami:latest"
}
