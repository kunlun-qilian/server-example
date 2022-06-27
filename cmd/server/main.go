package main

import (
	_ "KunLunQiLian/server-example/cmd/server/docs"
	"KunLunQiLian/server-example/cmd/server/global"
	"KunLunQiLian/server-example/cmd/server/router"

	"github.com/KunLunQiLian/confserver"
	"github.com/spf13/cobra"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v1
func main() {

	confserver.Execute(func(cmd *cobra.Command, args []string) {
		s := global.Config.Server
		router.NewRooter(s.Engine())
		s.Serve()
	})
}
