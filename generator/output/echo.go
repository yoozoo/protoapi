package output

import (
	"bytes"
	"strings"
	"text/template"

	"version.uuzu.com/Merlion/protoapi/generator/data"
)

// MethodOptions is the map of method data type and option name in method options
var MethodOptions = map[string]string{
	"ServiceType": "service_method",
	"ErrorType":   "custom_error",
}

func toGoType(dataType string, label string) string {
	// check if the field is repeated
	if label == data.FieldRepeatedLabel {
		return "[]*" + dataType
	}
	// if not primary type return data type and ignore the . in the data type
	return dataType
}

type echoGen struct {
	ApplicationName string
	PackageName     string
	structTpl       *template.Template
	serviceTpl      *template.Template
	enumTpl         *template.Template
}

func newEchoGen(applicationName, packageName string) *echoGen {
	gen := &echoGen{
		ApplicationName: applicationName,
		PackageName:     packageName,
	}
	gen.init()
	return gen
}

func (g *echoGen) getTpl(path string) *template.Template {
	var err error
	tpl := template.New("tpl")
	tplStr := data.LoadTpl(path)
	result, err := tpl.Parse(tplStr)
	if err != nil {
		panic(err)
	}
	return result
}

func (g *echoGen) init() {
	g.structTpl = g.getTpl("/generator/template/echo_struct.gogo")
	g.serviceTpl = g.getTpl("/generator/template/echo_service.gogo")
	g.enumTpl = g.getTpl("/generator/template/echo_enum.gogo")
}

func (g *echoGen) getStructFilename(packageName string, msg *data.MessageData) string {
	return strings.Replace(packageName, ".", "/", -1) + "/" + msg.Name + ".go"
}

func (g *echoGen) genStruct(msg *data.MessageData) string {
	buf := bytes.NewBufferString("")

	obj := newEchoStruct(msg, g.PackageName)
	err := g.structTpl.Execute(buf, obj)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (g *echoGen) getEnumFilename(packageName string, enum *data.EnumData) string {
	return strings.Replace(packageName, ".", "/", -1) + "/" + enum.Name + ".go"
}

func (g *echoGen) genEnum(enum *data.EnumData) string {
	buf := bytes.NewBufferString("")

	obj := newEchoEnum(enum, g.PackageName)
	err := g.enumTpl.Execute(buf, obj)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (g *echoGen) genServie(service *data.ServiceData) string {
	buf := bytes.NewBufferString("")

	obj := newEchoService(service, g.PackageName)
	err := g.serviceTpl.Execute(buf, obj)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func genEchoFileName(packageName string, service *data.ServiceData) string {
	return strings.Replace(packageName, ".", "/", -1) + "/" + service.Name + "Base.go"
}

func genEchoPackageName(packageName string) string {
	return strings.Replace(packageName, ".", "_", -1)
}

func genEchoCode(applicationName string, packageName string, service *data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options []*data.Option) (result map[string]string, err error) {
	packageName = genEchoPackageName(packageName)
	gen := newEchoGen(applicationName, packageName)
	result = make(map[string]string)

	for _, msg := range messages {
		filename := gen.getStructFilename(packageName, msg)
		content := gen.genStruct(msg)

		result[filename] = content
	}

	for _, enum := range enums {
		filename := gen.getEnumFilename(packageName, enum)
		content := gen.genEnum(enum)

		result[filename] = content
	}

	// make file name same as go file name
	filename := genEchoFileName(packageName, service)
	content := gen.genServie(service)
	result[filename] = content

	return
}

func init() {
	data.OutputMap["echo"] = genEchoCode
}
