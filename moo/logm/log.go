package logm

import (
	"fmt"
	"github.com/han-joker/moo-layout/moo/toolm"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"time"
)

//常量

//日志文件后缀
const fileExt = ".log"

//日志格式
const (
	Text = iota
	Json
)

// OutMode 预设值
const (
	Console      = iota // 控制台
	File                // 文件，单文件模式
	FilePerDay          // 每天一个文件
	FilePerWeek         // 每周一个文件
	FilePerMonth        // 每月一个文件
	FilePerHour         // 每小时一个文件
	FilePerSize         // 文件固定大小
	User
)

//变量

//池
var pool = map[string]*log{}

//格式集合
var fmtContainer = []int{Text, Json}

//输出模式
var outModeContainer = []int{Console, File, FilePerDay, FilePerWeek, FilePerMonth, FilePerHour, FilePerSize, User}

//默认选项
var optionDefault = Option{
	Fmt:        Text,
	Caller:     false,
	OutMode:    Console,
	Path:       "./logs/",
	FilePrefix: "moo",
	SizeMax:    100 * 1024 * 1024, // 100 M Bytes
	Output:     os.Stdout,
}

//类型

//日志
type log struct {
	*logrus.Logger
	option Option
	// cache
	writers map[string]*os.File
}

// Option 选项
type Option struct {
	Name string

	Fmt        int    // Text, Json
	Caller     bool   // true, false
	OutMode    int    // Console, FILE, User
	Path       string // some path
	FilePrefix string //
	SizeMax    int64
	Output     io.Writer
}

type Fields = logrus.Fields

// New 创建对象
func New(option ...Option) *log {
	verifiedOption := optionVerify(option...)
	l := create(verifiedOption)
	l.refreshOutMode()
	return l
}

// Get 存在直接返回，否则创建、存储再返回
func Get(option ...Option) *log {
	verifiedOption := optionVerify(option...)
	if !Has(verifiedOption.Name) {
		pool[verifiedOption.Name] = create(verifiedOption)
	}
	pool[verifiedOption.Name].refreshOutMode()
	return pool[verifiedOption.Name]
}

// Has 存在返回 true，否则返回 false
func Has(name string) bool {
	_, has := pool[name]
	return has
}


func (l *log) syncLoggerOption() *log {
	switch l.option.Fmt {
	case Json:
		l.Logger.SetFormatter(&logrus.JSONFormatter{})
	case Text:
		l.Logger.SetFormatter(&logrus.TextFormatter{})
	default:
		l.Logger.SetFormatter(&logrus.TextFormatter{})
	}

	l.Logger.SetReportCaller(l.option.Caller)

	l.Logger.SetOutput(l.option.Output)

	return l
}
func (l *log) refreshOutMode() *log {
	// 初始默认值
	l.Logger.SetOutput(optionDefault.Output)
	filename := ""
	switch l.option.OutMode {
	case File:
		filename = l.option.Path + "/" + l.option.FilePrefix + fileExt
	case FilePerHour:
		now := time.Now()
		filename = l.option.Path + "/" +
			l.option.FilePrefix +
			fmt.Sprintf("%04d-%02d-%02d-%02d", now.Year(), now.Month(), now.Day(), now.Hour()) +
			fileExt
	case FilePerDay:
		now := time.Now()
		filename = l.option.Path + "/" +
			l.option.FilePrefix +
			fmt.Sprintf("%04d-%02d-%02d", now.Year(), now.Month(), now.Day()) +
			fileExt
	case FilePerMonth:
		now := time.Now()
		filename = l.option.Path + "/" +
			l.option.FilePrefix +
			fmt.Sprintf("%04d-%02d", now.Year(), now.Month()) +
			fileExt
	case FilePerWeek:
		now := time.Now()
		year, week := now.ISOWeek()
		filename = l.option.Path + "/" +
			l.option.FilePrefix +
			fmt.Sprintf("%04d-%02d", year, week) +
			fileExt
	case FilePerSize:
		filename = l.option.Path + "/" + l.option.FilePrefix + fileExt
		if fileinfo, err := os.Stat(filename); err == nil {
			if fileinfo.Size() >= l.option.SizeMax {
				basename := path.Base(filename)
				if file, exists := l.writers[basename]; exists {
					if err := file.Close(); err != nil {
						l.Info("can not close file")
					} else {
						delete(l.writers, basename)
					}
				}
				now := time.Now()
				newFilename := l.option.Path + "/" +
					l.option.FilePrefix +
					fmt.Sprintf("%s", now.Format("2006-01-02-15-04-05")) +
					fileExt
				if err := os.Rename(filename, newFilename); err != nil {
					l.Info("can not rename log file")
				}
			}
		}
	case User:
		l.Logger.SetOutput(l.option.Output)
	default:
		l.Logger.SetOutput(optionDefault.Output)
	}
	if filename != "" {
		filepath := path.Dir(filename)
		if err := os.MkdirAll(filepath, 0644); err == nil || os.IsExist(err) {
			basename := path.Base(filename)
			file, exists := l.writers[basename]
			if !exists {
				if file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err != nil {
					l.Info("can not create log file")
				}
			}
			if file != nil {
				l.Logger.SetOutput(file)
				l.writers[basename] = file
			} else {
				l.Info("can not open log file")
			}
		} else {
			l.Info("can not make log path")
		}

	}

	return l
}
func create(option Option) *log {
	return (&log{
		Logger:  logrus.New(),
		option:  option,
		writers: map[string]*os.File{},
	}).syncLoggerOption()
}
func optionVerify(option ...Option) Option {
	opt := optionDefault
	if len(option) > 0 {
		//设置选项
		opt.Name = option[0].Name
		if toolm.IntSliceContains(option[0].Fmt, fmtContainer) {
			opt.Fmt = option[0].Fmt
		}
		opt.Caller = option[0].Caller
		if toolm.IntSliceContains(option[0].OutMode, outModeContainer) {
			opt.OutMode = option[0].OutMode
		}
		opt.Path = option[0].Path
		opt.FilePrefix = toolm.StringDefault(option[0].FilePrefix, opt.Name, optionDefault.FilePrefix)
		opt.SizeMax = toolm.Int64Default(option[0].SizeMax, optionDefault.SizeMax)
		if _, ok := option[0].Output.(io.Writer); ok {
			opt.Output = option[0].Output
		}
	}
	return opt
}
