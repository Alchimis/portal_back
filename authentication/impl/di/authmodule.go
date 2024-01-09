package di

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"portal_back/authentication/api/frontend"
	"portal_back/authentication/api/internalapi"
	"portal_back/authentication/impl/app/auth"
	"portal_back/authentication/impl/app/authrequest"
	"portal_back/authentication/impl/app/token"
	"portal_back/authentication/impl/infrastructure/sql"
	"portal_back/authentication/impl/infrastructure/transport"
	"time"
)

func InitAuthModule() (internalapi.AuthRequestService, *pgx.Conn) {
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

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, 5432, dbUser, dbPassword, dbName)

	conn, err := ConnectLoop(connStr, 30*time.Second)

	if err != nil {
		fmt.Printf("Error asdfasdfasdf!!!!! %s", err)
	} else {
		fmt.Printf("CONNECTED adsfadsfafds!!")
	}

	repo := sql.NewTokenStorage(conn)
	tokenService := token.NewService(repo)

	authRepo := sql.NewAuthRepository(conn)
	authService := auth.NewService(authRepo, tokenService)
	server := transport.NewServer(authService, tokenService)
	authRequestService := authrequest.NewService()

	router := mux.NewRouter()
	router.MethodNotAllowedHandler = methodNotAllowedHandler()

	options := frontendapi.GorillaServerOptions{
		BaseRouter: router,
		Middlewares: []frontendapi.MiddlewareFunc{func(handler http.Handler) http.Handler {
			return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				setCorsHeaders(w)
				handler.ServeHTTP(w, r)
			}))
		}},
	}
	r := frontendapi.HandlerWithOptions(server, options)
	http.Handle("/authorization/", r)

	return authRequestService, conn
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
			if db != nil {
				fmt.Println("Connected???")
			}
			fmt.Println("Error asdfasdfasdf!!!!!")
			fmt.Println(err)
		}
	}
}

func methodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			setCorsHeaders(w)
		}
		w.WriteHeader(http.StatusNoContent)
	})
}

func setCorsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, X-user-id, X-organization-id")
	w.Header().Set("Access-Control-Allow-Origin", "https://dev4.env.teamtells.ru")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
