package openapi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/go-log"
	"net/http"
)

func authorize(ctx iris.Context) {

	err := Srv.HandleAuthorizeRequest(ctx.ResponseWriter(), ctx.Request())
	if err != nil {
		log.Info(err)
		ctx.StopWithStatus(http.StatusBadRequest)
		return
	}
}
