package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"speed/filecontrol"
	"speed/models"
)

type JsonController struct {
	beego.Controller
}

func (this *JsonController) Get() {
	this.Ctx.WriteString("test the get method.")
}

func (this *JsonController) Post() {
	hi := models.HostInfo{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &hi)

	hi.ClientIP = this.Ctx.Input.IP()
	hi.ServerIp = this.Ctx.Request.Host

	filecontrol.RecordMsg(&hi)
	this.Data["json"] = &hi
	this.ServeJson()
	return
}
