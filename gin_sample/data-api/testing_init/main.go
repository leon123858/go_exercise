package testing_init

import (
	"os"
	"path"
	"runtime"
)

var MostInit = mostInit()

func mostInit() int {
	_, filename, _, _ := runtime.Caller(0)
	println(filename)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
	return 0
}
