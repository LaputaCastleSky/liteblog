package controllers


type IndexController struct {
	BaseController
}

// @router / [get]
func (c *IndexController) Get() {
	c.TplName = "index.html"
}

// @router /message [get]
func (c *IndexController) GetMessage() {
	c.TplName = "message.html"
}

// @router /about [get]
func (c *IndexController) GetAbout() {
	c.TplName = "about.html"
}

// @router /user [get]
func (c *IndexController) User() {
	c.TplName = "user.html"
}

// @router /reg [get]
func (c *IndexController) Reg() {
	c.TplName = "reg.html"
}
