package zztlog

import (
	"bufio"
	"github.com/gookit/color"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type logHandler struct {
	m           sync.Mutex
	out         io.Writer
	buf         []byte
	isInit      bool
	currentLine int64
	newFileName string
}

type BaseConfig struct {
	LogConfig LogConfig `json:"log_config"`
}

type LogConfig struct {
	TimeFormat        string `json:"time_format"`
	Prefix            string `json:"prefix"`
	FileOutput        bool   `json:"file_output"`
	CmdOutput         bool   `json:"cmd_output"`
	FileAllPathOutput bool   `json:"file_all_path_output"`
	FuncNameOutput    bool   `json:"func_name_output"`
	MaxFileLine       int64  `json:"max_file_line"`
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
	config.LogConfig.TimeFormat = "2006-01-02 15:04:05"
}

func (l *logHandler) handlerOutput(name string) bool {
	switch name {
	case err:
		if !config.LogConfig.ErrorOutput {
			return false
		}
	case debug:
		if !config.LogConfig.DebugOutput {
			return false
		}
	case fatal:
		if !config.LogConfig.FatalOutput {
			return false
		}
	case info:
		if !config.LogConfig.InfoOutput {
			return false
		}
	case warn:
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
		case debug:
			temp = []byte(color.Cyan.Sprintf(string(l.buf)))
		case err:
			temp = []byte(color.Magenta.Sprintf(string(l.buf)))
		case fatal:
			temp = []byte(color.Red.Sprintf(string(l.buf)))
		case info:
			temp = []byte(color.Green.Sprintf(string(l.buf)))
		case warn:
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
	fileName := "zztlog.log"
	if config.LogConfig.SaveFileName != "" {
		fileName = config.LogConfig.SaveFileName
	}
	if l.newFileName != "" {
		fileName = l.newFileName
	}
	fileInfo, e := os.Stat(fileName)
	var setSize int64
	var fileSize int64
	if !os.IsNotExist(e) {
		l.currentLine = fileLineNumber(fileName)
		fileSize = fileInfo.Size()
	} else {
		fileSize = -1
		l.currentLine++
	}
	// 文件大小初始化
	if config.LogConfig.MaxSizeM != 0 {
		setSize = config.LogConfig.MaxSizeM * 1024 * 1024
	} else {
		setSize = 10 * 1024 * 1024
	}
	// 切割路径
	paths := filepath.Dir(fileName)
	//if strings.Contains(fileName, "/") {
	//	if string(fileName[0]) != "." {
	//		fileName = "." + fileName
	//	}
	//	t := strings.Split(fileName, "/")
	//	for _, v := range t[:len(t)-1] {
	//		if paths == "" {
	//			paths = v
	//		} else {
	//			paths = paths + string(os.PathSeparator) + v
	//		}
	//	}
	//}

	if config.LogConfig.MaxFileLine != 0 {
		l.fileCutting(l.currentLine, config.LogConfig.MaxFileLine, fileName, paths)
		if l.currentLine >= config.LogConfig.MaxFileLine {
			l.currentLine = 0
		}
		return
	}
	l.fileCutting(fileSize, setSize, fileName, paths)
}

// 文件切割
func (l *logHandler) fileCutting(a, b int64, fileName, paths string) {
	m := sync.Mutex{}
	m.Lock()
	defer m.Unlock()
	// 如果目录不存在，创建
	if paths != "" {
		if _, err := os.Stat(paths); os.IsNotExist(err) {
			if err = os.MkdirAll(paths, os.ModePerm); err != nil {
				log.Panic(err)
				return
			}
		}
	}
	if a >= b {
		// 创建新的文件
		fileName = time.Now().Format(config.LogConfig.TimeFormat) + ".log"
		if paths != "" {
			fileName = paths + string(os.PathSeparator) + fileName
		}
		l.newFileName = fileName
		createFile(fileName)
	} else {
		createFile(fileName)
	}
}

func fileLineNumber(path string) int64 {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
		return 0
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	var lineCount int64
	for fileScanner.Scan() {
		lineCount++
	}
	return lineCount
}

func createFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Panic(err)
		return
	}
	defer file.Close()
	_, err = file.Write(l.buf)
	if err != nil {
		log.Panic(err)
		return
	}
}

func (l *logHandler) debug(s string) {
	l.m.Lock()
	defer l.m.Unlock()
	l.output(debug, s)
}

func (l *logHandler) info(s string) {
	l.m.Lock()
	defer l.m.Unlock()
	l.output(info, s)
}

func (l *logHandler) warn(s string) {
	l.m.Lock()
	defer l.m.Unlock()
	l.output(warn, s)
}

func (l *logHandler) fatal(s string) {
	l.m.Lock()
	defer l.m.Unlock()
	l.output(fatal, s)
}

func (l *logHandler) error(s string) {
	l.m.Lock()
	defer l.m.Unlock()
	l.output(err, s)
}
