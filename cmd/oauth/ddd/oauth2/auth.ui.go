package oauth2

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/oauth2/cmd/oauth/static"
	"net/http"
)

// auth 授权界面
func authUi(ctx iris.Context) {

	file, err := static.FS.Open("./auth.html")
	if err != nil {
		_, _ = ctx.HTML("%s", err.Error())
		return
	}
	fi, _ := file.Stat()
	http.ServeContent(ctx.ResponseWriter(), ctx.Request(), fi.Name(), fi.ModTime(), file)
}
