package output

import (
	"bytes"
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
}

func genPhpCode(applicationName string, packageName string, service *data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options []*data.Option) (result map[string]string, err error) {
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
	isObject := func(fieldType string) bool {
		switch fieldType {
		case data.StringFieldType,
			data.DoubleFieldType,
			data.IntFieldType,
			data.BooleanFieldType:
			return false
		default:
			return true //ignore enum currently
		}
	}

	funcMap := template.FuncMap{
		"isObject": isObject,
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
