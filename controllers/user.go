package controllers

import (
	"liteblog/models"
	"strings"
)

type UserController struct {
	BaseController
}

// @router /login [post]
func (c *UserController) Login() {
	name := c.GetString("name")
	password := c.GetString("password")
	u, result := models.AuthUser(name, password)
	if result == false {
		c.Abort("name or password is invalid.")
	}

	c.SetSession(SESSION_USER_KEY, u)
	c.JSONOk("login ok.", "/")
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
	c.JSONOk("reg ok", "/user")
}
