package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	connStr := "postgres://postgres:postgres@127.0.0.1/blog?sslmode=disable"
	orm.RegisterDataBase("default", "postgres", connStr)

	orm.RegisterModelWithPrefix("blog_", new(User))
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
}
