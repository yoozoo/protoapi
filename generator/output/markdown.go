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
	// get the string value based on the label Type
	getLabel := func(labelType string) string {
		switch labelType {
		case "LABEL_OPTIONAL":
			return "Optional"
		default:
			return "Undefined"
		}
		return "Undefined"
	}

	// get the default value of each data type
	getDefVal := func(fieldType string) string {
		switch fieldType {
		case data.BooleanFieldType:
			return "true"
		case data.DoubleFieldType,
			data.IntFieldType,
			data.Int32FieldType,
			data.Int64FieldType:
			return "0"
		case data.StringFieldType:
			return "Success"
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
	isObject := func(fieldType string) bool {
		switch fieldType {
		case data.StringFieldType,
			data.DoubleFieldType,
			data.IntFieldType,
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

	funcMap := template.FuncMap{
		"getLabel":  getLabel,
		"getDefVal": getDefVal,
		"isRepeat":  isRepeat,
		"isObject":  isObject,
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
