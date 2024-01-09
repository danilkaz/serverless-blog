# Serverless Blog

Простой блог, построенный на бессерверных технологиях (Serverless Containers, Serverless YDB)

Развёрнутое приложение - тут https://blog.tumba-yumba.fun

## Необходимые утилиты для развёртывания

- `yc`
- `terraform`
- `docker`
- `jq`
- `npm`
- `s3cmd`

## Инструкция по развёртыванию

### 0. Инициализируйте `yc` или настройте на нужное облако

```sh
yc init
```

### 1. Инициализируйте необходимые переменные

```sh
export YC_TOKEN=$(yc iam create-token)
export YC_CLOUD_ID=$(yc config get cloud-id)
export YC_FOLDER_ID=$(yc config get folder-id)
```

### 2. Создайте Container Registry, сервисный аккаунт, экземпляр YDB и раздайте необходимые права

```sh
terraform -chdir=terraform init
terraform -chdir=terraform apply -target yandex_container_registry.registry -target yandex_iam_service_account.service_account -target yandex_ydb_database_serverless.ydb -target yandex_resourcemanager_folder_iam_member.roles 
```

### 3. Обновите версию бэкенда

```sh
yc container registry configure-docker
export TF_VAR_SERVERLESS_CONTAINER_ID=$(yc sls container create --name serverless-blog --format json | jq -r '.id')
./backend/update.sh
```

### 4. Создайте оставшуюся инфраструктуру

```sh
terraform -chdir=terraform apply
```

### 5. Обновите версию фронтенда

```sh
./frontend/update.sh
```

Приложение находится по адресу из вывода команды `echo "https://$(terraform -chdir=terraform output -raw api_url)"`

## Скрипты для автоматизации

Для использования скриптов выполните пункты [0](#0-инициализируйте-yc-или-настройте-на-нужное-облако) и [1](#1-инициализируйте-необходимые-переменные) из инструкции выше

- `./backend/scale.sh` - добавление новой реплики контейнера
- `./backend/update.sh` - обновление и деплой новой версии бэкенда
- `./frontend/update.sh` - обновление и деплой новой версии фронтенда