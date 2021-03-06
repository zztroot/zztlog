package main

import (
	"zztlog"
)

func main() {
	//通过结构体配置
	//zztlog.InitConfig(zztlog.BaseConfig{LogConfig:zztlog.LogConfig{DebugOutput: true, ErrorOutput: true, CmdOutput: true, ColourOutput: true}})

	//通过文件配置
	err := zztlog.InitConfigFile("zztlog.json")
	if err != nil {
		zztlog.Debug(err)
	}
	zztlog.ErrorF("%s+%s", "我", "你")
	zztlog.Fatal(123456)
	zztlog.FatalF("%s+%s", "我", "你")
	zztlog.Warn(123456)
	zztlog.WarnF("%s+%s", "我", "你")
	//如果文件和结构体都未配置，将默认输出
	for {
		zztlog.Info(123456)
		zztlog.InfoF("%s+%s", "我", "你")
		zztlog.Debug(123456)
		zztlog.DebugF("%s+%s", "我", "你")
		zztlog.Error(123456)
		zztlog.ErrorF("%s+%s", "我", "你")
		zztlog.Fatal(123456)
		zztlog.FatalF("%s+%s", "我", "你")
		zztlog.Warn(123456)
		zztlog.WarnF("%s+%s", "我", "你")
	}
	//zztlog.Info(123456)
}
