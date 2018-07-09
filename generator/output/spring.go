package output

import (
	"bytes"
	"protoapi/generator/data"
	"strings"
	"text/template"
)

var javaTypes = map[string]string{
	// https://developers.google.com/protocol-buffers/docs/proto#scalar
	"double":   "double",
	"float":    "float",
	"int32":    "int",
	"int64":    "long",
	"uint32":   "int",
	"uint64":   "long",
	"sint32":   "int",
	"sint64":   "long",
	"fixed32":  "int",
	"fixed64":  "long",
	"sfixed32": "int",
	"sfixed64": "long",
	"bool":     "boolean",
	"string":   "String",
	"bytes":    "ByteString",
}

// JavaPackageOption is Java package option constant
const JavaPackageOption = "javaPackageOption"

func toJavaType(dataType string) string {
	// check if primary type
	if primaryType, ok := javaTypes[dataType]; ok {
		return primaryType
	}
	// if not primary type return data type and ignore the . in the data type
	return dataType[1:]
}

type springGen struct {
	ApplicationName string
	PackageName     string
	structTpl       *template.Template
	serviceTpl      *template.Template
}

func newSpringGen(applicationName, packageName string) *springGen {
	gen := &springGen{
		ApplicationName: applicationName,
		PackageName:     packageName,
	}
	gen.init()
	return gen
}

func (g *springGen) getTpl(path string) *template.Template {
	var err error
	tpl := template.New("tpl")
	tplStr := data.LoadTpl(path)
	result, err := tpl.Parse(tplStr)
	if err != nil {
		panic(err)
	}
	return result
}

func (g *springGen) init() {
	g.structTpl = g.getTpl("/generator/template/spring_struct.gojava")
	g.serviceTpl = g.getTpl("/generator/template/spring_service.gojava")
}

func (g *springGen) getStructFilename(msg *data.MessageData) string {
	return msg.Name + ".java"
}

func (g *springGen) genStruct(msg *data.MessageData) string {
	buf := bytes.NewBufferString("")

	obj := newSpringStruct(msg, g.PackageName)
	err := g.structTpl.Execute(buf, obj)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (g *springGen) genServie(service *data.ServiceData) string {
	buf := bytes.NewBufferString("")

	obj := newSpringService(service, g.PackageName)
	err := g.serviceTpl.Execute(buf, obj)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func genSpringPackageName(packageName string, options []*data.Option) string {
	if options != nil {
		for _, option := range options {
			if option.Name == JavaPackageOption {
				return option.Value
			}
		}
	}

	return packageName
}

func genSpringCode(applicationName string, packageName string, service *data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options []*data.Option) (result map[string]string, err error) {
	packageName = genSpringPackageName(packageName, options)
	gen := newSpringGen(applicationName, packageName)
	result = make(map[string]string)

	for _, msg := range messages {
		filename := strings.Replace(packageName, ".", "/", -1) + "/" + gen.getStructFilename(msg)
		content := gen.genStruct(msg)

		result[filename] = content
	}

	// make file name same as java class name
	filename := strings.Replace(packageName, ".", "/", -1) + "/" + service.Name + "Base.java"
	content := gen.genServie(service)
	result[filename] = content

	return
}

func init() {
	data.OutputMap["spring"] = genSpringCode
}
