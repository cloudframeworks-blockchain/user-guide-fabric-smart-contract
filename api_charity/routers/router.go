// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"api_charity/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/donation",
			beego.NSInclude(
				&controllers.DonationController{},
			),
		),
		beego.NSNamespace("/querydealonce",
			beego.NSInclude(
				&controllers.QueryOnceController{},
			),
		),
		beego.NSNamespace("/querydealall",
			beego.NSInclude(
				&controllers.QueryALLController{},
			),
		),
		beego.NSNamespace("/queryuserinfo",
			beego.NSInclude(
				&controllers.UserInfoController{},
			),
		),
		beego.NSNamespace("/donationRules",
			beego.NSInclude(
				&controllers.DonationRulsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
