CONTAINER_REGISTRY_ID=$(terraform -chdir=terraform output -raw container_registry_id)
BACKEND_VERSION=$(jq -r '.version' backend/version.json)
IMAGE_NAME="cr.yandex/${CONTAINER_REGISTRY_ID}/serverless-blog:${BACKEND_VERSION}"
SERVICE_ACCOUNT_ID=$(terraform -chdir=terraform output -raw service_account_id)
YDB_ENDPOINT=$(terraform -chdir=terraform output -raw ydb_endpoint)
YDB_SERVICE_ACCOUNT_KEY_ENCODED=$(terraform -chdir=terraform output -raw ydb_service_account_key | base64 -w 0)

REPLICAS_COUNT=$(./backend/replicas_count.sh)

yc serverless container revision deploy \
  --container-name serverless-blog \
  --image ${IMAGE_NAME} \
  --cores 1 \
  --memory 128M \
  --concurrency 1 \
  --execution-timeout 30s \
  --service-account-id ${SERVICE_ACCOUNT_ID} \
  --environment SERVERLESS_BLOG_YDB_ENDPOINT=${YDB_ENDPOINT},SERVERLESS_BLOG_YDB_SERVICE_ACCOUNT_KEY_ENCODED=${YDB_SERVICE_ACCOUNT_KEY_ENCODED},SERVERLESS_BLOG_BACKEND_VERSION=${BACKEND_VERSION} \
  --min-instances $((REPLICAS_COUNT+1))
