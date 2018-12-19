package output

import (
	"bytes"
	"strings"
	"text/template"
	"time"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/yoozoo/protoapi/generator/data"
	"github.com/yoozoo/protoapi/generator/data/tpl"
)

// create template data struct
type markdownStruct struct {
	Services *data.ServiceData
	Messages []*data.MessageData
	Methods  []*data.Method
	Time     string
	//ComErr   *data.MessageData
}

//contains logic to plug in values to the template specified
type markdownGen struct{}

func (g *markdownGen) Init(request *plugin.CodeGeneratorRequest) {
}

func (g *markdownGen) Gen(applicationName string, packageName string, service *data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options data.OptionMap) (result map[string]string, err error) {
	//获取可能的package name
	if len(packageName) == 0 {
		packageName = "Yoozoo\\Agent"
	}

	fileName := strings.Replace(packageName, "\\", "/", -1)
	if len(fileName) > 0 {
		fileName += "/"
	}
	fileName += service.Name + ".md"

	//读取template文件
	markdownTemplate := tpl.FSMustString(false, "/generator/template/markdown.gomd")

	// create template function map
	// get the default value of each data type
	getDefVal := func(dataType string) string {
		switch dataType {
		case data.BooleanFieldType:
			return "false"
		case data.DoubleFieldType,
			data.IntFieldType,
			data.Int32FieldType,
			data.Int64FieldType:
			return "0"
		case data.StringFieldType:
			return "\"Success\""
		}
		return ""
	}

	// check if the field label is repeated
	isRepeat := func(labelType string) bool {
		if labelType == data.FieldRepeatedLabel {
			return true
		}
		return false
	}

	// return false for primitive data type and enum
	isMessage := func(fieldType string) bool {
		switch fieldType {
		case data.StringFieldType,
			data.DoubleFieldType,
			data.IntFieldType,
			data.BooleanFieldType:
			return false
		default:
			// check if it is enum
			for _, enum := range enums {
				if enum.Name == fieldType {
					return false
				}
			}
			return true
		}
	}

	// get the messageData that matches the datatype and return the fields
	getFields := func(fieldType string) []data.MessageField {
		for _, message := range messages {
			if message.Name == fieldType {
				return message.Fields
			}
		}
		return make([]data.MessageField, 0)
	}

	//check if a field is not the last field in message
	isNotLast := func(fieldName string, fields []data.MessageField) bool {
		return fieldName != fields[len(fields)-1].Name
	}

	//check if a method is part of an input parameter or an output parameter
	var isOfType func(messageName string, typeName string) bool
	isOfType = func(messageName string, typeName string) bool {
		for _, field := range getFields(typeName) {
			if isMessage(field.DataType) {
				return isOfType(messageName, field.DataType)
			} else if messageName == typeName {
				return true
			}
		}
		return false
	}

	funcMap := template.FuncMap{
		"getDefVal": getDefVal,
		"isRepeat":  isRepeat,
		"isMessage": isMessage,
		"getFields": getFields,
		"isNotLast": isNotLast,
		"isOfType":  isOfType,
	}

	// var comError *data.MessageData
	// for i, msg := range messages {
	// 	if msg.Name == data.ComErrMsgName {
	// 		comError = msg
	// 		messages = append(messages[:i], messages[i+1:]...)
	// 		break
	// 	}
	// }
	// if comError == nil {
	// 	return nil, errors.New("Cannot find common error message")
	// }

	// fill in data
	templateData := markdownStruct{
		Services: service,
		Messages: messages,
		Methods:  service.Methods,
		Time:     time.Now().Format(time.RFC822),
		//ComErr:   comError,
	}

	//create a template
	tmpl, err := template.New("markdown template").Funcs(funcMap).Parse(string(markdownTemplate))
	if err != nil {
		return nil, err
	}

	//parse file and generate file content according to the template
	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, templateData)
	if err != nil {
		return nil, err
	}
	fileContent := buf.String()

	result = make(map[string]string)
	result[fileName] = fileContent
	return result, nil
}

func init() {
	data.OutputMap["markdown"] = &markdownGen{}
}
