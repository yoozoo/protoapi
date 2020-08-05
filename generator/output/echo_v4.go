package output

import (
	"github.com/yoozoo/protoapi/generator/data"
)

type echoV4Gen struct {
	goGen
}

func (g *echoV4Gen) Gen(applicationName string, packageName string, services []*data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options data.OptionMap) (result map[string]string, err error) {
	g.DataTypes = messages
	serviceResult := make(map[string]string)
	for _, service := range services {
		g.serviceTpl = g.getTpl("/generator/template/echo_v4/service.gogo")
		serviceContent := g.genGoService(service)
		serviceFilename := genEchoFileName(g.PackageName, service)
		g.serviceTpl = nil
		serviceResult[serviceFilename] = serviceContent
	}

	result, err = g.echoGen.Gen(applicationName, packageName, services, messages, enums, options)
	for k, v := range serviceResult {
		result[k] = v
	}

	return
}

func init() {
	data.OutputMap["echo_v4"] = &echoV4Gen{}
}
