package di

import (
	"context"
	"fmt"
	"net/http"
	"os"
	frontendapi "portal_back/role/api/frontend"
	"portal_back/role/impl/app/role"
	"portal_back/role/impl/infrasructure/sql"
	"portal_back/role/impl/infrasructure/transport"

	"github.com/jackc/pgx/v5"
)

func InitRolesModule() (role.RoleService, *pgx.Conn, error) {

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "password"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "app"
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	dbStringConnection := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", dbUser, dbPassword, dbHost, dbName)
	conn, _ := pgx.Connect(context.Background(), dbStringConnection)

	roleRepository := sql.NewRepository(conn)
	roleService := role.NewService(roleRepository)
	roleServer := transport.NewServer(roleService)
	http.Handle("/role/", frontendapi.Handler(roleServer))
	return roleService, conn, nil
}
