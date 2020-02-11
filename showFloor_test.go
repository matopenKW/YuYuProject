package main

import (
	"YuYuProject/internal/apps"
	"testing"
)

func TestGetFloorData(t *testing.T) {

	form, err := apps.GetData("e_1")

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	t.Log(form)

}
