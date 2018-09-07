package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

func TestCode(t *testing.T) {
	fmt.Println("C:\\Users\\admin\\Desktop\\i519")
	fmt.Println(strings.Replace("C:\\Users\\admin\\Desktop\\i519", "\\", "_", -1))

}

func TestStop(t *testing.T) {

	// 產生臨時的檔案夾
	tmp, err := ioutil.TempDir(os.TempDir(), "testDeploy")
	//tmp, err := ioutil.TempDir("C:\\Users\\admin\\Desktop\\i519", "coppy")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("tmp = %v\n", tmp)

	// 完成測試後刪除檔案夾
	//defer os.RemoveAll(tmp)

	// 在臨時檔案夾中, 建立 testCoppy.bat 檔案。
	// bat內容：將自己(testCoppy.bat)另存新檔案為testCoppy_時分秒.bat
	batPath := path.Join(tmp, "testCoppy.bat")
	content := "@ECHO OFF\n"
	content += ("FOR /F \"tokens=1-4 delims=:.\" %%a IN (\"%time%\") DO (\n")
	content += ("set a=%%a\n")
	content += ("set b=%%b\n")
	content += ("set c=%%c\n")
	content += (")\n")
	content += ("CD " + tmp + "\n")
	content += ("ping 127.0.0.1 -n 10 -w 1000\n")
	content += ("COPY  testCoppy.bat " + tmp + "\\testCoppy_%a%%b%%c%.bat\n")
	content += ("ping 127.0.0.1 -n 60 -w 1000\n")
	content += ("ECHO finish")
	ioutil.WriteFile(batPath, []byte(content), os.ModePerm)
	fmt.Printf("batPath = %v\n", batPath)

	// 驗證
	err = stop(batPath)
	//err = stop("C:\\Users\\admin\\Desktop\\i519\\testCoppy1.bat")
	if err != nil {
		fmt.Printf("有錯！")
		t.Error(err)
	}
	err = stop(batPath)
	//err = stop("C:\\Users\\admin\\Desktop\\i519\\testCoppy1.bat")
	if err != nil {
		fmt.Printf("有錯！")
		t.Error(err)
	}

}

func TestBackupWar(t *testing.T) {
	/*
		// 產生臨時的檔案夾
		tmp, err := ioutil.TempDir(os.TempDir(), "backup-war")
		if err != nil {
			t.Error(err)
		}
		// 完成測試後刪除檔案夾
		defer os.RemoveAll(tmp)

		// 在臨時檔案夾中, 建立 myapp.war 檔案
		warPath := path.Join(tmp, "myapp.war")
		content := "hello"
		ioutil.WriteFile(warPath, []byte(content), os.ModePerm)

		// 執行邏輯
		backupWar()

		// 驗證
		expected := path.Join(tmp, "backup", "myapp.war")
		read, err := ioutil.ReadFile(expected)
		if err != nil {
			t.Error(err)
		}
		if actual := string(read); actual != content {
			t.Errorf("expected content '%s', but got '%s'", content, actual)
		}
	*/
}
