package svr

import (
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/oauth2/internal/db/model"
	"time"
)

var (
	AccessTokenExpires int64 = 60 * 60 * 24 * 30
	RefreshToken       int64 = 24 * 30
)

type TokenInfo struct {
	AccessToken         string
	RefreshToken        string
	AccessTokenExpires  int64
	RefreshTokenExpires int64
}

func GenCredentialToken(data model.CredentialsApp, provider *token.JwtProvider) (ti TokenInfo, err error) {
	payload := token.JwtPayload{
		Org: data.OrgCode,
		Uid: data.AppId,
	}
	expireIn := time.Second * time.Duration(AccessTokenExpires)
	t, err := provider.GenWithTTL(payload, expireIn)
	if err != nil {
		return
	}
	ti.AccessToken = string(t)
	ti.AccessTokenExpires = AccessTokenExpires
	return
}
