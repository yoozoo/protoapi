package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"

	"protoapi/generator"
)

func main() {

	argsWithoutProg := os.Args[1:]
	var input []byte
	if len(argsWithoutProg) > 0 {
		var err error
		input, err = ioutil.ReadFile(argsWithoutProg[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading file %s error: %s\n", argsWithoutProg[0], err.Error())
			os.Exit(1)
		}
	} else {
		var err error
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading stdin error: %s\n", err.Error())
			os.Exit(1)
		}
	}

	request := new(plugin.CodeGeneratorRequest)

	files := request.GetProtoFile()

	for _, v := range files {
		fmt.Println("file: ", v)
	}

	proto.Unmarshal(input, request)
	if len(request.FileToGenerate) != 1 {
		fmt.Fprintf(os.Stderr, "input files are： %v\nwe only support one proto file\n", request.FileToGenerate)
		os.Exit(2)
	}

	response, err := generator.Generate(request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "generate response failed: %s\n", err.Error())
		os.Exit(3)
	}

	output, err := proto.Marshal(response)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create response failed: %s\n", err.Error())
		os.Exit(4)
	}
	os.Stdout.Write(output)
}
