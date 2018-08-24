package main

import (
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
	return deploy()
}

func deploy() error {

	// call 1~5

	return nil
}

func stop() error {
	// 1. Shutdown Tomcat
	return nil
}

func deleteWebapp() error {
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
