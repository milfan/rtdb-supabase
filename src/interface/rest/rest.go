package interface_rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/milfan/neoten-lib/lib_http_request"
	"github.com/milfan/neoten-lib/lib_response_gin"
	controller_modules "github.com/milfan/rtdb-supabase/src/app/modules/controller_module"
	"github.com/milfan/rtdb-supabase/src/app/modules/repository_module"
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

	repositoryModules := repository_module.LoadRepositoryModules(
		conf.RTDBConfig,
		httpRequestUtils,
	)

	usecaseModules := usecase_module.LoadUsecaseModules(
		repositoryModules,
	)
	controllerModule := controller_modules.LoadControllers(
		response,
		usecaseModules,
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

func StartService(srv *http.Server) {
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	gracefullShutdown(srv)
}

func gracefullShutdown(srv *http.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// Instead of using select with a single case, we directly wait for the context's Done signal.
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")

	log.Println("Server exiting")
}
