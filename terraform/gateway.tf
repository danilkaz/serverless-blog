variable "SERVERLESS_CONTAINER_ID" {
  type    = string
  default = ""
}

resource "yandex_api_gateway" "gateway" {
  name = "serverless-blog"

  spec = templatefile("openapi.yaml", {
    SERVICE_ACCOUNT_ID      = yandex_iam_service_account.service_account.id,
    SERVERLESS_CONTAINER_ID = var.SERVERLESS_CONTAINER_ID,
    BUCKET_NAME             = yandex_container_registry.registry.name
  })
}

output "api_url" {
  value = yandex_api_gateway.gateway.domain
}