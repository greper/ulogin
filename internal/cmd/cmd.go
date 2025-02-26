package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/greper/ulogin/internal/middleware"

	admin_app "github.com/greper/ulogin/internal/controller/admin/app"
	admin_auth "github.com/greper/ulogin/internal/controller/admin/auth"
	admin_mine "github.com/greper/ulogin/internal/controller/admin/mine"

	user_app "github.com/greper/ulogin/internal/controller/user/app"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/api/admin", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(middleware.JwtAuthMiddleware)
				group.Bind(
					admin_app.NewV1(),
					admin_auth.NewV1(),
					admin_mine.NewV1(),
				)
			})

			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					user_app.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
