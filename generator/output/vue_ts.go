package output

import (
	"bytes"
	"log"
	"protoapi/generator/data"
	"strings"
	"text/template"
)

// create template data struct
type vueResource struct {
	ClassName    string
	DataTypes    []*data.MessageData
	DataTypeFile string
	Functions    []data.Method
}

type vueInterface struct {
	DataTypes []*data.MessageData
}

func generateFuncName(title string) string {
	log.Printf("title is %s\n", title)
	titles := strings.Split(title, "/")
	result := "get"
	for _, t := range titles {
		result += strings.Title(t)
	}
	log.Printf("result is %s\n", result)
	return result
}

func isGet(function string) bool {
	return strings.Contains(function, "Get")
}

func generateVueTsCode(applicationName string, packageName string, service *data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options []*data.Option) (map[string]string, error) {

	// Code file
	serviceFile := strings.Replace(packageName, ".", "/", -1)
	serviceFile = serviceFile + "/" + service.Name + ".ts"

	// Data Struct file
	dataFile := strings.Replace(packageName, ".", "/", -1)
	dataFile = dataFile + "/" + strings.Title(strings.Replace(applicationName, ".proto", "", -1)) + ".ts"
	log.Printf("dataFile is %s\n", dataFile)

	// Helper file
	helperFile := strings.Replace(packageName, ".", "/", -1) + "/Helper.ts"

	// tsconfig.json file
	tsConfigFile := strings.Replace(packageName, ".", "/", -1) + "/tsconfig.json"

	// Get template path: one for class generation，one for data type （interface） generation
	vueTpl := data.LoadTpl("/generator/template/vue.gots")
	interfaceTpl := data.LoadTpl("/generator/template/interface.gots")
	helperTpl := data.LoadTpl("/generator/template/helper.gots")
	tsConfigTpl := data.LoadTpl("/generator/template/tsconfig.gojson")

	// map messages and service
	serviceData := vueResource{
		ClassName:    service.Name,
		DataTypeFile: strings.Title(strings.Replace(applicationName, ".proto", "", -1)),
		DataTypes:    messages,
		Functions:    service.Methods,
	}

	interfaceData := vueInterface{
		DataTypes: messages,
	}

	// function map
	funcMap := template.FuncMap{
		"Title": generateFuncName,
		"isGet": strings.Contains,
	}

	// create templates
	tmpl, err := template.New("hello").Funcs(funcMap).Parse(vueTpl)
	if err != nil {
		return nil, err
	}
	tmpl2, err := template.New("data").Funcs(funcMap).Parse(interfaceTpl)
	if err != nil {
		return nil, err
	}

	// combine data with template
	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, serviceData)

	buf2 := bytes.NewBufferString("")
	err = tmpl2.Execute(buf2, interfaceData)

	serviceContent := buf.String()
	dataContent := buf2.String()

	if err != nil {
		return nil, err
	}

	var result = make(map[string]string)

	// append generated file in result
	result[serviceFile] = serviceContent
	result[dataFile] = dataContent
	result[helperFile] = helperTpl
	result[tsConfigFile] = tsConfigTpl

	return result, nil
}

func init() {
	data.OutputMap["ts"] = generateVueTsCode
}
