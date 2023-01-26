# Fly Database
resource "fly_app" "dicetrace_database_app" {
  name = var.fly_database_name
  org  = var.fly_org_name
}

resource "fly_ip" "dicetrace_database_ipv4" {
  app = fly_app.dicetrace_database_app.name
  type = "v4"
  depends_on = [
    fly_app.dicetrace_database_app
  ]
}

resource "fly_volume" "dicetrace_database_volume" {
  for_each = var.fly_database_volume_regions
  name = "${var.fly_database_volume_name}_${each.value}"
  app = fly_app.dicetrace_database_app.name
  size = var.fly_database_volume_size
  region = each.value

  depends_on = [
    fly_app.dicetrace_database_app
  ]
}

resource "fly_machine" "dicetrace_database_machine" {
  for_each = var.fly_database_regions
  app    = fly_app.dicetrace_database_app.name 
  region = each.value
  name   = "${fly_app.dicetrace_database_app.name}-${each.value}"
  image  = var.fly_database_image
  services = [
    {
      ports = [
        {
          port     = 27017
          handlers = ["http"]
        },
      ]
      "protocol" : "tcp",
      "internal_port" : 27017
    },
  ]
  cpus = 1
  memorymb = 256
  mounts = [
    {
      volume = fly_volume.dicetrace_database_volume[each.value].name
      path = "/data/db"
      encrypted = false
      size_gb = var.fly_database_volume_size
    }
  ]

  depends_on = [
    fly_app.dicetrace_database_app,
    fly_volume.dicetrace_database_volume
  ]
}

# Fly Backend
resource "fly_app" "dicetrace_backend_app" {
  name = var.fly_backend_name
  org  = var.fly_org_name
}

resource "fly_ip" "dicetrace_backend_ipv4" {
  app = fly_app.dicetrace_backend_app.name
  type = "v4"
  depends_on = [
    fly_app.dicetrace_backend_app
  ]
}

resource "fly_ip" "dicetrace_backend_ipv6" {
  app = fly_app.dicetrace_backend_app.name
  type = "v6"
  depends_on = [
    fly_app.dicetrace_backend_app
  ]
}

resource "fly_machine" "dicetrace_backend_machine" {
  for_each = var.fly_backend_regions
  app    = fly_app.dicetrace_backend_app.name 
  region = each.value
  name   = "${fly_app.dicetrace_backend_app.name}-${each.value}"
  image  = var.fly_backend_image

  env = {
    "DB_IP" = "${each.value}.${fly_app.dicetrace_database_app.name}.internal"
  }

  services = [
    {
      ports = [
        {
          port     = 443
          handlers = ["tls", "http"]
        },
        {
          port     = 80
          handlers = ["http"]
        }
      ]
      "protocol" : "tcp",
      "internal_port" : 80
    },
  ]
  cpus = 1
  memorymb = 256

  depends_on = [
    fly_machine.dicetrace_database_machine,
    fly_app.dicetrace_backend_app,
  ]
}