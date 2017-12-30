package wxDB

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {

	dburl := beego.AppConfig.String("dburl")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbName := beego.AppConfig.String("db")
	//注册mysql Driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//构造conn连接
	conn := dbuser + ":" + dbpassword + "@tcp(" + dburl + ")/" + dbName + "?charset=utf8"
	//注册数据库连接
	orm.RegisterDataBase("default", "mysql", conn)

	fmt.Printf("数据库连接成功！%s\n", conn)
}