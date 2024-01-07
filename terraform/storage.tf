resource "yandex_storage_bucket" "bucket" {
  bucket     = "serverless-blog"
  access_key = yandex_iam_service_account_static_access_key.static_access_key.access_key
  secret_key = yandex_iam_service_account_static_access_key.static_access_key.secret_key
}

output "bucket_name" {
  value = yandex_storage_bucket.bucket.bucket
}