package logm

import (
	"fmt"
	"github.com/han-joker/moo-layout/moo/toolm"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

//常量

const file_ext = ".log"
//日志格式
const (
	TEXT = iota
	JSON
)
//存储类型
const (
	CONSOLE = iota
	FILE
	USER
)

//类型
//日志
type log struct {
	//选项
	opt Opt

	// cache
	cacheFilename string
	cacheFile *os.File
}
type Opt struct {
	Fmt int // TEXT, JSON
	Caller bool // true, false
	Mode int // CONSOLE, FILE, USER
	Path string // some path
	Filename string // %y, %m, %d, %w
	Out io.Writer
}

//变量
//日志单例
var instance *log

var fmtContainer = []int{TEXT, JSON}
var modeContainer = []int{CONSOLE, FILE, USER}

//默认选项
var dftOpt = Opt{
	Fmt: TEXT,
	Caller: false,
	Mode: CONSOLE,
	Path: "./logs/",
	Filename: "moo-%y-%m-%d",
	Out: os.Stdout,
}

// Inst 配置对象单例工厂
func Inst(opt ...Opt) *log {
	if instance == nil {
		//创建实例
		instance = &log{
			opt: dftOpt,
		}
		//设置选项
		if len(opt) > 0 {
			opt := opt[0]
			//基于用户选项，设置实例
			if toolm.IntSliceContains(opt.Fmt, fmtContainer) {
				instance.opt.Fmt = opt.Fmt
			}

			instance.opt.Caller = opt.Caller

			if toolm.IntSliceContains(opt.Mode, modeContainer) {
				instance.opt.Mode = opt.Mode
				switch opt.Mode {
				case CONSOLE:
					instance.opt.Out = os.Stdout
				case FILE:
					instance.opt.Path = path.Clean(toolm.StringDefault(opt.Path, dftOpt.Path))
					instance.opt.Filename = toolm.StringDefault(opt.Filename, dftOpt.Filename)
					err := os.MkdirAll(instance.opt.Path, 0644)
					if err != nil && !os.IsExist(err) {
						// 存在错误，但不是目录存在，创建日志目录失败，改为 CONSOLE 行为
						logrus.Info("创建日志目录失败，调整为 Console 模式")
						instance.opt.Mode = CONSOLE
						instance.opt.Path = dftOpt.Path
						instance.opt.Filename = dftOpt.Filename
					}
				case USER:
				}
			}
			if opt.Out != nil {
				if _, ok := opt.Out.(io.Writer); ok {
					instance.opt.Out = opt.Out
				}
			}
		}
		// 基于选项配置 logrus
		switch instance.opt.Fmt {
		case JSON:
			logrus.SetFormatter(&logrus.JSONFormatter{})
		case TEXT:
			logrus.SetFormatter(&logrus.TextFormatter{})
		}

		logrus.SetReportCaller(instance.opt.Caller)

	}
	return instance
}

//可导出方法
func (log) Info(args ...interface{}) {
	if instance.opt.Mode == FILE {
		now := time.Now()
		filename := instance.opt.Filename
		filename = strings.ReplaceAll(filename, "%y", fmt.Sprintf("%d", now.Year()))
		filename = strings.ReplaceAll(filename, "%m",  fmt.Sprintf("%d", now.Month()))
		filename = strings.ReplaceAll(filename, "%d",  fmt.Sprintf("%d", now.Day()))
		filename = strings.ReplaceAll(filename, "%w",  fmt.Sprintf("%d", now.Weekday()))

		// 日志文件句柄未缓存
		if filename != instance.cacheFilename {
			// 缓存上
			instance.cacheFilename = filename
			// 删除旧句柄，使用新句柄
			if instance.cacheFile != nil {
				instance.cacheFile.Close()
			}
			file, err := os.OpenFile(instance.opt.Path + "/" + filename + file_ext, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err == nil {
				instance.cacheFile = file
				logrus.SetOutput(file)
			} else {
				logrus.SetOutput(dftOpt.Out)
				logrus.Info("Failed to log to file, using default")
			}

		}
	}
	logrus.Info(args...)
}