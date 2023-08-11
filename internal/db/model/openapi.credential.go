package model

import "github.com/lishimeng/app-starter"

// CredentialsApp 凭证类App
type CredentialsApp struct {
	app.Pk
	OrgCode string `orm:"column(org_code)"`
	AppId   string `orm:"column(app_id);unique"`
	Secret  string `orm:"column(secret)"`
	Domain  string `orm:"column(domain)"`
	app.TableChangeInfo
}
