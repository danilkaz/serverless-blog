resource "yandex_ydb_database_serverless" "ydb" {
  name = "serverless-blog"
}

output "ydb_endpoint" {
  value = yandex_ydb_database_serverless.ydb.ydb_full_endpoint
}