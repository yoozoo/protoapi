package output

import (
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"
	"protoapi/generator/data"
	"strings"
	"text/template"
)

// create template data struct
type vueResource struct {
	ClassName    string
	DataTypes    []*data.MessageData
	DataTypeFile string
	FunctionName string
	Request      string
	Response     string
}

type vueInterface struct {
	DataTypes []*data.MessageData
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
	serviceFile := strings.Replace(packageName, ".", "/", -1)
	serviceFile = serviceFile + "/" + service.Name + ".ts"

	// Data Struct file
	dataFile := strings.Replace(packageName, ".", "/", -1)
	dataFile = dataFile + "/" + strings.Title(strings.Replace(applicationName, ".proto", "", -1)) + ".ts"
	log.Printf("dataFile is %s\n", dataFile)

	// Get template path: one for class generation，one for data type （interface） generation
	vuePath, err := filepath.Abs("generator/template/vue.gots")
	if err != nil {
		return nil, err
	}

	interfacePath, err := filepath.Abs("generator/template/interface.gots")
	if err != nil {
		return nil, err
	}

	helperPath, err := filepath.Abs("generator/template/helper.gots")
	if err != nil {
		return nil, err
	}

	// read templates
	vueTemplate, err := ioutil.ReadFile(vuePath)
	interfaceTemplate, err := ioutil.ReadFile(interfacePath)
	if err != nil {
		return nil, err
	}

	// map messages and service
	serviceData := vueResource{
		ClassName:    service.Name,
		DataTypeFile: strings.Title(strings.Replace(applicationName, ".proto", "", -1)),
		DataTypes:    messages,
		FunctionName: service.Methods[0].Name,
		Request:      service.Methods[0].InputType,
		Response:     service.Methods[0].OutputType,
	}

	interfaceData := vueInterface{
		DataTypes: messages,
	}

	// function map
	funcMap := template.FuncMap{
		"Title": generateFuncName,
	}

	// create templates
	tmpl, err := template.New("hello").Funcs(funcMap).Parse(string(vueTemplate))
	if err != nil {
		return nil, err
	}
	tmpl2, err := template.New("data").Funcs(funcMap).Parse(string(interfaceTemplate))
	if err != nil {
		return nil, err
	}
	// tmpl3, err := ioutil.ReadFile(helperPath)

	// combine data with template
	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, serviceData)

	buf2 := bytes.NewBufferString("")
	err = tmpl2.Execute(buf2, interfaceData)

	serviceContent := buf.String()
	dataContent := buf2.String()
	helperContent, err := ioutil.ReadFile(helperPath)

	if err != nil {
		return nil, err
	}

	var result = make(map[string]string)

	// append generated file in result
	result[serviceFile] = serviceContent
	result[dataFile] = dataContent
	result["com/yoozoo/ts/Helper.ts"] = string(helperContent)

	return result, nil
}

func init() {
	data.OutputMap["ts"] = generateVueTsCode
}
