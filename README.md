# Reporter
Rest api for working reports

# Requirements

- docker
- docker-compose

# Starting

1) ```cd reporter```
2) ```docker-compose up --build``` (Для запуска в фоне, использовать команду: ```docker-compose up -d --build```)
3) ```docker-compose down --volumes``` (Еслинужно отчистить БД)

# Generate documentation swagger

1) ```go get github.com/swaggo/swag/cmd/swag```
2) ```swag init``` (если команда не срабатывает: ```export GOBIN=$(go env GOPATH)/bin```)
3) http://0.0.0.0:8888/docs/index.html

# useful links
- `https://github.com/swaggo/http-swagger`
- `https://github.com/swaggo/swag`
- `https://github.com/swaggo/swag#api-operation`