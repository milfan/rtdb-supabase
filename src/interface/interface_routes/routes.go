package interface_routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milfan/neoten-lib/lib_response_gin"
	controller_modules "github.com/milfan/rtdb-supabase/src/app/modules/controller_module"
	"github.com/milfan/rtdb-supabase/src/infra/configs/app_config"
)

func DefaultRoute(route *gin.Engine) {
	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "online!")
	})
}

func AllRoute(
	route *gin.Engine,
	response lib_response_gin.IResponseClient,
	appConf app_config.Config,
	controllerModule controller_modules.ControllerDicts,
) {
	mainRoute := route.Group("v1/")

	routes := mainRoute.Group("/order")
	routes.POST("insert", controllerModule.OrderController.Persist)
}
