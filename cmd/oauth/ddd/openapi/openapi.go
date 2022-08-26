package openapi

import (
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/oauth2/internal/etc"
	"net/http"
	"sync"
)
import (
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	oredis "github.com/go-oauth2/redis/v4"
)

type Req struct {
	AccessToken  string
	RefreshToken string
	AppKey       string
	AppSecret    string
}

var (
	Srv  *server.Server
	once sync.Once
)

func Init() {

	manager := manage.NewDefaultManager()
	clientStore := NewClientStore()

	manager.MapClientStorage(clientStore)
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate(
		"",
		[]byte("00000000"),
		jwt.SigningMethodHS512),
	)

	if etc.Config.Redis.Enable {
		manager.MustTokenStorage(oredis.NewRedisStore(&redis.Options{
			Addr: etc.Config.Redis.Addr,
		}), nil)
	} else {
		manager.MustTokenStorage(store.NewMemoryTokenStore())
	}

	once.Do(func() {
		Srv = server.NewDefaultServer(manager)
		Srv.SetAllowGetAccessRequest(true)
		Srv.SetClientInfoHandler(server.ClientFormHandler)
		Srv.SetAuthorizeScopeHandler(func(w http.ResponseWriter, r *http.Request) (scope string, err error) {
			return
		})
		Srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
			log.Info("Internal Error:%s", err.Error())
			return
		})
		//Srv.SetPasswordAuthorizationHandler(func(ctx context.Context, clientID, username, password string) (userID string, err error) {
		//	err = errors.New("401")
		//	return
		//})
		Srv.SetUserAuthorizationHandler(func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
			userID = r.Header.Get("UID")
			if len(userID) == 0 {
				userID = r.Header.Get("User")
			}
			return
		})
	})

}
