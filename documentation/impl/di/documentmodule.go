package di

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"net/http"
	"portal_back/authentication/api/internalapi"
	frontendapi "portal_back/documentation/api/frontend"
	"portal_back/documentation/cmd"
	"portal_back/documentation/impl/app/sections"
	"portal_back/documentation/impl/infrastructure/sql"
	"portal_back/documentation/impl/infrastructure/transport"
)

func InitDocumentModule(authRequestService internalapi.AuthRequestService, config cmd.Config) *pgx.Conn {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	conn, _ := pgx.Connect(context.Background(), connStr)

	sectionRepository := sql.NewSectionRepository(conn)
	service := sections.NewSectionService(sectionRepository)
	server := transport.NewFrontendServer(service, authRequestService)
	http.Handle("/documentation/", frontendapi.Handler(server))
	return conn
}
