package main

import (
	"YuYuProject/internal/dao"
	"YuYuProject/tool"
	"io/ioutil"
	"strings"
	"testing"
)

func TestInsertFloorList(t *testing.T) {

	dir := dao.JSON_FOLDER_PATH + "twins/"

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	for _, file := range files {
		path := file.Name()
		index := strings.Index(path, ".")
		tool.InsertFloorList(path[0:index])
	}

	t.Log("success")

}
