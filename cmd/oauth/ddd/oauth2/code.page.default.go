package oauth2

import (
	"github.com/kataras/iris/v12"
)

func defaultCodePage(ctx iris.Context) {

	codeValue := ctx.URLParam("code")
	_, _ = ctx.HTML(`<p>code:%s</p>`, codeValue)
}
