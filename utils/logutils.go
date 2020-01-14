package utils

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

//GetLogger 获取自定义logger
func GetLogger(prefix string) (*os.File, *log.Logger) {
	currentTime := time.Now()
	filePath := currentTime.Format("2006/01/02")
	err := os.MkdirAll(fmt.Sprintf("./logs/%s", filePath), os.ModePerm)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, nil
	}
	fileName := fmt.Sprintf("./logs/%s.log", currentTime.Format("2006/01/02/15"))
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, nil
	}
	retValue := log.New(file, prefix, log.Ltime|log.LstdFlags|log.Lshortfile)
	return file, retValue
}

func stack() string {
	var buf [2 << 10]byte
	return string(buf[:runtime.Stack(buf[:], true)])
}

//Error 错误信息
func Error(msg string) {
	file, logger := GetLogger("[Error]")
	defer func() {
		if file != nil {
			file.Close()
		}
	}()
	logger.Printf("%s\n%s", msg, stack())
}

//Errorf 格式化错误信息
func Errorf(format string, a ...interface{}) {
	Error(fmt.Sprintf(format, a...))
}

//Warning 警告信息
func Warning(msg string) {
	file, logger := GetLogger("[Warning]")
	defer func() {
		if file != nil {
			file.Close()
		}
	}()
	logger.Println(msg)
}

//Info 提示信息
func Info(msg string) {
	file, logger := GetLogger("[Info]")
	defer func() {
		if file != nil {
			file.Close()
		}
	}()
	logger.Println(msg)
}
