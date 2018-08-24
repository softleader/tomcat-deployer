package main

import (
	"fmt"
	"time"
)

type deployCmd struct {
	warPath string
	at      string
	layout  string
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

	// 1. Shutdown Tomcat
	// 2. 刪除 .../tomcat/webapps 目錄下的 payment.war 及 payment 資料夾
	// 3. 將 .../oldWar/payment.war 複製到 .../tomcat/webapps/
	// 4. 執行backupDb.bat 備份DB資料
	// 5. Startup Tomcat

	fmt.Println("!")
	return nil
}
