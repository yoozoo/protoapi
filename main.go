package main

import (
	"io/ioutil"
	"log"
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
			log.Fatalf("reading file %s error: %s\n", argsWithoutProg[0], err.Error())
			// os.Exit(1)
		}
	} else {
		var err error
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("reading stdin error: %s\n", err.Error())
			// os.Exit(1)
		}
	}

	request := new(plugin.CodeGeneratorRequest)

	proto.Unmarshal(input, request)
	if len(request.FileToGenerate) != 1 {
		log.Fatalf("input files are： %v\nwe only support one proto file\n", request.FileToGenerate)
		// os.Exit(2)
	}

	response, err := generator.Generate(request)
	if err != nil {
		log.Fatalf("generate response failed: %s\n", err.Error())
		// os.Exit(3)
	}

	output, err := proto.Marshal(response)
	if err != nil {
		log.Fatalf("create response failed: %s\n", err.Error())
		// os.Exit(4)
	}
	os.Stdout.Write(output)
}
