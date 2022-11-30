# Reporter
Rest api for working reports

# Deploy
#### Requirements deploy: `docker`, `docker-compose`


1) ```git clone https://github.com/Ovsienko023/reporter.git```
2) ```cd reporter```
3) ```make run```

# Build
#### Requirements build: `golang 3.18+`

Для корректной работы прилажения необходимо накатить БД и подправить настройки подключения к бд`/reporter/api/infrastructure/configuration/config.go`: 
1) ```cd reporter/api``` 
2) ```go get```
3) ```go build main.go```


# Generate documentation swagger

1) ```go get github.com/swaggo/swag/cmd/swag``` или `go install github.com/swaggo/swag/cmd/swag@latest`
2) ```swag init``` (если команда не срабатывает: ```export GOBIN=$(go env GOPATH)/bin```)
3) http://0.0.0.0:8888/docs/index.html


# Init db

1) `cd reporter`
2) `make docker`

Так же можно руками накатить файл `init.sql` на установленную БД `postgres 14`. 

Путь к файлу: `/reporter/database/init.sql`


# useful links
- `https://github.com/swaggo/http-swagger`
- `https://github.com/swaggo/swag`
- `https://github.com/swaggo/swag#api-operation`


Чтобы удалить все остановленные контейнеры и неиспользуемые образы (а не только образы, не связанные с контейнерами
docker system prune -a