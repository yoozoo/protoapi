package main

import (
	"bytes"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/yoozoo/protoapi/cmd"
)

func init() {
	// make sure running in source directory
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Dir(filename)
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
func TestCmd(t *testing.T) {
	cli := cmd.RootCmd
	buf := new(bytes.Buffer)
	cli.SetOutput(buf)
	cli.SetArgs([]string{
		"gen", "--lang=go", "test/result/go", "test/proto/test.proto",
	})
	// cli.SetArgs([]string{
	// 	"help",
	// })

	err := cli.Execute()
	if err != nil {
		t.Error(err)
	}

	println(buf.String())
}
