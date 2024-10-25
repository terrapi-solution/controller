package main

import (
	"github.com/terrapi-solution/controller/internal/core"
	"github.com/terrapi-solution/controller/internal/servers"
)

func main() {

	coreSvc := core.GetInstance()
	servers.StartServers(coreSvc)
}
