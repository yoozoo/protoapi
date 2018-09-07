package output

import (
	"bytes"
	"errors"
	"strings"
	"text/template"
	"time"

	"version.uuzu.com/Merlion/protoapi/generator/data"
	"version.uuzu.com/Merlion/protoapi/generator/data/tpl"
)

// create template data struct
type phpStruct struct {
	NameSpace string
	Name      string
	Messages  []*data.MessageData
	Methods   []data.Method
	Enums     []*data.EnumData
	Time      string
	ComErr    *data.MessageData
}

func genPhpCode(applicationName string, packageName string, service *data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options data.OptionMap) (result map[string]string, err error) {
	//获取可能的package name
	if len(packageName) == 0 {
		packageName = "Yoozoo\\Agent"
	}
	nameSpace := strings.Replace(packageName, ".", "\\", -1)

	fileName := strings.Replace(packageName, "\\", "/", -1)
	if len(fileName) > 0 {
		fileName += "/"
	}
	fileName += service.Name + ".php"

	//读取template文件
	phpTemplate := tpl.FSMustString(false, "/generator/template/php.gophp")

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
		"isObject": isObject,
		"isBizErr": isBizErr,
		"isComErr": isComErr,
		"title":    strings.Title,
	}

	// fill in data
	templateData := phpStruct{
		NameSpace: nameSpace,
		Messages:  messages,
		Name:      service.Name,
		Methods:   service.Methods,
		Enums:     enums,
		Time:      time.Now().Format(time.RFC822),
		ComErr:    comError,
	}

	//create a template
	tmpl, err := template.New("php template").Funcs(funcMap).Parse(string(phpTemplate))
	if err != nil {
		return nil, err
	}

	//parse file and generate file content
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
	data.OutputMap["php"] = genPhpCode
}
