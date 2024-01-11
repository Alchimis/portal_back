package di

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"net/http"
	"portal_back/authentication/api/internalapi"
	"portal_back/core/network"
	frontendapi "portal_back/documentation/api/frontend"
	"portal_back/documentation/cmd"
	"portal_back/documentation/impl/app/sections"
	"portal_back/documentation/impl/infrastructure/sql"
	"portal_back/documentation/impl/infrastructure/transport"
)

func InitDocumentModule(authRequestService internalapi.AuthRequestService, config cmd.Config) *pgx.Conn {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, 5432, config.DBUser, config.DBPassword, config.DBName)

	conn, _ := pgx.Connect(context.Background(), connStr)

	sectionRepository := sql.NewSectionRepository(conn)
	service := sections.NewSectionService(sectionRepository)
	server := transport.NewFrontendServer(service, authRequestService)

	router := mux.NewRouter()
	router.MethodNotAllowedHandler = network.MethodNotAllowedHandler()

	options := frontendapi.GorillaServerOptions{
		BaseRouter: router,
		Middlewares: []frontendapi.MiddlewareFunc{func(handler http.Handler) http.Handler {
			return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				network.SetCorsHeaders(w, r)
				handler.ServeHTTP(w, r)
			}))
		}},
	}
	r := frontendapi.HandlerWithOptions(server, options)

	http.Handle("/documentation/", r)
	return conn
}
