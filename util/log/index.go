package log

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func println(
	color string,
	level string,
	data interface{},
	err error,
	tags []string,
) {
	type json struct {
		level string
		time  string
		tags  string
		data  interface{}
		err   string
	}
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	s := json{
		level: level,
		time:  time.Now().Format(time.RFC3339),
		tags:  strings.Join(tags, ","),
		data:  data,
		err:   errStr,
	}
	fmt.Fprintf(os.Stdout, "\033[%sm%#v\033[0m\n", color, s)
}

func Debug(data interface{}, tags ...string) {
	println(color.blue, level.debug, data, nil, tags)
}

func Info(data interface{}, tags ...string) {
	println(color.green, level.info, data, nil, tags)
}

func Warn(data interface{}, tags ...string) {
	println(color.yellow, level.warn, data, nil, tags)
}

func Error(data interface{}, err error, tags ...string) {
	println(color.red, level.err, data, err, tags)
}
