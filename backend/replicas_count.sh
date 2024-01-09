yc serverless container revision list \
  --container-id $(yc serverless container get --name serverless-blog --format json | jq -r '.id') \
  --format json | jq -r '[.[] | select(.status == "ACTIVE")] | if length == 0 then 0 else .[0].provision_policy.min_instances // 1 end'
