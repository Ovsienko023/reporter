# Reporter
Rest api for working reports

# Requirements

- docker
- docker-compose

# Starting

1) ```cd reporter```
2) ```docker-compose up --build``` (Для запуска в фоне, использовать команду: ```docker-compose up -d --build```)
3) ```docker-compose down --volumes``` (Еслинужно отчистить БД)
4) `sudo docker stop $(docker ps -a -q)` (Остановка всех запущенных контейнеров)
# Generate documentation swagger

1) ```go get github.com/swaggo/swag/cmd/swag```
2) ```swag init``` (если команда не срабатывает: ```export GOBIN=$(go env GOPATH)/bin```)
3) http://0.0.0.0:8888/docs/index.html

# useful links
- `https://github.com/swaggo/http-swagger`
- `https://github.com/swaggo/swag`
- `https://github.com/swaggo/swag#api-operation`


Чтобы удалить все остановленные контейнеры и неиспользуемые образы (а не только образы, не связанные с контейнерами
docker system prune -a

gRPC builds

go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc api/proto/*.proto --go_out=. --proto_path=. --go-grpc_out=.