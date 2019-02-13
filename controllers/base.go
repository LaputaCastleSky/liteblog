package controllers

import (
	"errors"
	"liteblog/models"

	"github.com/astaxie/beego"
)

type NestPreparer interface {
	NestPrepare()
}

type BaseController struct {
	beego.Controller
	IsLogin bool
	User    models.User
}

const SESSION_USER_KEY = "SESSION_USER_KEY"

func (ctx *BaseController) Prepare() {
	// 将页面路径 保存到 Path变量里面
	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI
	ctx.IsLogin = false
	tu := ctx.GetSession(SESSION_USER_KEY)
	if tu != nil {
		if u, ok := tu.(models.User); ok {
			ctx.User = u
			ctx.Data["User"] = u
			ctx.IsLogin = true
		}
	}
	ctx.Data["IsLogin"] = ctx.IsLogin
	if app, ok := ctx.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

func (ctx *BaseController) MustLogin() {
	if !ctx.IsLogin {
		ctx.Abort500(errors.New("must login"))
	}
}

func (c *BaseController) GetMustString(key string, msg string) string {
	v := c.GetString(key, "")
	if len(v) == 0 {
		c.Abort500(errors.New(msg))
	}
	return v
}

func (ctx *BaseController) Abort500(err error) {
	ctx.Data["error"] = err
	ctx.Abort("500")
}

type ResultJsonValue struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Action string      `json:"action,omitempty"`
	Count  int         `json:"count,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func (ctx *BaseController) JSONOk(msg string, actions ...string) {
	var action string
	if len(actions) > 0 {
		action = actions[0]
	}
	ctx.Data["json"] = &ResultJsonValue{
		Code:   0,
		Msg:    msg,
		Action: action,
	}
	ctx.ServeJSON()
}
