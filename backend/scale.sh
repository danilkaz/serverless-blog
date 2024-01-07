REPLICAS_COUNT=$(yc serverless container revision list \
                   --container-id ${TF_VAR_SERVERLESS_CONTAINER_ID} \
                   --format json | jq -r '[.[] | select(.status == "ACTIVE")] | if length == 0 then 0 else .[0].provision_policy.min_instances // 1 end')

yc serverless container revision deploy \
  --container-name serverless-blog \
  --image ${SERVERLESS_BLOG_IMAGE_NAME} \
  --cores 1 \
  --memory 128M \
  --concurrency 1 \
  --execution-timeout 30s \
  --service-account-id ${SERVERLESS_BLOG_SERVICE_ACCOUNT_ID} \
  --environment SERVERLESS_BLOG_YDB_ENDPOINT=$(terraform output -raw ydb_endpoint),SERVERLESS_BLOG_YDB_API_TOKEN=${YC_TOKEN}
  --min-instances $((REPLICAS_COUNT+1))
