package oauth2

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/hufu/internal/db/repo"
	"net/url"
)

func authorize(ctx iris.Context) {
	var err error
	var resp app.Response

	responseType := ctx.URLParam("response_type")
	clientId := ctx.URLParam("client_id")
	redirectUri := ctx.URLParam("redirect_uri")
	scope := ctx.URLParam("scope")
	log.Info("response_type:%s, client_id:%s, redirect:%s, scope:%s", responseType, clientId, redirectUri, scope)

	if !responseTypeSupport(ResponseType(responseType)) {

		resp.Code = iris.StatusBadRequest
		resp.Message = "unsupported response_type"
		tool.ResponseJSON(ctx, resp)
		return
	}

	ci, err := repo.GetClientByAppKey(clientId)
	if err != nil {

		resp.Code = iris.StatusForbidden
		resp.Message = "unknown client id"
		tool.ResponseJSON(ctx, resp)
	}
	log.Info("client id:%s", ci.AppId)
	code := tool.RandHexStr(4)

	redirect, err := buildRedirect(redirectUri, code)
	if err != nil {
		resp.Code = iris.StatusBadRequest
		resp.Message = "unsupported redirect_uri"
		tool.ResponseJSON(ctx, resp)
		return
	}

	err = storage.SaveCode(code, clientId)
	if err != nil {
		resp.Code = iris.StatusInternalServerError
		resp.Message = "internal error"
		tool.ResponseJSON(ctx, resp)
		return
	}

	ctx.Redirect(redirect)
}

func buildRedirect(redirect string, code string) (res string, err error) {
	d, err := url.Parse(redirect)
	if err != nil {
		return
	}
	query := d.Query()
	query.Add("code", code)
	q := query.Encode()
	d.RawQuery = q
	res = d.String()
	return
}

func responseTypeSupport(responseType ResponseType) bool {
	switch responseType {
	case Code:
		return true
	default:
		return false
	}
}
