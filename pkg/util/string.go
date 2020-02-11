package util

import (
	"errors"
)

func CheckNil(strs []string, errObjName string) error {
	if strs == nil || strs[0] == "" {
		return errors.New(errObjName + "が空です！")
	}

	return nil
}
