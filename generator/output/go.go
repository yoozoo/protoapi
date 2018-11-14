package output

import (
	"bytes"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/yoozoo/protoapi/generator/data"
	"github.com/yoozoo/protoapi/util"
)

var _goService *goService

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
	return g.ServiceData.CommonErrorType != ""
}

func (g *goService) hasCommonError(field string) bool {
	if !g.HasCommonError() {
		return false
	}

	for _, t := range g.Gen.DataTypes {
		if t.Name == g.ServiceData.CommonErrorType {
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
	_goService = &goService{obj, g}
	err := g.serviceTpl.Execute(buf, _goService)
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
	g.DataTypes = messages

	// Temporary hack from go server gen here
	// Should rewrite goGen completely later
	if service == nil {
		g.serviceTpl = nil
		result, err = g.echoGen.Gen(applicationName, packageName, service, messages, enums, options)
		return
	}

	g.serviceTpl = g.getTpl("/generator/template/go/service.gogo")
	serviceContent := g.genGoServie(service)
	serviceFilename := genEchoFileName(g.PackageName, service)
	g.serviceTpl = nil

	result, err = g.echoGen.Gen(applicationName, packageName, service, messages, enums, options)
	result[serviceFilename] = serviceContent

	return
}

func init() {
	data.OutputMap["go"] = &goGen{}
}
