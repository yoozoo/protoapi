package main

//go:generate esc -o generator/data/tpl/tpl.go -pkg=tpl generator/template

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"

	"protoapi/generator"
)

func main() {
	args := os.Args
	var input []byte
	if len(args[1:]) > 0 {
		var err error
		input, err = ioutil.ReadFile(args[2])
		if err != nil {
			log.Fatalf("reading file %s error: %s\n", args[2], err.Error())
		}
	} else {
		var err error
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("reading stdin error: %s\n", err.Error())
		}
	}

	// if run with protoapi then run protoc with plugin
	if args[0] == "protoapi" {
		executable, _ := os.Executable()

		args[0] = "--plugin=protoc-gen-custom=" + executable
		args[1] = strings.Replace(args[1], "--", "--custom_out=", 1)

		// Run protoc command
		cmd := exec.Command("protoc", args...)
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}

		return
	}

	request := new(plugin.CodeGeneratorRequest)

	proto.Unmarshal(input, request)
	if len(request.FileToGenerate) != 1 {
		log.Fatalf("input files are： %v\nwe only support one proto file\n", request.FileToGenerate)
	}

	response, err := generator.Generate(request)
	if err != nil {
		log.Fatalf("generate response failed: %s\n", err.Error())
	}

	output, err := proto.Marshal(response)
	if err != nil {
		log.Fatalf("create response failed: %s\n", err.Error())
	}
	os.Stdout.Write(output)
}
