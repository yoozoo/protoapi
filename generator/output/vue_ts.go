package output

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"protoapi/generator/data"
	"strings"
	"text/template"
)

// create template data struct
type VueResource struct {
	ClassName    string
	DataTypes    []*data.MessageData
	DataTypeFile string
	FunctionName string
	Request      string
	Response     string
}

func generateFuncName(title string) string {
	titles := strings.Split(title, "/")
	result := "get"
	for _, t := range titles {
		result += strings.Title(t)
	}
	return result
}

func generateVueTsCode(applicationName string, packageName string, service *data.ServiceData, messages []*data.MessageData, enums []*data.EnumData) (map[string]string, error) {

	// Code file
	fileName := strings.Replace(packageName, ".", "/", -1)
	fileName = fileName + "/" + service.Name + ".ts"

	// Data Struct file
	dataFile := strings.Title(strings.Replace(applicationName, ".proto", "", -1)) + ".ts"
	fmt.Fprintf(os.Stderr, "dataFile is %s\n", dataFile)

	//读取template文件: 一个是class generation， 一个是 data type （interface） generation
	vue_path, err := filepath.Abs("generator/template/vue_ts.tmpl")
	// interface_path, err := filepath.Abs("generator/template/interface_ts.tmpl")

	if err != nil {
		return nil, err
	}
	vue_template, err := ioutil.ReadFile(vue_path)
	// interface_template, err := ioutil.ReadFile(interface_path)
	if err != nil {
		return nil, err
	}

	//获取message， 转化成需要的格式

	vue_strut := VueResource{
		ClassName:    service.Name,
		DataTypeFile: dataFile,
		DataTypes:    messages,
		FunctionName: service.Methods[0].Name,
		Request:      service.Methods[0].InputType,
		Response:     service.Methods[0].OutputType,
	}

	//创建function map
	funcMap := template.FuncMap{
		"Title": generateFuncName,
	}

	tmpl, err := template.New("hello").Funcs(funcMap).Parse(string(vue_template)) //建立一个模板
	// tmpl2, err := template.New("data").Funcs(funcMap).Parse(string(interface_template)) //建立一个模板
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, vue_strut) //将struct与模板合成，合成结果放到buffer里
	// err2 := tmpl.Execute(buf, *messages)
	if err != nil {
		return nil, err
	}
	fileContent := buf.String()
	var result = make(map[string]string)
	result[fileName] = fileContent

	// dataContent := buf.String()
	// result[dataFile] = dataContent
	return result, nil
}

func init() {
	data.OutputMap["ts"] = generateVueTsCode
}
