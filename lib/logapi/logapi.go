package logapi

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

const (
	color_red = uint8(iota + 91)
	color_green
	color_yellow
	color_blue
	color_magenta //洋红

	info  = "[INFO]"
	trace = "[TRACE]"
	erro  = "[ERRO]"
	warn  = "[WARN]"
	debug = "[DEBUG]"
)

// see complete color rules in document in https://en.wikipedia.org/wiki/ANSI_escape_code#cite_note-ecma48-13
func TRACE(args ...interface{}) {
	f := inspect(args...)
	content := formatContent(trace, f)
	log.Println(output(color_yellow, content))
}

func INFO(args ...interface{}) {
	f := inspect(args...)
	content := formatContent(info, f)
	log.Println(output(color_blue, content))
}

func DEBUG(args ...interface{}) {
	f := inspect(args...)
	content := formatContent(debug, f)
	log.Println(output(color_green, content))
}

func WARN(args ...interface{}) {
	f := inspect(args...)
	content := formatContent(warn, f)
	log.Println(output(color_magenta, content))
}

func ERROR(args ...interface{}) {
	f := fmt.Sprintf("%s:%s", fileInfo(2), inspect(args...))
	content := formatContent(erro, f)
	log.Println(output(color_red, content))
}

func Inspect(dump ...interface{}) string {
	return inspect(dump...)
}

func inspect(dump ...interface{}) string {
	return spew.Sdump(dump...)
}

func output(d uint8, s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", d, s)
}

func formatContent(prefix string, content string) string {
	return time.Now().Format("2006/01/02 15:04:05") + " " + prefix + " " + content + " "
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d\n", file, line)
}
