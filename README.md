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
    "max_file_line": 0,
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
max_size_m: 写入文件的最大大小，单位M，当达到最大时将会写入新的日志文件中。(默认大小10M)    
max_file_line: 文件行数切割，(为0表示不通过文件行数切割，默认为0)，此项如果打开，max_size_m设置将自动忽略。

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

import "github.com/zztroot/zztlog"

func main() {
	//通过文件配置
	if err := zztlog.InitConfigFile("zztlog.json"); err != nil {
		zztlog.Error(err)
		return
	}
	
	//通过结构体配置
	zztlog.InitConfig(zztlog.BaseConfig{LogConfig:zztlog.LogConfig{DebugOutput: true, ErrorOutput: true, CmdOutput: true, ColourOutput: true}})

	//可以获取结构体
	loggler := zztlog.Default()
	loggler.Info(456789)
	loggler.Error("sdfdsfsdffffffffffff")

	//如果什么配置都不做，也可以直接使用，但格式默认
	zztlog.Info(123)
	zztlog.Error(456545646)
	zztlog.ErrorF(`%s`, "你是我的眼睛")
}


```
#### 输出
```
[Info-] 2021/04/04 14:34:12 F:/golang/test/main.go:14 [main.main] 456789
[Error] 2021/04/04 14:34:13 F:/golang/test/main.go:15 [main.main] sdfdsfsdffffffffffff
[Debug] 2021/04/04 14:34:13 F:/golang/test/main.go:16 [main.main] 我是debug
[Info-] 2021/04/04 14:34:13 F:/golang/test/main.go:19 [main.main] 123
[Error] 2021/04/04 14:34:13 F:/golang/test/main.go:20 [main.main] 456545646
[Error] 2021/04/04 14:34:13 F:/golang/test/main.go:21 [main.main] 你是我的眼睛

```

![在这里插入图片描述](https://img-blog.csdnimg.cn/20210404143512599.PNG?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NDU0NjM0MA==,size_16,color_FFFFFF,t_70#pic_center)



