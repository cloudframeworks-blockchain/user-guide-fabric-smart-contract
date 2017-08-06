package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["api_chraity/controllers:DonationController"] = append(beego.GlobalControllerRouter["api_chraity/controllers:DonationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_chraity/controllers:DonationRulsController"] = append(beego.GlobalControllerRouter["api_chraity/controllers:DonationRulsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_chraity/controllers:QueryALLController"] = append(beego.GlobalControllerRouter["api_chraity/controllers:QueryALLController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_chraity/controllers:QueryOnceController"] = append(beego.GlobalControllerRouter["api_chraity/controllers:QueryOnceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_chraity/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["api_chraity/controllers:UserInfoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
