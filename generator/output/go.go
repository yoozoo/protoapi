package output

import (
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"version.uuzu.com/Merlion/protoapi/generator/data"
)

// Re-use everything in echoGen, only use different template
type goGen struct {
	echoGen
}

func (g *goGen) Init(request *plugin.CodeGeneratorRequest) {
	g.echoGen.Init(request)

	g.structTpl = g.getTpl("/generator/template/go/struct.gogo")
	g.serviceTpl = g.getTpl("/generator/template/go/service.gogo")
	g.enumTpl = g.getTpl("/generator/template/go/enum.gogo")
}

func init() {
	data.OutputMap["go"] = &goGen{}
}
