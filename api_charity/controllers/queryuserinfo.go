package controllers

import (
	"api_charity/models"
	"fmt"

	"github.com/astaxie/beego"
)

// Operations show user s info
type UserInfoController struct {
	beego.Controller
}

// @Title Show user info
// @Description show user info
// @Param	account		query 	string	true		"user account"
// @Success 200 {object} models.CharityUser
// @Failure 403 query args is empty
// @router / [post]
func (u *UserInfoController) Post() {
	username := u.GetString("account")
	cUser, err := models.GetUserInfo(username)
	if err != nil {
		u.Data["json"] = fmt.Sprintf("Get user s info error. %v", err)
	}
	u.Data["json"] = cUser
	u.ServeJSON()
}
