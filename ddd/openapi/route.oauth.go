package openapi

import "github.com/kataras/iris/v12"

func Router(root iris.Party) {
	root.Get("/authorize", authorize)
	root.Get("/token", token)
}
