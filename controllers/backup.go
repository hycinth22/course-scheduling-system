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

// @router / [get]
func (this *BackupController) List() {
	list := models.ListBackup()
	this.Data["json"] = list
	err := this.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @router /new [post]
func (this *BackupController) Create() {
	err := models.Backup()
	if err != nil {
		log.Println(err)
		return
	}
}

// @router /:name [delete]
func (this *BackupController) Delete() {
	name := this.GetString("name")
	path := models.GetBackupAbsPath(name)
	newpath := models.GetBackupAbsPath(name)
	err := os.Rename(path, newpath)
	if err != nil {
		log.Println(err)
		return
	}
}

// @router /new/upload [post]
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

// @router /download/:name [put]
func (this *BackupController) Download() {
	name := this.GetString("name")
	this.Ctx.Output.Download(models.GetBackupAbsPath(name))
}
