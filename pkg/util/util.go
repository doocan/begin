package util

import (
	"path"
	"runtime"
)

func GetCurrentPath() string {
	_, filename, _, _ := runtime.Caller(2)
	return path.Dir(filename)
}
