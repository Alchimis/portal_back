package cmd

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"net/http"
	"portal_back/authentication/api/internalapi"
	"portal_back/company/api/frontend"
	"portal_back/company/impl/app/department"
	"portal_back/company/impl/app/employeeaccount"
	"portal_back/company/impl/infrastructure/sql"
	"portal_back/company/impl/infrastructure/transport"
	rolesapi "portal_back/roles/api/internalapi"
)

func InitCompanyModule(config Config, authApi internalapi.AuthRequestService, userApi internalapi.UserRequestService, rolesApi rolesapi.RolesRequestService) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBName)

	conn, _ := pgx.Connect(context.Background(), connStr)

	accountRepo := sql.NewEmployeeAccountRepository(conn)
	accountService := employeeaccount.NewService(accountRepo, userApi)

	departmentRepo := sql.NewDepartmentRepository(conn)
	departmentService := department.NewService(departmentRepo, accountService)

	server := transport.NewServer(accountService, departmentService, rolesApi, authApi)

	http.Handle("/", frontendapi.Handler(server))
}
