# zztlog
golang日志库，支持输出到终端、文件，可以设置文件最大大小，达到大小后将写入新的日志文件中，终端颜色显示，显示文件名称或全路径，显示行数，显示函数名称等等

##配置文件说明(如果没有配置文件，将输出默认格式)
```json
{
  "save_file_name": "zztlog.log",
  "time_format": "2006/01/02 15:04:05",
  "max_size_m":100,
  "file_output": true,
  "cmd_output": true,
  "file_all_path_output": false,
  "colour_output": true,
  "func_name_output": true,
  "error_output": true,
  "fatal_output": true,
  "warn_output": true,
  "info_output": true,
  "debug_output": true
}
```
