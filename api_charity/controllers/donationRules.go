package controllers

import (
	"api_chraity/models"
	"fmt"

	"github.com/astaxie/beego"
)

// Operations donation rules
type DonationRulsController struct {
	beego.Controller
}

// @Title donation rules
// @Description input user  donation models and direction
// @Param	account			query 	string	true		"user account"
// @Param	model			query 	string	true		"donation model"
// @Param	Direction		query 	string	false		"donation to"
// @Success 200 {object} models.RcRules
// @Failure 403 body is empty
// @router / [post]
func (d *DonationRulsController) Post() {
	username := d.GetString("account")
	model := d.GetString("model")
	direction := d.GetString("direction")
	rr, err := models.DonationRulesUser(username, model, direction)
	if err != nil {
		d.Data["json"] = fmt.Sprintf("donation money error. %v", err)
	}
	d.Data["json"] = rr
	d.ServeJSON()
}
