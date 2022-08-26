package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/oauth2/cmd/oauth/ddd/openapi"
)

func Route(root *iris.Application) {
	openapi.Init()
	group := root.Party("/oauth2")
	openapi.Router(group)
}
