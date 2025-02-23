package cmd

import (
	"context"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"net/http"
	"portal_back/authentication/api/frontend"
	"portal_back/authentication/api/internalapi"
	"portal_back/authentication/impl/app/auth"
	"portal_back/authentication/impl/app/authrequest"
	"portal_back/authentication/impl/app/token"
	"portal_back/authentication/impl/app/userrequest"
	"portal_back/authentication/impl/infrastructure/sql"
	"portal_back/authentication/impl/infrastructure/transport"
	"portal_back/core/network"
	"time"
)

func InitAuthModule(config Config) (internalapi.AuthRequestService, internalapi.UserRequestService, *pgx.Conn, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, 5432, config.DBUser, config.DBPassword, config.DBName)

	conn, _ := ConnectLoop(connStr, 30*time.Second)

	repoId := sql.NewTokenStorage(conn)
	tokenService := token.NewService(repoId)

	authRepo := sql.NewAuthRepository(conn)
	authService := auth.NewService(authRepo, tokenService)
	server := transport.NewServer(authService, tokenService)
	authRequestService := authrequest.NewService()
	userRequestService := userrequest.NewService(authService)

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
	http.Handle("/authorization/", r)

	return authRequestService, userRequestService, conn, nil
}

func ConnectLoop(connStr string, timeout time.Duration) (*pgx.Conn, error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeoutExceeded := time.After(timeout)
	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %s timeout", timeout)

		case <-ticker.C:
			db, err := pgx.Connect(context.Background(), connStr)
			if err == nil {
				return db, nil
			}
		}
	}
}
