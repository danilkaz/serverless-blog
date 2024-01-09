BUCKET_NAME=$(terraform -chdir=terraform output -raw bucket_name)
BUCKET_ACCESS_KEY=$(terraform -chdir=terraform output -raw bucket_access_key)
BUCKET_SECRET_KEY=$(terraform -chdir=terraform output -raw bucket_secret_key)
export VUE_APP_SERVERLESS_BLOG_API_URL="https://$(terraform -chdir=terraform output -raw api_url)"

npm --prefix frontend install
npm --prefix frontend version patch
npm --prefix frontend run build

s3cmd --access_key=${BUCKET_ACCESS_KEY} --secret_key=${BUCKET_SECRET_KEY} --host="storage.yandexcloud.net" --host-bucket="%(bucket)s.storage.yandexcloud.net" sync frontend/dist/ s3://${BUCKET_NAME}/
s3cmd --access_key=${BUCKET_ACCESS_KEY} --secret_key=${BUCKET_SECRET_KEY} --host="storage.yandexcloud.net" --host-bucket="%(bucket)s.storage.yandexcloud.net" -r modify --add-header=content-type:application/javascript s3://${BUCKET_NAME}/js/
s3cmd --access_key=${BUCKET_ACCESS_KEY} --secret_key=${BUCKET_SECRET_KEY} --host="storage.yandexcloud.net" --host-bucket="%(bucket)s.storage.yandexcloud.net" -r modify --add-header=content-type:text/css s3://${BUCKET_NAME}/css/
