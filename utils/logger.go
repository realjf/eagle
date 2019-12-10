package utils

import (
	"fmt"
	"github.com/gogf/gf/os/gfile"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const (
	LogDir = "../log"
	LogFile = LogDir + "/eagle"
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

		// 判断日志目录是否存在
		if !gfile.Exists(LogDir) {
			err = os.Mkdir(LogDir, 0755)
			if err != nil {
				panic(err)
			}
		}

		Logger.Out = src
		Logger.SetLevel(logrus.DebugLevel)

		logWriter, err := rotatelogs.New(
			LogFile+".%Y-%m-%d-%H-%M.log",
			rotatelogs.WithLinkName(LogFile),          // 生成软链，指向最新日志文件
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
