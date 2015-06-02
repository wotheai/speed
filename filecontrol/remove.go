package filecontrol

import (
	"log"
	"os"
	"path"
)

func Del() {
	log.Println("begin to delete file.")
	func() {
		err := os.Remove(path.Join("attach", "upload"))
		if err != nil {
			log.Println(err)
			return
		} else {
			log.Println(path.Join("attach", "upload") + "remove file success.")
		}
	}()
}
