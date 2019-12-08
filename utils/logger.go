package utils

import (
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const (
	LogPath = "../log/eagle"
)

var Logger *logrus.Logger

func init() {
	if Logger == nil {
		Logger = logrus.New()
		// 禁止logrus的输出
		src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("err: ", err)
		}

		Logger.Out = src
		Logger.SetLevel(logrus.DebugLevel)

		logWriter, err := rotatelogs.New(
			LogPath+".%Y-%m-%d-%H-%M.log",
			rotatelogs.WithLinkName(LogPath),          // 生成软链，指向最新日志文件
			rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
			rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
		)
		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.WarnLevel: logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.FatalLevel: logWriter,
		}

		lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
		Logger.AddHook(lfHook)
	}
}
