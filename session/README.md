# Use Redis Session in Golang

## set redis local

須設置環境變數 "REDISHOST=localhost","REDISPORT=6379"

`docker run --name session -p 6379:6379 -d redis`

## deploy to GCP

發布 cloud run

`gcloud run deploy session --region=asia-east1  --source .`

連接與創建 redis

```sh
gcloud beta run integrations create \
--type=redis \
--service=session \
--parameters=memory-size-gb=1
```

note: my print like this `[redis] integration [redis-1] has been created successfully.`

清除連接和刪除 redis (INTEGRATION_NAME 可以從 list 查看)

```sh
gcloud beta run integrations list
gcloud beta run integrations delete INTEGRATION_NAME
```
