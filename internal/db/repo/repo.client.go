package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/oauth2/internal/db/model"
)

func GetClientByAppKey(appKey string) (ci model.OpenApp, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.OpenApp)).Filter("AppId", appKey).One(&ci)
	return
}
