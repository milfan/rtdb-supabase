package interface_rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milfan/neoten-lib/lib_http_request"
	"github.com/milfan/neoten-lib/lib_response_gin"
	controller_modules "github.com/milfan/rtdb-supabase/src/app/modules/controller_module"
	"github.com/milfan/rtdb-supabase/src/app/modules/usecase_module"
	"github.com/milfan/rtdb-supabase/src/infra/configs/app_config"
	"github.com/milfan/rtdb-supabase/src/interface/interface_routes"
)

func NewServer(
	server *gin.Engine,
	conf app_config.Config,
	response lib_response_gin.IResponseClient,
) *http.Server {
	httpRequestUtils := lib_http_request.NewHttpRequestUtils()

	usecaseModule := usecase_module.LoadUsecaseModules(
		httpRequestUtils,
	)
	controllerModule := controller_modules.LoadControllers(
		response,
		usecaseModule,
	)

	interface_routes.DefaultRoute(server)
	interface_routes.AllRoute(
		server,
		response,
		conf,
		controllerModule,
	)

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", conf.HttpConfig.HttpPort),
		Handler: server.Handler(),
	}
}
