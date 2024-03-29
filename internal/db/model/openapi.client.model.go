package model

import "github.com/lishimeng/app-starter"

// OpenApp 个人类App
type OpenApp struct {
	app.Pk
	OrgCode string `orm:"column(org_code)"`
	AppId   string `orm:"column(app_id);unique"`
	Secret  string `orm:"column(secret)"`
	Domain  string `orm:"column(domain)"`
	UserId  string `orm:"column(user_id);null"`
	app.TableChangeInfo
}
