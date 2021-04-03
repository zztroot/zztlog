package zztlog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

var (
	l      logHandler
	config logConfig
	m      sync.Mutex
)

func InitConfig(name string) error {
	r, err := ioutil.ReadFile(name)
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
