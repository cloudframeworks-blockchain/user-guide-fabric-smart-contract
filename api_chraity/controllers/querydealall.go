package controllers

import (
	"api_chraity/models"
	"fmt"

	"github.com/astaxie/beego"
)

// Query user s all record
type QueryALLController struct {
	beego.Controller
}

// @Title Query all of user's transaction record
// @Description input username
// @Param	account		query 	string	true		"user account"
// @Success 200 {object} models.QueryALLRecords
// @Failure 403 query args is empty
// @router / [post]
func (q *QueryALLController) Post() {
	username := q.GetString("account")
	rs, err := models.QueryALLRecords(username)
	if err != nil {
		q.Data["json"] = fmt.Sprintf("Query user s records Error. %v", err)
	}
	q.Data["json"] = rs
	q.ServeJSON()
}
