# Buildx Bake file: https://docs.docker.com/build/bake/
# Target platform is fixed to linux/amd64 (x86_64).
# Usage: docker buildx bake

variable "PLATFORMS" {
  default = ["linux/amd64"]
}

group "default" {
  targets = ["admin", "web"]
}

target "admin" {
  context    = "."
  dockerfile = "docker/admin.Dockerfile"
  platforms  = PLATFORMS
  tags       = ["sh/s/admin:latest"]
}

target "web" {
  context    = "."
  dockerfile = "docker/web.Dockerfile"
  platforms  = PLATFORMS
  tags       = ["sh/s/web:latest"]
}
