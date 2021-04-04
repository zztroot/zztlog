package zztlog

import (
	"github.com/gookit/color"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type logHandler struct {
	m      sync.Mutex
	out    io.Writer
	buf    []byte
	isInit bool
}

type base struct {
	LogConfig logConfig `json:"log_config"`
}

type logConfig struct {
	TimeFormat        string `json:"time_format"`
	Prefix            string `json:"prefix"`
	FileOutput        bool   `json:"file_output"`
	CmdOutput         bool   `json:"cmd_output"`
	FileAllPathOutput bool   `json:"file_all_path_output"`
	FuncNameOutput    bool   `json:"func_name_output"`
	MaxSizeM          int64  `json:"max_size_m"`
	SaveFileName      string `json:"save_file_name"`
	ErrorOutput       bool   `json:"error_output"`
	FatalOutput       bool   `json:"fatal_output"`
	WarnOutput        bool   `json:"warn_output"`
	InfoOutput        bool   `json:"info_output"`
	DebugOutput       bool   `json:"debug_output"`
	ColourOutput      bool   `json:"colour_output"`
}

func (l *logHandler) output(name, s string) {
	mutex := sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	if !l.isInit {
		l.initConfig()
	}
	if !l.handlerOutput(name) {
		return
	}
	l.out = os.Stdout
	l.m.Unlock()
	ptr, file, line, ok := runtime.Caller(3)
	l.m.Lock()
	if !ok {
		file = "/???"
		line = 0
	}
	//判断文件路径显示
	var f string
	if config.LogConfig.FileAllPathOutput {
		f = file
	} else {

		fileList := strings.Split(file, "/")
		f = fileList[len(fileList)-1]
	}
	if config.LogConfig.Prefix != "" {
		l.buf = append(l.buf, config.LogConfig.Prefix+" "...)
	}
	l.buf = append(l.buf, "["...)
	l.buf = append(l.buf, name...)
	l.buf = append(l.buf, "] "...)
	nowTime := time.Now().Format(config.LogConfig.TimeFormat + " ")
	l.buf = append(l.buf, nowTime...)
	l.buf = append(l.buf, f...)
	l.buf = append(l.buf, ":"...)
	l.buf = append(l.buf, strconv.Itoa(line)...)
	if config.LogConfig.FuncNameOutput {
		funcName := runtime.FuncForPC(ptr).Name()
		funcName = path.Base(funcName)
		l.buf = append(l.buf, " ["...)
		l.buf = append(l.buf, funcName...)
		l.buf = append(l.buf, "]"...)
	}
	l.buf = append(l.buf, " "+s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}

	//判断是否写入文件还是cmd
	if config.LogConfig.CmdOutput && config.LogConfig.FileOutput {
		l.outputCmd(name)
		l.outputFile()
	} else if config.LogConfig.FileOutput {
		l.outputFile()
	} else {
		l.outputCmd(name)
	}
	l.buf = nil
	return
}

func (l *logHandler) initConfig() {
	config.LogConfig.WarnOutput = true
	config.LogConfig.InfoOutput = true
	config.LogConfig.FatalOutput = true
	config.LogConfig.ErrorOutput = true
	config.LogConfig.DebugOutput = true
	config.LogConfig.CmdOutput = true
	config.LogConfig.MaxSizeM = 100
	config.LogConfig.TimeFormat = "2006-01-02 15:04:05"
}

func (l *logHandler) handlerOutput(name string) bool {
	switch name {
	case ERROR:
		if !config.LogConfig.ErrorOutput {
			return false
		}
	case DEBUG:
		if !config.LogConfig.DebugOutput {
			return false
		}
	case FATAL:
		if !config.LogConfig.FatalOutput {
			return false
		}
	case INFO:
		if !config.LogConfig.InfoOutput {
			return false
		}
	case WARN:
		if !config.LogConfig.WarnOutput {
			return false
		}
	}
	return true
}

func (l *logHandler) outputCmd(name string) {
	var temp []byte
	if config.LogConfig.ColourOutput {
		switch name {
		case DEBUG:
			temp = []byte(color.Cyan.Sprintf(string(l.buf)))
		case ERROR:
			temp = []byte(color.Magenta.Sprintf(string(l.buf)))
		case FATAL:
			temp = []byte(color.Red.Sprintf(string(l.buf)))
		case INFO:
			temp = []byte(color.Green.Sprintf(string(l.buf)))
		case WARN:
			temp = []byte(color.Yellow.Sprintf(string(l.buf)))
		}
	} else {
		temp = l.buf
	}
	_, err := l.out.Write(temp)
	if err != nil {
		log.Println(err)
	}
}

func (l *logHandler) outputFile() {
	nowName := time.Now().Format("20060102150405")
	fileName := "zztlog.log"
	if config.LogConfig.SaveFileName != "" {
		fileName = config.LogConfig.SaveFileName
	}
	fileInfo, e := os.Stat(fileName)
	var setSize int64
	var fileSize int64
	if !os.IsNotExist(e) {
		setSize = config.LogConfig.MaxSizeM * 1024 * 1024
		fileSize = fileInfo.Size()
	} else {
		fileSize = -1
	}
	var paths string
	if strings.Contains(fileName, "/") {
		t := strings.Split(fileName, "/")
		for _, v := range t[:len(t)-1] {
			if paths == "" {
				paths = v
			} else {
				paths = paths + "/" + v
			}
		}
	}
	if fileSize >= setSize {
		err := os.Rename(fileName, paths+"/"+nowName+".log")
		if err != nil {
			log.Println("Failed to change file name")
			return
		}
		createFile(fileName)
	} else {
		createFile(fileName)
	}
}

func createFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("File opening failed")
		return
	}
	defer file.Close()
	_, err = file.Write(l.buf)
	if err != nil {
		log.Println("File write failed")
		return
	}
}

func (l *logHandler) debug(s string) {
	l.m.Lock()
	defer l.m.Unlock()
	l.output(DEBUG, s)
}

func (l *logHandler) info(s string) {
	l.m.Lock()
	defer l.m.Unlock()
	l.output(INFO, s)
}

func (l *logHandler) warn(s string) {
	l.m.Lock()
	defer l.m.Unlock()
	l.output(WARN, s)
}

func (l *logHandler) fatal(s string) {
	l.m.Lock()
	defer l.m.Unlock()
	l.output(FATAL, s)
}

func (l *logHandler) error(s string) {
	l.m.Lock()
	defer l.m.Unlock()
	l.output(ERROR, s)
}
