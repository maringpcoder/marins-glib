package log

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"log"
)

type LoggerConfig struct {
	FileName            string `json:"filename"`
	Level               int    `json:"level"`    //日志保存的时候的级别，默认是 Trace 级别
	MaxLines            int    `json:"maxlines"` //每个文件保存的最大行数，默认值 1000000
	MaxSize             int    `json:"maxsize"`  //每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
	Daily               bool   `json:"daily"`    //是否按照每天 logrotate，默认是 true
	Maxdays             int    `json:"maxdays"`  //文件最多保存多少天，默认保存 7 天
	Rotate              bool   `json:"rotate"`   //是否开启 logrotate，默认是 true
	Perm                string `json:"perm"`     //日志文件权限
	RotatePerm          string `json:"rotate_perm"`
	EnableFuncCallDepth bool   `json:"-"` //输出文件名行号
	LogFuncCallDepth    int    `json:"-"` //函数调用层级
}

func LogsInit(path string, lev int, rotatePerm string, perm string) error {
	config := LoggerConfig{
		FileName:            path,
		Level:               lev,
		EnableFuncCallDepth: true,
		LogFuncCallDepth:    3,
		RotatePerm:          rotatePerm,
		Perm:                perm,
	}
	marshal, err := json.Marshal(config)
	if err != nil {
		return err
	}
	if err := logs.SetLogger(logs.AdapterFile, string(marshal)); err != nil {
		log.Fatalf("init logs err:%v", err.Error())
	}
	logs.EnableFuncCallDepth(config.EnableFuncCallDepth)
	logs.SetLogFuncCallDepth(config.LogFuncCallDepth)
	logs.Info("beego logs init success")
	return nil
}
