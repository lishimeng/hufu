package ddd

import (
	"github.com/kataras/iris/v12"
	openapi2 "github.com/lishimeng/oauth2/ddd/openapi"
)

func Route(root *iris.Application) {
	openapi2.Init()
	group := root.Party("/oauth2")
	openapi2.Router(group)
}
