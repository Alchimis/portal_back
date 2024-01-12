package di

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	authcmd "portal_back/authentication/cmd"
	di "portal_back/authentication/cmd"
	companycmd "portal_back/company/cmd"
	"portal_back/documentation/cmd"
	documentationDi "portal_back/documentation/impl/di"
	rolesDi "portal_back/roles/impl/di"
)

func InitAppModule() {
	authService, userRequestService, authConn, err := di.InitAuthModule(authcmd.NewConfig())
	if authConn == nil {
		fmt.Printf("Can't connect to teamtells database")
		return
	}
	defer authConn.Close(context.Background())

	documentConnection := documentationDi.InitDocumentModule(authService, cmd.NewConfig())
	defer documentConnection.Close(context.Background())

	rolesModule := rolesDi.InitRolesModule()

	companycmd.InitCompanyModule(companycmd.NewConfig(), authService, userRequestService, rolesModule)

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	err = http.ListenAndServe(":"+appPort, nil)
	if err != nil {
		log.Panic("ListenAndServe: ", err)
	}
}

func Migrate() {
	err := authcmd.Migrate(authcmd.NewConfig())
	if err != nil {
		log.Fatal("failed migrate auth module:", err)
	}

	err = cmd.Migrate(cmd.NewConfig())

	if err != nil {
		log.Fatal("failed migrate documentation module:", err)
	}
}
