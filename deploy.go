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

	// 1. 將新的 payment.war 備份到 .../oldWar 資料夾
	// 2. Shutdown Tomcat
	// 3. 刪除 .../tomcat/webapps 目錄下的 payment.war 及 payment 資料夾
	// 4. 將 .../oldWar/payment.war 複製到 .../tomcat/webapps/
	// 5. 執行backupDb.bat 備份DB資料
	// 6. Startup Tomcat
	// 7. 完成

	local, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}
	tt, err := time.ParseInLocation(cmd.layout, cmd.at, local)
	if err != nil {
		return err
	}
	d := tt.Sub(time.Now())
	t := time.After(d)
	<-t
	return deploy()
}

func deploy() error {
	fmt.Println("!")
}
