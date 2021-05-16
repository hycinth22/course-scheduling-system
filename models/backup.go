package models

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const backupDir = "./backup/"

func ListBackup() []string {
	var files []string
	err := filepath.Walk(backupDir, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return nil
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".sql" {
			files = append(files, info.Name())
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return files
}

func Backup() error {
	filename := time.Now().Format("2006_01_02_15_04_05.00000.sql")
	f, err := os.Create(filepath.Join(backupDir, filename))
	if err != nil {
		return err
	}
	_, err = f.WriteString(ExportDB())
	if err != nil {
		return err
	}
	return nil
}

func ExportDB() string {
	var output bytes.Buffer
	cmd := exec.Command("D:\\bstool\\mariadb-10.5.9-winx64\\bin\\mysqldump", "--opt", "-hlocalhost", "-uroot", "-proot", "coursescheduling")
	cmd.Dir = `D:\bstool\mariadb-10.5.9-winx64\bin`
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		return ""
	}
	return output.String()
}
