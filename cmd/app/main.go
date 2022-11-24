package main

import "github.com/samuellfa/copa-do-mundo-golang/internal/api"

// @title          Swagger Example API
// @version        1.0
// @description    This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:3333
// @BasePath /v1
func main() {
	api := api.New()
	api.SetupAndListen()
}
