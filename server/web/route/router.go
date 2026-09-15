package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/herhe-com/framework/http/middleware"
)

func Router(router *server.Hertz) {

	router.Use(middleware.Jwt())

	BasicRouter(router)

}
