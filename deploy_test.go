package main

import (
	"testing"
	"io/ioutil"
	"os"
	"path"
)

func TestStop(t *testing.T) {
	err := stop()
	if err != nil {
		t.Error(err)
	}
}

func TestBackupWar(t *testing.T) {
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
}
