package output

import (
	"bytes"
	"encoding/json"
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
	Enums    []*data.EnumData
	Time     string
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

	msgMap := make(map[string]*data.MessageData)
	for _, message := range messages {
		msgMap[message.Name] = message
	}

	enumMap := make(map[string]*data.EnumData)
	for _, enum := range enums {
		enumMap[enum.Name] = enum
	}

	fileName := strings.Replace(packageName, "\\", "/", -1)
	if len(fileName) > 0 {
		fileName += "/"
	}
	fileName += service.Name + ".md"

	//读取template文件
	markdownTemplate := tpl.FSMustString(false, "/generator/template/markdown.gomd")

	// check if a field is of type enum
	isEnum := func(fieldType string) bool {
		if _, exist := enumMap[fieldType]; exist {
			return true
		}
		return false
	}

	// return the first enum field name
	getDefEnum := func(fieldType string) string {
		if _, exist := enumMap[fieldType]; exist {
			return enumMap[fieldType].Fields[0].Name
		}
		return ""
	}

	// create template function map
	// get the default value of each data type
	getDefVal := func(fieldType string) string {
		switch fieldType {
		case data.BooleanFieldType:
			return "false"
		case data.DoubleFieldType,
			data.IntFieldType,
			data.Int32FieldType,
			data.Int64FieldType:
			return "0"
		case data.StringFieldType:
			return "Success"
		}

		if isEnum(fieldType) {
			return getDefEnum(fieldType)
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
			return !isEnum(fieldType)
		}
	}

	// get the messageData that matches the messageName and return the fields
	getFields := func(messageName string) []data.MessageField {
		if _, exist := msgMap[messageName]; exist {
			return msgMap[messageName].Fields
		}
		return make([]data.MessageField, 0)
	}

	// return the messageData that matches messageName
	getMessage := func(messageName string) *data.MessageData {
		if _, exist := msgMap[messageName]; exist {
			return msgMap[messageName]
		}
		return &data.MessageData{}
	}

	// filter the messages that is used in the field,
	// return array of message data including the nested messages structure
	var getMessagesOfType func(messageName string, rootName string) []*data.MessageData
	getMessagesOfType = func(messageName string, rootName string) []*data.MessageData {
		var filteredMess []*data.MessageData
		mData := getMessage(messageName)

		filteredMess = append(filteredMess, mData)

		for _, field := range mData.Fields {
			if isMessage(field.DataType) {
				filteredMess = append(filteredMess, getMessagesOfType(field.DataType, rootName)...)
			}
		}
		return filteredMess
	}

	// make a map of string and interface and map message fields to it to be converted into json
	var makeJSONMap func(fields []data.MessageField) map[string]interface{}
	makeJSONMap = func(fields []data.MessageField) map[string]interface{} {
		data := make(map[string]interface{})
		for _, field := range fields {
			if isMessage(field.DataType) {
				data[field.Name] = makeJSONMap(getFields(field.DataType))
			} else {
				data[field.Name] = getDefVal(field.DataType)
			}
		}
		return data
	}

	// convert the fields to map of string and interface and use MarshalIndent to generate Json
	// return the string of the json
	makeJSON := func(fields []data.MessageField) string {
		json, _ := json.MarshalIndent(makeJSONMap(fields), "", "   ")
		return string(json)
	}

	funcMap := template.FuncMap{
		"isEnum":            isEnum,
		"isRepeat":          isRepeat,
		"isMessage":         isMessage,
		"getFields":         getFields,
		"getMessagesOfType": getMessagesOfType,
		"makeJSON":          makeJSON,
	}

	// fill in data
	templateData := markdownStruct{
		Services: service,
		Messages: messages,
		Methods:  service.Methods,
		Enums:    enums,
		Time:     time.Now().Format(time.RFC822),
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
