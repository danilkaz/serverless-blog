yc serverless container revision list \
  --container-id ${TF_VAR_SERVERLESS_CONTAINER_ID} \
  --format json | jq -r '[.[] | select(.status == "ACTIVE")] | if length == 0 then 0 else .[0].provision_policy.min_instances // 1 end'
