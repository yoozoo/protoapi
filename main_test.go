package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
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
	cli := cmd.RootCmd
	buf := new(bytes.Buffer)
	cli.SetOutput(buf)

	cli.SetArgs(strings.Split(args, " "))

	err := cli.Execute()
	if err != nil {
		println(args)
		t.Error(err)
		println(buf.String())
	}
}

func TestCmd(t *testing.T) {
	executable := os.Getenv("PROTOAPI_EXE")
	if executable == "" {
		t.Error("PROTOAPI_EXE is not set")
		return
	}

	testCmds := `
	gen --lang=go test/result/go test/proto/test.proto
	gen --lang=go test/result/go test/proto/echo.proto
	gen --lang=go test/result/go test/proto/calc.proto
	gen --lang=go test/result/go test/proto/todolist.proto
	gen --lang=go test/result/go test/proto/nested.proto

	gen --lang=go test/result/package/go test/proto/package/common.proto
	gen --lang=go test/result/package/go test/proto/package/gopackage_addReqFull.proto
	gen --lang=go test/result/package/go test/proto/package/gopackage_addReq.proto
	gen --lang=go test/result/package/go test/proto/package/gopackage_calcFull.proto
	gen --lang=go test/result/package/go test/proto/package/gopackage_calc.proto
	gen --lang=go test/result/package/go test/proto/package/gopackage_calc_warn.proto
	gen --lang=go test/result/package/go test/proto/package/mixpackage_addReq.proto
	gen --lang=go test/result/package/go test/proto/package/mixpackage_calc.proto
	gen --lang=go test/result/package/go test/proto/package/nopackage_calc.proto
	gen --lang=go test/result/package/go test/proto/package/nopackage_calc_warn.proto
	gen --lang=go test/result/package/go test/proto/package/package_addReq.proto
	gen --lang=go test/result/package/go test/proto/package/package_calc_commonerror.proto
	gen --lang=go test/result/package/go test/proto/package/package_calc.proto
	gen --lang=go test/result/package/go test/proto/package/package_calc._without_commonerror.proto

	gen --lang=ts test/result/ts test/proto/test.proto
	gen --lang=ts-fetch test/result/ts/fetch test/proto/test.proto
	gen --lang=ts-axios test/result/ts/axios test/proto/test.proto
	gen --lang=ts-wechat test/result/ts/wechat test/proto/test.proto

	gen --lang=spring test/result/ test/proto/test.proto

	gen --lang=phpclient test/result/ test/proto/test.proto

	gen --lang=yii2 test/result/ test/proto/todolist.proto

	gen --lang=markdown test/result/ test/proto/login.proto

	`
	testCmds = strings.TrimSpace(testCmds)

	for _, cmd := range strings.Split(testCmds, "\n") {
		cmd = strings.TrimSpace(cmd)
		if cmd != "" {
			test(t, cmd)
		}
	}

	files, err := filePathWalkDir("test/expected")
	if err != nil {
		t.Fatal(err)
	}

	for _, expectedFileName := range files {
		resultFileName := strings.Replace(expectedFileName, "expected", "result", 1)

		expectedFile, err := ioutil.ReadFile(expectedFileName)
		if err != nil {
			t.Error("expectedFile read err: ", expectedFileName, err)
		}

		resultFile, err := ioutil.ReadFile(resultFileName)
		if err != nil {
			t.Error("resultFile read err: ", resultFileName, err)
		}

		if string(expectedFile) != string(resultFile) {
			dmp := diffmatchpatch.New()
			diffs := dmp.DiffMain(string(expectedFile), string(resultFile), false)
			t.Error("file different: ", expectedFileName, resultFileName)
			fmt.Println(dmp.DiffPrettyText(diffs))
		}
	}
}

func filePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
