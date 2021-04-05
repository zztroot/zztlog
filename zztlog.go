package zztlog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

var (
	l      logHandler
	config BaseConfig
	m      sync.Mutex
)

type initLog struct{}

func InitConfigFile(path string) error {
	r, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(r, &config)
	if err != nil {
		return err
	}
	l.isInit = true
	return nil
}
func InitConfig(b BaseConfig) {
	handleBaseConfig(&b)
	config = b
	l.isInit = true
}

func handleBaseConfig(b *BaseConfig) {
	if !b.LogConfig.CmdOutput {
		b.LogConfig.CmdOutput = true
	}
	if b.LogConfig.TimeFormat == "" {
		b.LogConfig.TimeFormat = "2006-01-02 15:04:05"
	}
	if b.LogConfig.MaxSizeM == 0 {
		b.LogConfig.MaxSizeM = 100
	}
}

func Default() *initLog { return &initLog{} }

func Debug(s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.debug(fmt.Sprint(s...))
}

func Info(s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.info(fmt.Sprint(s...))
}

func Fatal(s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.fatal(fmt.Sprint(s...))
}

func Warn(s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.warn(fmt.Sprint(s...))
}

func Error(s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.error(fmt.Sprint(s...))
}

func DebugF(format string, s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.debug(fmt.Sprintf(format, s...))
}

func InfoF(format string, s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.info(fmt.Sprintf(format, s...))
}

func FatalF(format string, s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.fatal(fmt.Sprintf(format, s...))
}

func WarnF(format string, s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.warn(fmt.Sprintf(format, s...))
}

func ErrorF(format string, s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.error(fmt.Sprintf(format, s...))
}

func (i *initLog) Debug(s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.debug(fmt.Sprint(s...))
}

func (i *initLog) Info(s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.info(fmt.Sprint(s...))
}

func (i *initLog) Fatal(s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.fatal(fmt.Sprint(s...))
}

func (i *initLog) Warn(s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.warn(fmt.Sprint(s...))
}

func (i *initLog) Error(s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.error(fmt.Sprint(s...))
}

func (i *initLog) DebugF(format string, s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.debug(fmt.Sprintf(format, s...))
}

func (i *initLog) InfoF(format string, s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.info(fmt.Sprintf(format, s...))
}

func (i *initLog) FatalF(format string, s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.fatal(fmt.Sprintf(format, s...))
}

func (i *initLog) WarnF(format string, s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.warn(fmt.Sprintf(format, s...))
}

func (i *initLog) ErrorF(format string, s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	l.error(fmt.Sprintf(format, s...))
}
