package controllers

import "fmt"

type IndexController struct {
	BaseController
}

//首页
// @router / [get]
func (c *IndexController) Get() {
	c.TplName = "index.html"
}

//留言
// @router /message [get]
func (c *IndexController) GetMessage() {
	fmt.Println("aaaa")
	c.TplName = "message.html"
}

//关于
// @router /about [get]
func (c *IndexController) GetAbout() {
	c.TplName = "about.html"
}
