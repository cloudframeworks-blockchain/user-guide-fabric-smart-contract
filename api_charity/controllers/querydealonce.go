package controllers

import (
	"api_charity/models"
	"fmt"

	"github.com/astaxie/beego"
)

// Operations query once
type QueryOnceController struct {
	beego.Controller
}

// @Title query user s once record
// @Description input user and record number
// @Param	account		query 	string	true		"user account"
// @Param	nums		query 	string	true		"record numbs"
// @Success 200 {object} models.CharityNote
// @Failure 403 args is empty
// @router / [post]
func (q *QueryOnceController) Post() {
	username := q.GetString("account")
	nums := q.GetString("nums")
	cNote, err := models.QueryOnce(username, nums)
	if err != nil {
		q.Data["json"] = fmt.Sprintf("Query users record error. %v", err)
	}
	q.Data["json"] = cNote
	q.ServeJSON()
}
