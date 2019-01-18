package output

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"strings"
	"text/template"
	"time"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/yoozoo/protoapi/generator/data"
	"github.com/yoozoo/protoapi/generator/data/tpl"
	"github.com/yoozoo/protoapi/util"
)

// create template data struct
type goStruct struct {
	Package  string
	Name     string
	Messages []*data.MessageData
	Methods  []*data.Method
	Enums    []*data.EnumData
	Time     string
	ComErr   *data.MessageData
}

type goClientGen struct{}

func (g *goClientGen) Init(request *plugin.CodeGeneratorRequest) {
}

func (g *goClientGen) Gen(applicationName string, packageName string, services []*data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options data.OptionMap) (result map[string]string, err error) {
	var service *data.ServiceData
	if len(services) > 1 {
		util.Die(fmt.Errorf("goclient found %d services; only 1 service is supported now", len(services)))
	} else if len(services) == 1 {
		service = services[0]
	}

	//获取可能的package name
	if len(packageName) == 0 {
		packageName = "yoozooagent"
	}
	nameSpace := strings.Replace(packageName, ".", "", -1)

	fileName := strings.Replace(packageName, "\\", "/", -1)
	if len(fileName) > 0 {
		fileName += "/"
	}
	fileName += service.Name + ".go"

	//读取template文件
	goTemplate := tpl.FSMustString(false, "/generator/template/go_client.gogo")

	// create template function map
	bizErrorMsgs := make(map[string]bool)
	for _, serv := range service.Methods {
		errorMsgName, found := serv.Options["error"]
		if found {
			bizErrorMsgs[errorMsgName] = true
		}
	}
	isBizErr := func(name string) bool {
		for k := range bizErrorMsgs {
			if k == name {
				return true
			}
		}
		return false
	}

	var comError *data.MessageData
	for i, msg := range messages {
		if msg.Name == data.ComErrMsgName {
			comError = msg
			messages = append(messages[:i], messages[i+1:]...)
			break
		}
	}
	if comError == nil {
		return nil, errors.New("Cannot find common error message")
	}
	isComErr := func(name string) bool {
		for _, field := range comError.Fields {
			if field.DataType == name {
				return true
			}
		}
		return false
	}

	isObject := func(fieldType string) bool {
		switch fieldType {
		case data.StringFieldType,
			data.DoubleFieldType,
			data.IntFieldType,
			data.Int32FieldType,
			data.Int64FieldType,
			data.BooleanFieldType:
			return false
		default:
			// check if is enum
			for _, enum := range enums {
				if enum.Name == fieldType {
					return false
				}
			}
			return true
		}
	}

	toType := func(s *data.MessageField) string {
		dataType := s.DataType
		// if not primary type return data type and ignore the . in the data type
		if isObject(dataType) {
			dataType = "*" + dataType
		}

		// check if the field is repeated
		if s.Label == data.FieldRepeatedLabel {
			dataType = "[]" + dataType
		}

		return dataType
	}

	funcMap := template.FuncMap{
		"isObject": isObject,
		"isBizErr": isBizErr,
		"isComErr": isComErr,
		"title":    strings.Title,
		"type":     toType,
	}

	// fill in data
	templateData := goStruct{
		Package:  nameSpace,
		Messages: messages,
		Name:     service.Name,
		Methods:  service.Methods,
		Enums:    enums,
		Time:     time.Now().Format(time.RFC822),
		ComErr:   comError,
	}

	//create a template
	tmpl, err := template.New("go client template").Funcs(funcMap).Parse(string(goTemplate))
	if err != nil {
		return nil, err
	}

	//parse file and generate file content
	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, templateData)
	if err != nil {
		return nil, err
	}
	fileContent, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, err
	}

	result = make(map[string]string)
	result[fileName] = string(fileContent)
	return result, nil
}

func init() {
	data.OutputMap["goclient"] = &goClientGen{}
}
