package openapi

import (
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/oauth2/internal/db/model"
)

func GetClientByAppKey(ctx persistence.OrmContext, appKey string) (ci model.OpenClient, err error) {
	err = ctx.Context.QueryTable(new(model.OpenClient)).Filter("AppId", appKey).One(&ci)
	return
}
