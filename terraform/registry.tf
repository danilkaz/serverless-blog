resource "yandex_container_registry" "registry" {
  name = "serverless-blog"
}

output "container_registry_id" {
  value = yandex_container_registry.registry.id
}
