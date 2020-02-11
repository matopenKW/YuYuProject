package main

import (
	"YuYuProject/tool"
	"testing"
)

func TestInserSerial(t *testing.T) {

	err := tool.InsertSerial()

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

}
