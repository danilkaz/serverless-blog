locals {
  roles = toset([
    "container-registry.images.puller",
    "serverless.containers.invoker",
    "storage.editor",
    "storage.configurer"
  ])
}

resource "yandex_iam_service_account" "service_account" {
  name = "serverless-blog-sa"
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

resource "yandex_iam_service_account" "ydb_service_account" {
  name = "serverless-blog-ydb-sa"
}

resource "yandex_resourcemanager_folder_iam_member" "ydb_role" {
  folder_id = yandex_iam_service_account.service_account.folder_id
  role      = "yds.editor"
  member    = "serviceAccount:${yandex_iam_service_account.ydb_service_account.id}"
}

resource "yandex_iam_service_account_key" "ydb_service_account_key" {
  service_account_id = yandex_iam_service_account.ydb_service_account.id
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

output "ydb_service_account_key" {
  value = jsonencode({
    "id"=yandex_iam_service_account_key.ydb_service_account_key.id,
    "service_account_id"=yandex_iam_service_account_key.ydb_service_account_key.service_account_id,
    "created_at"=yandex_iam_service_account_key.ydb_service_account_key.created_at,
    "key_algorithm"=yandex_iam_service_account_key.ydb_service_account_key.key_algorithm,
    "public_key"=yandex_iam_service_account_key.ydb_service_account_key.public_key,
    "private_key"=yandex_iam_service_account_key.ydb_service_account_key.private_key
  })  
  sensitive = true
}
