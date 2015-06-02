package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"net/url"
	"os"
)

type ImgController struct {
	beego.Controller
}

func (this *ImgController) Get() {
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
