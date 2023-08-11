package oauth2

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
)

func checkToken(ctx iris.Context) {
	var resp app.Response

	resp.Code = iris.StatusOK
	tool.ResponseJSON(ctx, resp)
}
