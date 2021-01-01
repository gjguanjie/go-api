package routers

import (
	"github.com/astaxie/beego"
	"go-api/controllers"
)

func init()  {
	beego.Router(
		"/queryPerson",
		&controllers.PersonController{},
		"post:QueryPerson",
	)

}
