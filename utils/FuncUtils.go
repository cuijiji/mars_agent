package utils

import (
	"fmt"
	"github.com/pkg/errors"
	osLog "log"
	"os"
	"path/filepath"
	"mars_agent/configs/log"
	"strings"
)

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Logger.Sugar().Error(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func HandleError(err error, message ...string) {
	stack := GetErrorStack(err)
	if log.Logger == nil {
		osLog.Panic(stack)
		return
	}
	if message == nil || len(message) == 0 {
		log.Logger.Sugar().Error(stack)
		return
	}
	log.Logger.Sugar().Error(message, stack)
}

func GetErrorStack(err error) string {
	return fmt.Sprintf("%+v", errors.WithStack(err))
}

func LocalFileExist(localPath string) (bool, error) {
	_, err := os.Stat(localPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
