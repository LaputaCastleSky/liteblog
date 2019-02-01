package controllers

import (
	"github.com/astaxie/beego"
)

type NestPreparer interface {
	NestPrepare()
}

type BaseController struct {
	beego.Controller
}

func (ctx *BaseController) Prepare() {
	// 将页面路径 保存到 Path变量里面
	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI
	if app, ok := ctx.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}
