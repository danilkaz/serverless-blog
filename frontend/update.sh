npm install
npm version patch
npm run build

s3cmd --access_key=${SERVERLESS_BLOG_BUCKET_ACCESS_KEY} --secret_key=${SERVERLESS_BLOG_BUCKET_SECRET_KEY} --host="storage.yandexcloud.net" --host-bucket="%(bucket)s.storage.yandexcloud.net" sync dist/ s3://${SERVERLESS_BLOG_BUCKET_NAME}/
s3cmd --access_key=${SERVERLESS_BLOG_BUCKET_ACCESS_KEY} --secret_key=${SERVERLESS_BLOG_BUCKET_SECRET_KEY} --host="storage.yandexcloud.net" --host-bucket="%(bucket)s.storage.yandexcloud.net" -r modify --add-header=content-type:application/javascript s3://${SERVERLESS_BLOG_BUCKET_NAME}/js/
s3cmd --access_key=${SERVERLESS_BLOG_BUCKET_ACCESS_KEY} --secret_key=${SERVERLESS_BLOG_BUCKET_SECRET_KEY} --host="storage.yandexcloud.net" --host-bucket="%(bucket)s.storage.yandexcloud.net" -r modify --add-header=content-type:text/css s3://${SERVERLESS_BLOG_BUCKET_NAME}/css/
