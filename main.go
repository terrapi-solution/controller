package main

import (
	"github.com/terrapi-solution/controller/internal/core"
	"github.com/terrapi-solution/controller/internal/servers"
)

// @title TerrAPI Controller
// @version 1.0
// @description TerrAPI is a service designed to simplify and automate your deployments.
// @contact.name Support
// @contact.url https://github.com/terrapi-solution
// @contact.email contact@thomas-illiet.fr
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {

	coreSvc := core.GetInstance()
	servers.StartServers(coreSvc)
}
