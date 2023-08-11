package oauth2

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/factory"
	token2 "github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/hufu/cmd/oauth/ddd/svr"
	"github.com/lishimeng/hufu/internal/db/model"
)

type Access struct {
	AccessToken string    `json:"accessToken,omitempty"`
	TokenType   TokenType `json:"tokenType,omitempty"`
	Scope       string    `json:"scope,omitempty"`
	ExpiresIn   int64     `json:"expiresIn,omitempty"`
}

type CredentialTokenResp struct {
	app.Response
	Access
}

type AccessTokenResp struct {
	app.Response
	Access
	RefreshToken string `json:"refreshToken,omitempty"`
}

func token(ctx iris.Context) {

	grantType := ctx.URLParam("grant_type")
	clientId := ctx.URLParam("client_id")
	clientSecret := ctx.URLParam("client_secret")

	code := ctx.URLParam("code")
	refresh := ctx.URLParam("refresh_token")
	switch GrantType(grantType) {

	case AuthorizationCode:
		log.Info("auth_code")
		accessToken(ctx, clientId, clientSecret, code)
	case ClientCredentials:
		log.Info("credential")
		credential(ctx, clientId, clientSecret)
	case Refreshing:
		log.Info("refresh_token")
		refreshToken(ctx, clientId, clientSecret, refresh)
	default:
		var resp app.Response
		resp.Code = iris.StatusBadRequest
		tool.ResponseJSON(ctx, resp)
		return
	}
}

func credential(ctx iris.Context, clientId, secret string) {

	var err error
	var resp CredentialTokenResp

	client, err := getCredentialClient(clientId)
	if err != nil {
		log.Debug(err)
		resp.Code = iris.StatusInternalServerError // TODO client not found
		tool.ResponseJSON(ctx, resp)
		return
	}
	if client.Secret != secret {
		log.Debug("secret not match")
		resp.Code = iris.StatusInternalServerError // TODO client not found
		tool.ResponseJSON(ctx, resp)
		return
	}
	// TODO gen token
	var provider token2.JwtProvider
	err = factory.Get(&provider)
	if err != nil {
		log.Debug(err)
		resp.Code = iris.StatusInternalServerError // TODO client not found
		tool.ResponseJSON(ctx, resp)
		return
	}
	ti, err := svr.GenCredentialToken(client, &provider)
	if err != nil {
		log.Debug(err)
		resp.Code = iris.StatusInternalServerError // TODO client not found
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.AccessToken = ti.AccessToken
	resp.ExpiresIn = ti.AccessTokenExpires
	resp.TokenType = Bearer
	resp.Code = iris.StatusOK
	tool.ResponseJSON(ctx, resp)
}

func accessToken(ctx iris.Context, clientId, secret, code string) {

}

func refreshToken(ctx iris.Context, clientId, secret, refresh string) {

}

func getCredentialClient(clientId string) (client model.CredentialsApp, err error) {
	err = app.GetOrm().Context.
		QueryTable(new(model.CredentialsApp)).
		Filter("AppId", clientId).
		One(&client)
	return
}
