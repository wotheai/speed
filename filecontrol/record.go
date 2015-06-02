package filecontrol

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"speed/models"
)

func RecordMsg(h *models.HostInfo) {

	ws := fmt.Sprintln(h.ClientIP + "|" + h.ServerIp + "|" + h.Download + "|" + h.Upload)
	log := logs.NewLogger(10000)
	log.SetLogger("file", `{"filename":"logs/record.log"}`)
	log.Info(ws)
}
