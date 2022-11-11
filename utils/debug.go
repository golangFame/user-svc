package utils

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Debug prints a debug information to the log with file and line.
func Debug(format string, a ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	info := fmt.Sprintf(format, a...)

	log.Printf("[cgl] debug %s:%d %v", file, line, info)
}

// Function getting the function name of the previous caller
func Function(skips ...int) string {
	skip := 1
	if len(skips) > 0 {
		skip = skips[0]
	}

	counter, _, _, success := runtime.Caller(skip)

	if !success {
		log.Fatal("functionName: runtime.Caller: failed")
	}

	return runtime.FuncForPC(counter).Name()
}

func FunctionName(skips ...int) string {
	name := Function(skips...)
	names := strings.Split(name, ".")
	return names[len(names)-1]
}

func GetFunctionName(temp interface{}) string {
	strs := strings.Split(runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name(), ".")
	return strs[len(strs)-1]
}

func GetFunctionNameOnly(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
