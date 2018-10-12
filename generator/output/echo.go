package output

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"version.uuzu.com/Merlion/protoapi/generator/data"
	"version.uuzu.com/Merlion/protoapi/util"
)

const (
	googleDescriptorProtoName = "google/protobuf/descriptor.proto"
)

var (
	rgxSyntaxError = regexp.MustCompile(`(\d+):\d+: `)
)

func toGoType(dataType string, label string) string {
	// if not primary type return data type and ignore the . in the data type
	if _, ok := wrapperTypes[dataType]; !ok {
		dataType = "*" + dataType
	}

	// check if the field is repeated
	if label == data.FieldRepeatedLabel {
		dataType = "[]" + dataType
	}

	return dataType
}

type echoGen struct {
	ApplicationName string
	PackageName     string
	structTpl       *template.Template
	serviceTpl      *template.Template
	enumTpl         *template.Template
}

func (g *echoGen) init(applicationName string) {
	g.ApplicationName = applicationName
	g.structTpl = g.getTpl("/generator/template/echo_struct.gogo")
	g.serviceTpl = g.getTpl("/generator/template/echo_service.gogo")
	g.enumTpl = g.getTpl("/generator/template/echo_enum.gogo")
}

func (g *echoGen) getTpl(path string) *template.Template {
	var err error
	tpl := template.New("tpl")
	tplStr := data.LoadTpl(path)
	result, err := tpl.Parse(tplStr)
	if err != nil {
		util.Die(err)
	}
	return result
}

func formatBuffer(buf *bytes.Buffer) (string, error) {
	output, err := format.Source(buf.Bytes())
	if err == nil {
		return string(output), nil
	}

	matches := rgxSyntaxError.FindStringSubmatch(err.Error())
	if matches == nil {
		return "", errors.New("failed to format template")
	}

	lineNum, _ := strconv.Atoi(matches[1])
	scanner := bufio.NewScanner(buf)
	errBuf := &bytes.Buffer{}
	line := 1
	for ; scanner.Scan(); line++ {
		if delta := line - lineNum; delta < -5 || delta > 5 {
			continue
		}

		if line == lineNum {
			errBuf.WriteString(">>>> ")
		} else {
			fmt.Fprintf(errBuf, "% 4d ", line)
		}
		errBuf.Write(scanner.Bytes())
		errBuf.WriteByte('\n')
	}

	return "", fmt.Errorf("failed to format template\n\n%s", errBuf.Bytes())
}

func (g *echoGen) getStructFilename(packageName string, msg *data.MessageData) string {
	return strings.Replace(packageName, ".", "/", -1) + "/" + msg.Name + ".go"
}

func (g *echoGen) genStruct(msg *data.MessageData) string {
	buf := bytes.NewBufferString("")

	obj := newEchoStruct(msg, g.PackageName)
	err := g.structTpl.Execute(buf, obj)
	if err != nil {
		util.Die(err)
	}

	code, err := formatBuffer(buf)
	if err != nil {
		util.Die(err)
	}
	return code
}

func (g *echoGen) getEnumFilename(packageName string, enum *data.EnumData) string {
	return strings.Replace(packageName, ".", "/", -1) + "/" + enum.Name + ".go"
}

func (g *echoGen) genEnum(enum *data.EnumData) string {
	buf := bytes.NewBufferString("")

	obj := newEchoEnum(enum, g.PackageName)
	err := g.enumTpl.Execute(buf, obj)
	if err != nil {
		util.Die(err)
	}

	code, err := formatBuffer(buf)
	if err != nil {
		util.Die(err)
	}
	return code
}

func (g *echoGen) genServie(service *data.ServiceData) string {
	buf := bytes.NewBufferString("")

	obj := newEchoService(service, g.PackageName)
	err := g.serviceTpl.Execute(buf, obj)
	if err != nil {
		util.Die(err)
	}

	code, err := formatBuffer(buf)
	if err != nil {
		util.Die(err)
	}
	return code
}

func genEchoFileName(packageName string, service *data.ServiceData) string {
	return strings.Replace(packageName, ".", "/", -1) + "/" + service.Name + "Base.go"
}

func genEchoPackageName(packageName string) string {
	return strings.Replace(packageName, ".", "_", -1)
}

func (g *echoGen) Init(request *plugin.CodeGeneratorRequest) {
	for _, file := range request.ProtoFile {
		if file.GetName() == googleDescriptorProtoName {
			continue
		}

		opts := file.GetOptions()
		if opts == nil || opts.GetGoPackage() == "" {
			continue
		}

		if g.PackageName == "" {
			g.PackageName = opts.GetGoPackage()
		} else if g.PackageName != opts.GetGoPackage() {
			// Implement code gen for different go packages later
			util.Die(fmt.Errorf("different go package detected: %s, %s", g.PackageName, opts.GetGoPackage()))
		}
	}
}

func (g *echoGen) Gen(applicationName string, packageName string, service *data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options data.OptionMap) (result map[string]string, err error) {
	if g.PackageName == "" {
		g.PackageName = genEchoPackageName(packageName)
	}

	g.init(applicationName)
	result = make(map[string]string)

	for _, msg := range messages {
		filename := g.getStructFilename(g.PackageName, msg)
		content := g.genStruct(msg)

		result[filename] = content
	}

	for _, enum := range enums {
		filename := g.getEnumFilename(g.PackageName, enum)
		content := g.genEnum(enum)

		result[filename] = content
	}

	// make file name same as go file name
	filename := genEchoFileName(g.PackageName, service)
	content := g.genServie(service)
	result[filename] = content

	return
}

func init() {
	gen := &echoGen{}
	data.OutputMap["echo"] = gen
	data.OutputMap["go"] = gen
}
