echo `cat version.json | jq '.version = .version + 1'` > version.json;
SERVERLESS_BLOG_BACKEND_VERSION=$(jq -r '.version' version.json)
SERVERLESS_BLOG_IMAGE_NAME="cr.yandex/${SERVERLESS_BLOG_CONTAINER_REGISTRY_ID}/serverless-blog:${SERVERLESS_BLOG_BACKEND_VERSION}"

docker build -t ${SERVERLESS_BLOG_IMAGE_NAME} .
docker push ${SERVERLESS_BLOG_IMAGE_NAME}

yc serverless container revision deploy --container-name serverless-blog \
  --image ${SERVERLESS_BLOG_IMAGE_NAME} \
  --cores 1 \
  --memory 128M \
  --concurrency 1 \
  --execution-timeout 30s \
  --service-account-id ${SERVERLESS_BLOG_SERVICE_ACCOUNT_ID} \
  --environment SERVERLESS_BLOG_YDB_ENDPOINT=${SERVERLESS_BLOG_YDB_ENDPOINT},SERVERLESS_BLOG_YDB_ACCESS_TOKEN=${YC_TOKEN},SERVERLESS_BLOG_BACKEND_VERSION=${SERVERLESS_BLOG_BACKEND_VERSION}