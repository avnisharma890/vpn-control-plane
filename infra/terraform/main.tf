terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0.2"
    }
  }
}

provider "docker" {}

# Prometheus Container

resource "docker_container" "prometheus" {
  name  = "prometheus"
  image = "prom/prometheus"

  ports {
    internal = 9090
    external = 9090
  }

  # Copy config inside container
  upload {
    content = file("${path.module}/../prometheus/prometheus.yml")
    file    = "/etc/prometheus/prometheus.yml"
  }
}

# Grafana Container

resource "docker_container" "grafana" {
  name  = "grafana"
  image = "grafana/grafana"

  ports {
    internal = 3000
    external = 3000
  }
}
