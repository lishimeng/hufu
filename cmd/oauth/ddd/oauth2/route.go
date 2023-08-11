package oauth2

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/oauth2/cmd/oauth/ddd/svr"
	"github.com/lishimeng/oauth2/internal/midware"
)

var storage svr.Storage

func Route(root iris.Party) {
	storage = svr.NewMemStorage()
	root.Get("/authorize", authorize)
	root.Get("/token", token)
	root.Get("/default_code_page", defaultCodePage)

	root.Get("/user_info", midware.WithAuth(userInfo)...)
	root.Get("/auth", authUi)

	root.Get("/validate", midware.WithAuth(checkToken)...)
}
