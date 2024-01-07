# Serverless Blog

## Необходимые утилиты

- `yc`
- `terraform`
- `docker`
- `jq`

## Инструкция по развёртыванию

### 1. Инициализируйте необходимые переменные

```sh
yc init
export YC_TOKEN=$(yc iam create-token)
export YC_CLOUD_ID=$(yc config get cloud-id)
export YC_FOLDER_ID=$(yc config get folder-id)
```

### 2. Создайте Container Registry, сервисный аккаунт, экземпляр YDB и раздайте необходимые права

```sh
cd terraform
terraform init
terraform apply -target yandex_container_registry.registry -target yandex_iam_service_account.service_account -target yandex_resourcemanager_folder_iam_member.roles -target yandex_ydb_database_serverless.ydb
export SERVERLESS_BLOG_CONTAINER_REGISTRY_ID=$(terraform output -raw container_registry_id)
export SERVERLESS_BLOG_SERVICE_ACCOUNT_ID=$(terraform output -raw service_account_id)
export SERVERLESS_BLOG_YDB_ENDPOINT=$(terraform output -raw ydb_endpoint)
export SERVERLESS_BLOG_YDB_ACCESS_KEY=$(terraform output -raw service_account_access_key)
```

### 3. Обновите версию бэкенда

```sh
yc container registry configure-docker
export TF_VAR_SERVERLESS_CONTAINER_ID=$(yc sls container create --name serverless-blog --format json | jq -r '.id')
cd ../backend
./update.sh
```

### 4. Создайте оставшуюся инфраструктуру

```sh
cd ../terraform
terraform apply
```

### 5. Обновите версию фронтенда

```sh
export SERVERLESS_BLOG_BUCKET_NAME=$(terraform output -raw bucket_name)
export SERVERLESS_BLOG_BUCKET_ACCESS_KEY=$(terraform output -raw bucket_access_key)
export SERVERLESS_BLOG_BUCKET_SECRET_KEY=$(terraform output -raw bucket_secret_key)
export VUE_APP_SERVERLESS_BLOG_API_URL="https://$(terraform output -raw api_url)"
cd ../frontend
./update.sh
```