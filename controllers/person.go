package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"go-api/models"
)

var (
	personList map[string]*models.Person
)



type PersonController struct {
	beego.Controller
}

func (this *PersonController) QueryPerson()  {
	var person *models.Person
	json.Unmarshal(this.Ctx.Input.RequestBody,&person)
	person.Name = person.Name + "11"
	this.Data["json"] = person
	person.Age = 1000
	this.ServeJSON()
}

// @router /test
func (this *PersonController) test()  {
	this.Ctx.WriteString("000000")
}


