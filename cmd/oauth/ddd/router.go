package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/hufu/cmd/oauth/ddd/oauth2"
)

func Route(root *iris.Application) {

	group := root.Party("/")
	oauth2.Route(group)
}
