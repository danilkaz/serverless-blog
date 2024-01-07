locals {
  roles = toset([
    "container-registry.images.puller",
    "serverless.containers.invoker",
    "storage.admin",
    "ydb.admin",
    "yds.admin"
  ])
}

resource "yandex_iam_service_account" "service_account" {
  name = "serverless-blog-container-sa"
}

resource "yandex_resourcemanager_folder_iam_member" "roles" {
  for_each  = local.roles
  folder_id = yandex_iam_service_account.service_account.folder_id
  role      = each.key
  member    = "serviceAccount:${yandex_iam_service_account.service_account.id}"
}

resource "yandex_iam_service_account_static_access_key" "static_access_key" {
  service_account_id = yandex_iam_service_account.service_account.id
}

output "service_account_id" {
  value = yandex_iam_service_account.service_account.id
}

output "bucket_access_key" {
  value = yandex_iam_service_account_static_access_key.static_access_key.access_key
}

output "bucket_secret_key" {
  value = yandex_iam_service_account_static_access_key.static_access_key.secret_key
  sensitive = true
}
