package main

import "portal_back/di"

func main() {
	di.Migrate()
	di.InitAppModule()
}
