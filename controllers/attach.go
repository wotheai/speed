package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"net/url"
	"os"
	"path"
	"speed/filecontrol"
)

type AttachController struct {
	beego.Controller
}

func (this *AttachController) Download() {
	filePath, err := url.QueryUnescape(this.Ctx.Request.RequestURI[1:])
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	fmt.Println(filePath)

	f, err := os.Open(filePath)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	defer f.Close()

	_, err = io.Copy(this.Ctx.ResponseWriter, f)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	return
}

func (this *AttachController) Upload() {
	filecontrol.Del()

	_, fh, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}

	if fh != nil {
		err = this.SaveToFile("attachment", path.Join("attach", "upload"))
		if err != nil {
			beego.Error(err)
		}
	}

	this.Ctx.WriteString("上传成功")

	return
}

func (this *AttachController) Add() {
	this.TplNames = "topic_add.html"
}
