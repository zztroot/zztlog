package main

import "zztlog"

func main() {
	err := zztlog.InitConfig("zztlog.json")
	if err != nil {
		zztlog.Debug(err)
	}
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
}
