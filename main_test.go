package main

import (
	"bytes"
	"os"
	"path"
	"runtime"
	"strings"
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

	go serv()
}

func test(t *testing.T, args string) {
	println(args)
	cli := cmd.RootCmd
	buf := new(bytes.Buffer)
	cli.SetOutput(buf)

	cli.SetArgs(strings.Split(args, " "))

	err := cli.Execute()
	if err != nil {
		t.Error(err)
	}

	println(buf.String())
}

func TestCmd(t *testing.T) {
	executable := os.Getenv("PROTOAPI_EXE")
	if executable == "" {
		t.Error("PROTOAPI_EXE is not set")
		return
	}

	test(t, "gen --lang=go test/result/go test/proto/test.proto")
	// test(t, "gen --lang=go test/result/go test/proto/calc.proto")
	// test(t, "gen --lang=go test/result/go test/proto/todolist.proto")
	// test(t, "gen --lang=go test/result/go test/proto/nested.proto")

	test(t, "gen --lang=ts test/result/ts test/proto/test.proto")
	test(t, "gen --lang=spring test/result/ test/proto/test.proto")

	test(t, "gen --lang=phpclient test/result/ test/proto/test.proto")
	test(t, "gen --lang=yii2 test/result/ test/proto/todolist.proto")
	test(t, "gen --lang=markdown test/result/ test/proto/login.proto")
}
