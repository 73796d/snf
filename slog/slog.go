package slog

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	TraceLevel = iota
	DebugLevel
	WarnLevel
	ErrorLevel
	InfoLevel
	FatalLevel
)

var logger *log.Logger
var level = TraceLevel
var logFile *os.File
var isOutputScreen = true

// 获取日志级别
func GetLogLevel() int {
	return level
}

// 设置日志级别
func SetLogLevel(l int) {
	if l > FatalLevel || l < TraceLevel {
		level = TraceLevel
	} else {
		level = l
	}
}

// 设置日志是否输出到屏幕
func SetIsOutputScreen(isOutput bool) {
	isOutputScreen = isOutput
}

// 初始化日志
func InitLog(logFileName string) {
	var err error
	pid := os.Getgid()
	pidStr := strconv.FormatInt(int64(pid), 10)
	logFileName = "log/" + logFileName + "_" + pidStr + ".log"
	if err:= os.MkdirAll("log", 0666); err != nil {
		fmt.Println(err.Error())
		return
	}
	logFile, err = os.OpenFile(logFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	logger.Println("log to file sample")
}

// 跟踪级别的日志
func Trace(format string, v ...interface{}) {
	if level <= TraceLevel {
		var str string = "[T]" + format
		str = fmt.Sprintf(str, v...)
		logger.Output(2, str)
		if isOutputScreen {
			fmt.Println(str)
		}
	}
}

func Debug(format string, v ...interface{}) {
	if level <= DebugLevel {
		var str string
		str = "[D] " + format
		str = fmt.Sprintf(str, v...)
		logger.Output(2, str)

		if isOutputScreen {
			fmt.Println(str)
		}
	}
}
func Warn(format string, v ...interface{}) {
	if level <= WarnLevel {
		var str string
		str = "[W] " + format
		str = fmt.Sprintf(str, v...)
		logger.Output(2, str)

		if isOutputScreen {
			fmt.Println(str)
		}
	}
}

func Error(format string, v ...interface{}) {
	if level <= ErrorLevel {
		var str string
		str = "[E] " + format
		str = fmt.Sprintf(str, v...)
		logger.Output(2, str)

		if isOutputScreen {
			fmt.Println(str)
		}
	}
}

func Info(format string, v ...interface{}) {
	if level <= InfoLevel {
		var str string
		str = "[I] " + format
		str = fmt.Sprintf(str, v...)
		logger.Output(2, str)

		if isOutputScreen {
			fmt.Println(str)
		}
	}
}

func Fatal(format string, v ...interface{}) {
	if level <= FatalLevel {
		var str string
		str = "[F] " + format
		str = fmt.Sprintf(str, v...)
		logger.Output(2, str)

		if isOutputScreen {
			fmt.Println(str)
		}
	}
}
