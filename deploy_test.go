package main

import (
	"testing"
	"os"
	"io/ioutil"
	"path"
	"fmt"
)

func TestStop(t *testing.T) {
	err := stop()
	if err != nil {
		t.Error(err)
	}
}

//func TestBackupWar(t *testing.T) {
//	// 產生臨時的檔案夾
//	tmp, err := ioutil.TempDir(os.TempDir(), "backup-war")
//	if err != nil {
//		t.Error(err)
//	}
//	// 完成測試後刪除檔案夾
//	defer os.RemoveAll(tmp)
//
//	// 在臨時檔案夾中, 建立 myapp.war 檔案
//	warPath := path.Join(tmp, "myapp.war")
//	content := "hello"
//	ioutil.WriteFile(warPath, []byte(content), os.ModePerm)
//
//	// 執行邏輯
//	backupWar()
//
//	// 驗證
//	expected := path.Join(tmp, "backup", "myapp.war")
//	read, err := ioutil.ReadFile(expected)
//	if err != nil {
//		t.Error(err)
//	}
//	if actual := string(read); actual != content {
//		t.Errorf("expected content '%s', but got '%s'", content, actual)
//	}
//}

func TestDeleteFile(t *testing.T)  {

	//在指定的資料夾下(測試時是tmp,正式的時候是tomcatPath)產生webapps資料夾
	tmp, err := ioutil.TempDir(os.TempDir(), "webapps")
	os.Mkdir("webapps", os.ModePerm); // 當前目錄建立 A 資料夾

	if err != nil {
		t.Error(err)
	}

	//在資料夾(路徑+webapps)中產生payment.war及payment資料夾
	warPath := path.Join(tmp, "payment.war")
	content := "helloPayment"
	ioutil.WriteFile(warPath, []byte(content), os.ModePerm)

	filePath := path.Join(tmp, "payment")
	ioutil.WriteFile(filePath, []byte(content), os.ModePerm)

	//執行邏輯
	deleteWebapp(tmp)

	//驗證
	files, err := ioutil.ReadDir(tmp)
	if err != nil{
		fmt.Print("could't find folder")
	}else{
		for _, f := range files {
			if f.Name() == "payment.war"{
				t.Error("war still in there")
			}else{
				fmt.Print("war file deleted!")
			}
			if f.Name() == "payment"{
				t.Error("folder still in there")
			}else{
				fmt.Print("folder file deleted!")
			}
		}
	}

	//完成後刪除資料夾
	defer os.RemoveAll(tmp)
}

