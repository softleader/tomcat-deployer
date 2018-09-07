package main

import (
	"fmt"
	"os/exec"
	"time"
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
	return deploy(cmd)
}

func deploy(cmd *deployCmd) error {
	// call 1~5

	// 1. Shutdown Tomcat
	if err := stop(cmd.tomcatPath); err != nil {
		fmt.Printf("stop error")
		return err
	}

	// 2. 刪除 .../tomcat/webapps 目錄下的 payment.war 及 payment 資料夾
	if err := deleteWebapp(cmd.tomcatPath + "\\webapps"); err != nil {
		fmt.Printf("deleteWebapp error")
		return err
	}

	// 3. 將 .../oldWar/payment.war 複製到 .../tomcat/webapps/
	if err := backupWar(); err != nil {
		fmt.Printf("backupWar error")
		return err
	}

	// 4. 執行backupDb.bat 備份DB資料
	if err := backupDb(); err != nil {
		fmt.Printf("backupDb error")
		return err
	}

	// 5. Startup Tomcat
	if err := start(); err != nil {
		fmt.Printf("start error")
		return err
	}
}

func stop(absFilePath string) error {
	// 1. Shutdown Tomcat

	fmt.Println("執行Shutdown path[" + absFilePath + "]")
	//cmd := exec.Command("cmd.exe", "/C", "echo e")
	//cmd := exec.Command("cmd.exe", "/C", "C:/D/act.txt")
	cmd := exec.Command("cmd.exe", "/C", absFilePath)
	d, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println("這是甚麼？" + string(d))
	return nil
}

func deleteWebapp(absFilePath string) error {
	// 2. 刪除 .../tomcat/webapps 目錄下的 payment.war 及 payment 資料夾
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
