package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetAbout",
            Router: `/about`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetMessage",
            Router: `/message`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
