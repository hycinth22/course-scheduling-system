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
		if info.IsDir() && path != backupDir {
			return filepath.SkipDir
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

func GetBackupAbsPath(name string) string {
	return backupDir + name
}

func Backup() (string, error) {
	filename := time.Now().Format("2006_01_02_15_04_05.00000.sql")
	f, err := os.Create(filepath.Join(backupDir, filename))
	if err != nil {
		return filename, err
	}
	_, err = f.WriteString(ExportDB())
	if err != nil {
		return filename, err
	}
	return filename, nil
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

func RestoreDB(name string) error {
	sql, err := os.Open(GetBackupAbsPath(name))
	if err != nil {
		log.Println(err)
		return err
	}
	defer func(sql *os.File) {
		err := sql.Close()
		if err != nil {
			log.Println(err)
		}
	}(sql)
	var output bytes.Buffer
	cmd := exec.Command("D:\\bstool\\mariadb-10.5.9-winx64\\bin\\mysql", "-hlocalhost", "-uroot", "-proot", "coursescheduling")
	cmd.Dir = `D:\bstool\mariadb-10.5.9-winx64\bin`
	cmd.Stdin = sql
	cmd.Stdout = &output
	err = cmd.Run()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Restore using", name, output.String())
	return nil
}
