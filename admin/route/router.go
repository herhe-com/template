package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/herhe-com/framework/http/middleware"
	"github.com/herhe-com/template/admin/constants"
)

func Router(router *server.Hertz) {

	router.Use(middleware.Jwt(constants.JwtOfIssuerWithAdmin))

	BasicRouter(router)

	SiteRouter(router)
}
