package main

import (
	_ "liteblog/routers"
	"strings"

	"github.com/astaxie/beego"
)

func main() {
	initTemplate()
	beego.Run()
}

func initTemplate() {
	beego.AddFuncMap("equrl", func(x, y string) bool {
		s1 := strings.Trim(x, "/")
		s2 := strings.Trim(y, "/")
		return strings.Compare(s1, s2) == 0
	})
}
