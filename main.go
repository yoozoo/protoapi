package main

//go:generate esc -o generator/data/tpl/tpl.go -modtime 0 -pkg=tpl generator/template
//go:generate esc -o util/protoapi_include.go -modtime 0 -pkg=util proto

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/debug"

	"github.com/golang/protobuf/proto"

	"github.com/yoozoo/protoapi/cmd"
	"github.com/yoozoo/protoapi/generator"
	"github.com/yoozoo/protoapi/util"
)

func main() {
	defer func() {
		util.CleanIncludePath()
		if r := recover(); r != nil {
			log.Printf("%s: %s", r, debug.Stack())
			os.Exit(1)
		}
	}()

	stat, err := os.Stdin.Stat()
	args := os.Args

	// when no any parameter and not reading from char device, treat it as being called by protoc
	if len(args) == 1 && err == nil && (stat.Mode()&os.ModeCharDevice) == 0 {
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			util.Die(fmt.Errorf("reading stdin error: %s", err.Error()))
		}

		response := generator.Generate(input)

		output, err := proto.Marshal(response)
		if err != nil {
			util.Die(fmt.Errorf("create response failed: %s", err.Error()))
		}
		_, err = os.Stdout.Write(output)
		if err != nil {
			util.Die(err)
		}
	} else {
		cmd.Execute()
	}
}
