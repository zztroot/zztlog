# zztlog
（zztlog）golang日志库，支持输出到终端、文件，可以设置文件大小切割，终端颜色显示，显示文件名称或全路径，显示行数，显示函数名称等等

## 安装
```
go get github.com/zztroot/zztlog
```

## 配置文件说明(如果没有配置文件，将输出默认格式)
```json
{
  "log_config": {
    "save_file_name": "log/zztlog.log",
    "time_format": "2006/01/02 15:04:05",
    "max_size_m":1,
    "prefix": "[测试]",
    "file_output": true,
    "cmd_output": true,
    "file_all_path_output": true,
    "colour_output": true,
    "func_name_output": true,
    "error_output": true,
    "fatal_output": true,
    "warn_output": true,
    "info_output": true,
    "debug_output": true
  }
}
```
#### 写入文件相关:
file_output: 是否输出到文件(默认为false)  
save_file_name: 写入文件的名称(默认zztlog.log)  
max_size_m: 写入文件的最大大小，单位M，当达到最大时将会写入新的日志文件中。(默认大小100M)    

#### 写入终端相关:
cmd_output: 是否输出到终端(默认为true)  
colour_output: 输出是否带颜色(默认为false)  

#### 公共
prefix: 输出前缀(默认为空)  
time_format: 输出时间格式(默认格式:2006-01-02 15:04:05)  
file_all_path_output: 是否显示文件全路径(默认只显示文件名称faslse)  
func_name_output: 是否显示函数名称(默认不显示false)  
error_output: 是否输出error信息(默认为true)  
fatal_output: 是否输出fatal信息(默认为true)  
warn_output: 是否输出warn信息(默认为true)  
info_output: 是否输出info信息(默认为true)  
debug_output: 是否输出debug信息(默认为true)  

## 例子
```go
package main

import "zztlog"

func main() {
	err := zztlog.InitConfig("zztlog.json") //可配置
	if err != nil {
		zztlog.Debug(err)
	}
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

```
#### 输出
```
[Info-] 2021/04/04 01:19:58 main.go:11 [main.main] 123456
[Info-] 2021/04/04 01:19:58 main.go:12 [main.main] 我+你
[Debug] 2021/04/04 01:19:58 main.go:13 [main.main] 123456
[Debug] 2021/04/04 01:19:58 main.go:14 [main.main] 我+你
[Error] 2021/04/04 01:19:58 main.go:15 [main.main] 123456
[Error] 2021/04/04 01:19:58 main.go:16 [main.main] 我+你
[Fatal] 2021/04/04 01:19:58 main.go:17 [main.main] 123456
[Fatal] 2021/04/04 01:19:58 main.go:18 [main.main] 我+你
[Warn-] 2021/04/04 01:19:58 main.go:19 [main.main] 123456
[Warn-] 2021/04/04 01:19:58 main.go:20 [main.main] 我+你
[Info-] 2021/04/04 01:19:58 main.go:11 [main.main] 123456
[Info-] 2021/04/04 01:19:58 main.go:12 [main.main] 我+你
[Debug] 2021/04/04 01:19:58 main.go:13 [main.main] 123456
[Debug] 2021/04/04 01:19:58 main.go:14 [main.main] 我+你
[Error] 2021/04/04 01:19:58 main.go:15 [main.main] 123456
[Error] 2021/04/04 01:19:58 main.go:16 [main.main] 我+你
[Fatal] 2021/04/04 01:19:58 main.go:17 [main.main] 123456
[Fatal] 2021/04/04 01:19:58 main.go:18 [main.main] 我+你
```


