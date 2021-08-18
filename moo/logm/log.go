package logm

import (
	"fmt"
	"github.com/han-joker/moo-layout/moo/dftm"
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
	"strings"
)

//类型
//日志
type log struct {
	//选项
	opt Opt
}
type Opt struct {
	Fmt string
	Caller bool
	Out *os.File
}

//变量
//日志单例
var instance *log

//默认选项
var dftOpt = Opt{
	Fmt: "text",
	Caller: false,
	Out: os.Stderr,
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
			instance.opt.Fmt = dftm.String(opt.Fmt, dftOpt.Fmt)
			instance.opt.Caller = opt.Caller
			if opt.Out != nil {
				fmt.Println( "type:", reflect.TypeOf(opt.Out), reflect.TypeOf([]int{}).String() )
				if reflect.TypeOf(opt.Out).Name() == "" {

				}
				instance.opt.Out = opt.Out
			}
		}
		// 基于选项配置 logrus
		switch strings.ToLower(instance.opt.Fmt) {
		case "json":
			logrus.SetFormatter(&logrus.JSONFormatter{})
		case "text":
			fallthrough
		default:
			logrus.SetFormatter(&logrus.TextFormatter{})
		}
		logrus.SetReportCaller(instance.opt.Caller)

	}
	return instance
}

//可导出方法
func (log) Info(args ...interface{}) {
	logrus.Info(args...)
}