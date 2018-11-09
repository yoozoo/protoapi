package output

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"
	"text/template"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/yoozoo/protoapi/generator/data"
	yii2 "github.com/yoozoo/protoapi/generator/output/phpyii2"
	"github.com/yoozoo/protoapi/util"
)

type yii2Gen struct {
	result            map[string]string
	ModuleName        string
	NameSpace         string
	bizErrors         []string
	comError          *data.MessageData
	messageTpl        *template.Template
	errorTpl          *template.Template
	controllerTpl     *template.Template
	ErrorHandlerTpl   *template.Template
	RequestHandlerTpa *template.Template
	moduleTpl         *template.Template
}

func (g *yii2Gen) getTpl(path string) (*template.Template, error) {
	var err error
	tpl := template.New("tpl")
	tplStr := data.LoadTpl(path)
	result, err := tpl.Parse(tplStr)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (g *yii2Gen) genMessage(msg *data.MessageData) error {
	buf := bytes.NewBufferString("")

	obj := yii2.NewMessage(msg, g.NameSpace)
	err := g.messageTpl.Execute(buf, obj)
	if err != nil {
		return err
	}

	return nil
}

func (g *yii2Gen) Init(request *plugin.CodeGeneratorRequest) {
	for _, file := range request.ProtoFile {
		if file.GetName() == googleDescriptorProtoName {
			continue
		}

		opts := file.GetOptions()
		if opts == nil || opts.GetPhpNamespace() == "" {
			continue
		}

		if g.NameSpace == "" {
			g.NameSpace = opts.GetPhpNamespace()
		}
	}

	g.messageTpl, _ = g.getTpl("/generator/template/echo_struct.gogo")
}

func (g *yii2Gen) isBizErr(msg *data.MessageData) bool {
	for _, field := range g.bizErrors {
		if field == msg.Name {
			return true
		}
	}
	return false
}

func (g *yii2Gen) isComErr(msg *data.MessageData) bool {
	for _, field := range g.comError.Fields {
		if field.DataType == msg.Name {
			return true
		}
	}
	return false
}

func (g *yii2Gen) Gen(applicationName string, packageName string, service *data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options data.OptionMap) (result map[string]string, err error) {
	if g.NameSpace == "" {
		g.NameSpace = strings.Replace(packageName, ".", "\\", -1)

		if g.NameSpace == "" {
			util.Die(fmt.Errorf("No name space given"))
		}

		log.Printf("Use proto package name for php: %v", g.NameSpace)
	}

	g.ModuleName = applicationName
	g.result = make(map[string]string)

	// create error map
	for _, serv := range service.Methods {
		errorMsgName, found := serv.Options["error"]
		if found {
			g.bizErrors = append(g.bizErrors, errorMsgName)
		}
	}

	for i, msg := range messages {
		if msg.Name == data.ComErrMsgName {
			g.comError = msg
			messages = append(messages[:i], messages[i+1:]...)
			break
		}
	}
	if g.comError == nil {
		return nil, errors.New("Cannot find common error message")
	}

	for _, msg := range messages {
		if g.isBizErr(msg) {

		} else if g.isComErr(msg) {

		} else {

			filename := g.getStructFilename(g.PackageName, msg)
			content := g.genStruct(msg)

			result[filename] = content
		}
	}

	if g.serviceTpl != nil {
		filename := genEchoFileName(g.PackageName, service)
		content := g.genServie(service)
		result[filename] = content
	}

	return
}

func init() {
	data.OutputMap["yii2"] = &echoGen{}
}
