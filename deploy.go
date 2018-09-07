package main

import (
	"time"
	"os"
	"path/filepath"
	"log"
	"strings"
	"io/ioutil"
	"fmt"
)

type deployCmd struct {
	warPath    string
	at         string
	layout     string
	tomcatPath string
}

func (cmd *deployCmd) run() error {
	local, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation(cmd.layout, cmd.at, local)
	if err != nil {
		return err
	}
	d := t.Sub(time.Now())
	c := time.After(d)
	<-c
	return deploy(cmd.tomcatPath)
}

func deploy(tomcatPath string) error {

	deleteWebapp(tomcatPath)
	// call 1~5

	return nil
}

func stop() error {
	// 1. Shutdown Tomcat
	return nil
}

func deleteWebapp(tomcatPath string) error {
	// 2. 刪除 .../tomcat/webapps 目錄下的 payment.war 及 payment 資料夾

	//step1 進入變數tomcatPath下面的webapps目錄裡
	//step2 找到payment.war及payment資料夾
	//step3 刪除他們

	s := tomcatPath + "/webapps"

	files, err := ioutil.ReadDir(s)
	if err != nil{
		fmt.Print("could't find folder")
	}else{
		for _, f := range files {
			if f.Name() == "payment.war"{
				os.Remove(f.Name()) // 刪除檔案 (可多層)
			}
			if f.Name() == "payment"{
				os.RemoveAll(f.Name()) // 刪除資料料夾 (可多層)
			}
		}
	}

	return nil
}

func backupWar() error {
	// 3. 將 .../oldWar/payment.war 複製到 .../tomcat/webapps/
	return nil
}

func backupDb() error {
	// 4. 執行backupDb.bat 備份DB資料
	return nil
}

func start() error {
	// 5. Startup Tomcat
	return nil
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}