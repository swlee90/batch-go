# 사용 Libaray
- zap
- sql

# 사용방법
## 1. config.yml을 생성
config.yml을 사용합니다. dbconfig를 사용하지 않고 kubernetes의 ENV나 docker ENV를 사용가능합니다.
```yaml
logger:
  filename: log.log
  level: DEBUG
  env: Dev
dbconfig:
  db_url: <DB_URL 작성>
  db_user: <db_user 작성>
  db_password: <db_userpassword 작성>
  db_name: <db_dababase 작성>
  db_port: <db_port 작성>
  db_table: <db_table 작성>
```

## 2. kubernetes 배포 방법
yaml을 아래와같이 작성하여 배포하면 배포가능
config.yml은 필요합니다.. config.yml 없이 실행은 현재 불가능..
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: batch-go
  labels:
    app: batch-go
spec:
  replicas: 3
  selector:
    matchLabels:
      app: batch-go
  template:
    metadata:
      labels:
        app: batch-go
    spec:
      containers:
      - name: go-batch
        image: sw90lee/batch-go
        ports:
        - containerPort: 80
        env:
        - name: DB_URL
          value: <DB_URL 작성>
        - name: DB_USER
          value: <DB_URL 작성>
        - name: DB_PASSWORD
          value: <DB_URL 작성>
        - name: DB_NAME
          value: <DB_URL 작성>
        - name: DB_PORT
          value: <DB_URL 작성>
        - name: DB_TABLE
          value: <DB_TABLE 작성>
```
