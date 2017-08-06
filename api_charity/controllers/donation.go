package controllers

import (
	"api_chraity/models"
	"fmt"

	"github.com/astaxie/beego"
)

// Operations about Donation
type DonationController struct {
	beego.Controller
}

// @Title donationMoneyFromUser
// @Description input user and account
// @Param	account		query 	models.Donation	true		"user account"
// @Param	money		query 	models.money	true		"donation money"
// @Success 200 {object} models.CharityUser
// @Failure 403 Donation user account error
// @router / [post]
func (d *DonationController) Post() {
	username := d.GetString("account")
	money := d.GetString("money")
	cUser, err := models.DonationUser(username, money)
	if err != nil {
		d.Data["json"] = fmt.Sprintf("Donation user account error. %v", err)
	}
	//d.Data["json"] = "User account build success."
	d.Data["json"] = cUser
	d.ServeJSON()
}
