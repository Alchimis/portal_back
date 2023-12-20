Для запуска:
1. `go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest`
2. `oapi-codegen -generate gorilla,types authentication/api/frontend/frontendapi.yaml > authentication/api/frontend/frontendapi.gen.go`
3. `oapi-codegen -generate gorilla,types documentation/api/frontend/frontendapi.yaml > documentation/api/frontend/frontendapi.gen.go`
4. `oapi-codegen -generate gorilla,types company/api/frontend/frontendapi.yaml > company/api/frontend/frontendapi.gen.go`
4. `go mod tidy -v`
5. `go run main.go`

Configure GoLand run config with the following env vars:
- DB_USER
- DB_PASSWORD
- DB_NAME
- DB_HOST
- BACKEND_PORT
- DB_DOCUMENTATION_NAME

![img.png](img/envVars.png)

Для работы с бд
1. установить postgres
2. миграции выполнятся автоматически

## Миграции

[Инструкция по установке утилиты для миграций](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)

Создание миграции: `migrate create -ext sql -dir {path-to-migrations-dir} {migartion-name}`

Миграция базы вручную (up-миграции также выполняются при запуске приложения):
up: `migrate -path {path-to-migrations-dir} -database "postgres://{DB_USER}:{DB_PASSWORD}@{DB_HOST}:5432/{DB_NAME}" up`
down: `migrate -path {path-to-migrations-dir} -database "postgres://{DB_USER}:{DB_PASSWORD}@{DB_HOST}:5432/{DB_NAME}" down`
