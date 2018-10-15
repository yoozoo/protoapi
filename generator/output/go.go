package output

import (
	"bytes"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/yoozoo/protoapi/generator/data"
	"github.com/yoozoo/protoapi/util"
)

// Re-use everything in echoGen, only use different template
type goGen struct {
	DataTypes []*data.MessageData
	echoGen
}

type goService struct {
	*echoService
	Gen *goGen
}

func (g *goService) CommonError() string {
	return g.ServiceData.Options["common_error"]
}

func (g *goService) HasCommonError() bool {
	_, ok := g.ServiceData.Options["common_error"]
	return ok
}

func (g *goService) hasCommonError(field string) bool {
	errType, ok := g.ServiceData.Options["common_error"]
	if !ok {
		return false
	}

	for _, t := range g.Gen.DataTypes {
		if t.Name == errType {
			for _, f := range t.Fields {
				if f.Name == field {
					return true
				}
			}
		}
	}
	return false
}

func (g *goService) HasCommonBindError() bool {
	return g.hasCommonError("bindError")
}

func (g *goService) HasCommonValidateError() bool {
	return g.hasCommonError("validateError")
}

func (g *goGen) genGoServie(service *data.ServiceData) string {
	buf := bytes.NewBufferString("")

	obj := newEchoService(service, g.PackageName)
	err := g.serviceTpl.Execute(buf, &goService{obj, g})
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
	g.serviceTpl = nil
	g.DataTypes = messages
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
