package controllers

import (
	"liteblog/models"
	"strings"

	"github.com/astaxie/beego"
)

type UserController struct {
	BaseController
}

// @router /login [post]
func (c *UserController) Login() {
	name := c.GetString("name")
	password := c.GetString("password")

	beego.Info("name:", name, "password:", password)

	if false == models.AuthUser(name, password) {
		c.Abort("name or password is invalid.")
	}

	// c.ServeJSON()
	c.Redirect("/", 301)
}

// @router /reg [post]
func (c *UserController) Reg() {
	name := c.GetString("name")
	email := c.GetString("email")
	pwd1 := c.GetString("password1")
	pwd2 := c.GetString("password2")
	if strings.Compare(pwd1, pwd2) != 0 {
		c.Abort("505")
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: pwd1,
		Avatar:   "",
		Role:     1,
	}

	if user.Exist() {
		c.Abort("user is already exist.")
	}

	user.Save()

	c.ServeJSON()
}
