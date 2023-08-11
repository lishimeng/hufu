package midware

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter/midware/auth"
)

func WithAuth(handler ...iris.Handler) []iris.Handler {
	var handlers []iris.Handler
	handlers = append(handlers, auth.JwtBasic(), auth.Forbidden401Handler(auth.WithJsonResp()))
	handlers = append(handlers, handler...)
	return handlers
}
