package controllers

import (
	"io"
	"log"
	"os"

	"courseScheduling/models"
	beego "github.com/beego/beego/v2/server/web"
)

type BackupController struct {
	beego.Controller
}

// @router /list [get]
func (this *BackupController) List() {
	list := models.ListBackup()
	if list == nil {
		list = []string{}
	}
	this.Data["json"] = list
	err := this.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @router /new [post]
func (this *BackupController) Create() {
	filename, err := models.Backup()
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(500)
		return
	}
	this.Data["json"] = map[string]interface{}{
		"filename": filename,
	}
	err = this.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @router /:name [delete]
func (this *BackupController) Delete() {
	name := this.GetString(":name")
	path := models.GetBackupAbsPath(name)
	newpath := models.GetBackupAbsPath("deleted/" + name)
	log.Println(path, newpath)
	err := os.Rename(path, newpath)
	if err != nil {
		log.Println(err)
		return
	}
}

// @router /upload [post]
func (this *BackupController) Upload() {
	closeIO := func(f io.Closer) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}
	uploaded, header, err := this.GetFile("backupFile")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer closeIO(uploaded)
	created, err := os.Create(models.GetBackupAbsPath(header.Filename))
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(500)
		return
	}
	_, err = io.Copy(created, uploaded)
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(500)
		return
	}
	defer closeIO(created)
}

// @router /download/ [get]
func (this *BackupController) Download() {
	name := this.GetString("name")
	log.Println("download", name)
	filename := models.GetBackupAbsPath(name)
	log.Println(name, filename)
	this.Ctx.Output.Download(filename)
}

// @router /apply [put]
func (this *BackupController) Restore() {
	name := this.GetString("name")
	log.Println("Restore", name)
	err := models.RestoreDB(name)
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(500)
		return
	}
}
