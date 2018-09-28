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

	"version.uuzu.com/Merlion/protoapi/generator/data"
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
		die(err)
	}
	return result
}

func (g *echoGen) init() {
	g.structTpl = g.getTpl("/generator/template/echo_struct.gogo")
	g.serviceTpl = g.getTpl("/generator/template/echo_service.gogo")
	g.enumTpl = g.getTpl("/generator/template/echo_enum.gogo")
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

	return "", fmt.Errorf("failed to format template\n\n%s\n", errBuf.Bytes())
}

func die(err error) {
	print(err.Error())
	panic(err)
}

func (g *echoGen) getStructFilename(packageName string, msg *data.MessageData) string {
	return strings.Replace(packageName, ".", "/", -1) + "/" + msg.Name + ".go"
}

func (g *echoGen) genStruct(msg *data.MessageData) string {
	buf := bytes.NewBufferString("")

	obj := newEchoStruct(msg, g.PackageName)
	err := g.structTpl.Execute(buf, obj)
	if err != nil {
		die(err)
	}

	code, err := formatBuffer(buf)
	if err != nil {
		die(err)
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
		die(err)
	}

	code, err := formatBuffer(buf)
	if err != nil {
		die(err)
	}
	return code
}

func (g *echoGen) genServie(service *data.ServiceData) string {
	buf := bytes.NewBufferString("")

	obj := newEchoService(service, g.PackageName)
	err := g.serviceTpl.Execute(buf, obj)
	if err != nil {
		die(err)
	}

	code, err := formatBuffer(buf)
	if err != nil {
		die(err)
	}
	return code
}

func genEchoFileName(packageName string, service *data.ServiceData) string {
	return strings.Replace(packageName, ".", "/", -1) + "/" + service.Name + "Base.go"
}

func genEchoPackageName(packageName string) string {
	return strings.Replace(packageName, ".", "_", -1)
}

func genEchoCode(applicationName string, packageName string, service *data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options data.OptionMap) (result map[string]string, err error) {
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
