package output

import (
	"bytes"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"version.uuzu.com/Merlion/protoapi/generator/data"
	"version.uuzu.com/Merlion/protoapi/util"
)

// Re-use everything in echoGen, only use different template
type goGen struct {
	echoGen
}

type goService struct {
	*echoService
}

func (s *goService) Foo(method *echoMethod) string {
	return method.ErrorType()
}

func (s *goService) Foo1(method *echoMethod) string {
	return method.ErrorType()
}

func (g *goGen) genGoServie(service *data.ServiceData) string {
	buf := bytes.NewBufferString("")

	obj := newEchoService(service, g.PackageName)
	err := g.serviceTpl.Execute(buf, &goService{obj})
	if err != nil {
		util.Die(err)
	}

	return formatBuffer(buf)
}

func (g *goGen) Init(request *plugin.CodeGeneratorRequest) {
	g.echoGen.Init(request)

	g.structTpl = g.getTpl("/generator/template/go/struct.gogo")
	g.enumTpl = g.getTpl("/generator/template/go/enum.gogo")
}

func (g *goGen) Gen(applicationName string, packageName string, service *data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options data.OptionMap) (result map[string]string, err error) {
	result, err = g.echoGen.Gen(applicationName, packageName, service, messages, enums, options)

	// Temporary hack from go server gen here
	// Should rewrite goGen completely later
	g.serviceTpl = g.getTpl("/generator/template/go/service.gogo")

	filename := genEchoFileName(g.PackageName, service)
	content := g.genGoServie(service)
	result[filename] = content
	return
}

func init() {
	data.OutputMap["go"] = &goGen{}
}
