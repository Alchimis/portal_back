package main

import "portal_back/di"

func main() {
	migrate()
	di.InitAppModule()
}
