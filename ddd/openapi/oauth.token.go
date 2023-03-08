package openapi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
)

func token(ctx iris.Context) {

	var resp app.Response

	err := Srv.HandleTokenRequest(ctx.ResponseWriter(), ctx.Request())
	if err != nil {
		log.Info(err)

		tool.ResponseJSON(ctx, resp)
		return
	}

}
