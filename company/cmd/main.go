package cmd

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"net/http"
	"portal_back/authentication/api/internalapi"
	"portal_back/company/api/frontend"
	"portal_back/company/impl/app/department"
	"portal_back/company/impl/app/employeeaccount"
	"portal_back/company/impl/infrastructure/sql"
	"portal_back/company/impl/infrastructure/transport"
	"portal_back/core/network"
	rolesapi "portal_back/roles/api/internalapi"
)

func InitCompanyModule(config Config, authApi internalapi.AuthRequestService, userApi internalapi.UserRequestService, rolesApi rolesapi.RolesRequestService) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, 5432, config.DBUser, config.DBPassword, config.DBName)

	conn, _ := pgx.Connect(context.Background(), connStr)

	accountRepo := sql.NewEmployeeAccountRepository(conn)
	accountService := employeeaccount.NewService(accountRepo, userApi)

	departmentRepo := sql.NewDepartmentRepository(conn)
	departmentService := department.NewService(departmentRepo, accountService)

	server := transport.NewServer(accountService, departmentService, rolesApi, authApi)

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

	http.Handle("/", r)
}
