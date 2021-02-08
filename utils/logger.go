package utils

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// AppLogger 全局引用
var AppLogger = logrus.New()

var (
	// Trace 记录所有日志
	Trace *log.Logger

	// Info 重要的信息
	Info *log.Logger

	// Warning 需要注意的信息
	Warning *log.Logger

	// Error 非常严重的问题
	Error *log.Logger
)

/*
InitLogger 初始化logger --- 这个好像没用上的说
*/
func InitLogger(filePath string) {
	file, err := os.OpenFile(filePath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

/*
LogrusInit 全局Logger初始化
*/
func LogrusInit(path string) {
	customFormatter := new(logrus.JSONFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"

	AppLogger.SetFormatter(customFormatter)
	AppLogger.SetLevel(logrus.InfoLevel)
	AppLogger.SetReportCaller(true)

	/* 日志轮转相关函数
	    `WithLinkName` 为最新的日志建立软连接
	    `WithRotationTime` 设置日志分割的时间，隔多久分割一次
		`WithMaxAge` 设置文件清理前的最长保存时间
		`WithRotationCount` 设置文件清理前最多保存的个数
	    WithMaxAge 和 WithRotationCount二者只能设置一个
	*/
	writer, _ := rotatelogs.New(
		path+".%Y%m%d%",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		// rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		rotatelogs.WithRotationCount(7),
	)
	AppLogger.SetOutput(writer)
}
