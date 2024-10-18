package main

import (
	"github.com/gin-gonic/gin"
	"github.com/milfan/neoten-lib/lib_response_gin"
	"github.com/milfan/rtdb-supabase/src/infra/configs/app_config"
	interface_rest "github.com/milfan/rtdb-supabase/src/interface/rest"
)

func main() {
	appConf := app_config.InitAppConfig()

	ginServer := gin.Default()
	ginServer.Use(gin.Recovery())

	response := lib_response_gin.NewResponseClient()

	serv := interface_rest.NewServer(
		ginServer,
		appConf,
		response,
	)
	interface_rest.StartService(serv)
}
