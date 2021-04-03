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

type logConfig struct {
	TimeFormat        string `json:"time_format"`
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
	if config.FileAllPathOutput {
		f = file
	} else {

		fileList := strings.Split(file, "/")
		f = fileList[len(fileList)-1]
	}
	l.buf = append(l.buf, "["...)
	l.buf = append(l.buf, name...)
	l.buf = append(l.buf, "] "...)
	nowTime := time.Now().Format(config.TimeFormat + " ")
	l.buf = append(l.buf, nowTime...)
	l.buf = append(l.buf, f...)
	l.buf = append(l.buf, ":"...)
	l.buf = append(l.buf, strconv.Itoa(line)...)
	if config.FuncNameOutput {
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
	if config.CmdOutput && config.FileOutput {
		l.outputCmd(name)
		l.outputFile()
	} else if config.FileOutput {
		l.outputFile()
	} else {
		l.outputCmd(name)
	}
	l.buf = nil
	return
}

func (l *logHandler) initConfig() {
	config.WarnOutput = true
	config.InfoOutput = true
	config.FatalOutput = true
	config.ErrorOutput = true
	config.DebugOutput = true
	config.CmdOutput = true
	config.MaxSizeM = 100
	config.TimeFormat = "2006-01-02 15:04:05"
}

func (l *logHandler) handlerOutput(name string) bool {
	switch name {
	case "Debug":
		if !config.DebugOutput {
			return false
		}
	case "Error":
		if !config.ErrorOutput {
			return false
		}
	case "Fatal":
		if !config.FatalOutput {
			return false
		}
	case "Info-":
		if !config.InfoOutput {
			return false
		}
	case "Warn-":
		if !config.WarnOutput {
			return false
		}
	}
	return true
}

func (l *logHandler) outputCmd(name string) {
	var temp []byte
	if config.ColourOutput {
		switch name {
		case "Debug":
			temp = []byte(color.Cyan.Sprintf(string(l.buf)))
		case "Error":
			temp = []byte(color.Magenta.Sprintf(string(l.buf)))
		case "Fatal":
			temp = []byte(color.Red.Sprintf(string(l.buf)))
		case "Info-":
			temp = []byte(color.Green.Sprintf(string(l.buf)))
		case "Warn-":
			temp = []byte(color.Yellow.Sprintf(string(l.buf)))
		}
	}else {
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
	if config.SaveFileName != "" {
		fileName = config.SaveFileName
	}
	fileInfo, e := os.Stat(fileName)
	var setSize int64
	var fileSize int64
	if !os.IsNotExist(e) {
		setSize = config.MaxSizeM * 1024 * 1024
		fileSize = fileInfo.Size()
	} else {
		fileSize = -1
	}
	if fileSize >= setSize {
		err := os.Rename(fileName, nowName+".log")
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
	l.output("Debug", s)
}

func (l *logHandler) info(s string) {
	l.m.Lock()
	defer l.m.Unlock()
	l.output("Info-", s)
}

func (l *logHandler) warn(s string) {
	l.m.Lock()
	defer l.m.Unlock()
	l.output("Warn-", s)
}

func (l *logHandler) fatal(s string) {
	l.m.Lock()
	defer l.m.Unlock()
	l.output("Fatal", s)
}

func (l *logHandler) error(s string) {
	l.m.Lock()
	defer l.m.Unlock()
	l.output("Error", s)
}
