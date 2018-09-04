package output

import (
	"bytes"
	"strings"
	"text/template"

	"version.uuzu.com/Merlion/protoapi/generator/data"
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

var wrapperTypes = map[string]string{
	"double":   "Double",
	"float":    "Float",
	"int32":    "Integer",
	"int64":    "Long",
	"uint32":   "Integer",
	"uint64":   "Long",
	"sint32":   "Integer",
	"sint64":   "Long",
	"fixed32":  "Integer",
	"fixed64":  "Long",
	"sfixed32": "Integer",
	"sfixed64": "Long",
	"bool":     "Boolean",
	"string":   "String",
	"bytes":    "Byte",
}

func toJavaType(dataType string, label string) string {
	// check if the field is repeated
	if label == data.FieldRepeatedLabel {
		// check if wrapper type
		if wrapperType, ok := wrapperTypes[dataType]; ok {
			return "List<" + wrapperType + ">"
		}
		return "List<" + dataType + ">"
	}
	// if not repeated filed
	// check if primary type
	if primaryType, ok := javaTypes[dataType]; ok {
		return primaryType
	}
	// if not primary type return data type and ignore the . in the data type
	return dataType
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

func (g *springGen) getStructFilename(packageName string, msg *data.MessageData) string {
	return strings.Replace(packageName, ".", "/", -1) + "/" + msg.Name + ".java"
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

func genSpringPackageName(packageName string, options data.OptionMap) string {
	if javaPckName, ok := options[data.JavaPackageOption]; ok {
		return javaPckName
	}

	return packageName
}

func genServiceFileName(packageName string, service *data.ServiceData) string {
	return strings.Replace(packageName, ".", "/", -1) + "/" + service.Name + "Base.java"
}

func genSpringCode(applicationName string, packageName string, service *data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options data.OptionMap) (result map[string]string, err error) {
	// get java package name from options
	packageName = genSpringPackageName(packageName, options)
	gen := newSpringGen(applicationName, packageName)
	result = make(map[string]string)

	for _, msg := range messages {
		filename := gen.getStructFilename(packageName, msg)
		content := gen.genStruct(msg)

		result[filename] = content
	}

	// make file name same as java class name
	filename := genServiceFileName(packageName, service)
	content := gen.genServie(service)
	result[filename] = content

	return
}

func init() {
	data.OutputMap["spring"] = genSpringCode
}
