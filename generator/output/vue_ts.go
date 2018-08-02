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

	// base path of files generated
	prefix := strings.Replace(packageName, ".", "/", -1)

	/** 
	* Generate code files in /src folder 
	*/
	// Service file
	serviceFile := prefix + "/src/" + service.Name  + ".ts"

	// Data Struct file
	dataFile := prefix + "/src/" + strings.Title(strings.Replace(applicationName, ".proto", "", -1)) + ".ts"

	// Helper file
	helperFile := prefix + "/src/Helper.ts"

	// index html & ts for testing API functions
	// indexHtmlFile := prefix + "/src/index.html"
	// indexTSFile := prefix + "/src/index.ts"

	/** 
	* Other configurations files 
	*/
	// tsConfigFile := prefix + "/tsconfig.json"
	// webpackConfigFile := prefix + "/webpack.config.js"
	// babelConfigFile := prefix + "/babel.config.js"
	// packageFile := prefix + "/package.json"
	// publicIndexFile := prefix + "/public/index.html"
	// readMeFile := prefix + "/README.md"

	/**
	* Get TEMPLATE path
	*/
	vueTpl := data.LoadTpl("/generator/template/ts/vue.gots")
	interfaceTpl := data.LoadTpl("/generator/template/ts/interface.gots")
	helperTpl := data.LoadTpl("/generator/template/ts/helper.gots")
	// tsConfigTpl := data.LoadTpl("/generator/template/ts/tsconfig.gojson")
	
	// indexHtmlTpl := data.LoadTpl("/generator/template/ts/index.gohtml")
	// indexTsTpl := data.LoadTpl("/generator/template/ts/index.gots")

	// pkgTpl := data.LoadTpl("/generator/template/ts/package.gojson")
	// webpackConfigTpl := data.LoadTpl("/generator/template/ts/webpack.config.gojs")
	// publicIndexTpl := data.LoadTpl("/generator/template/ts/public_index.gohtml")
	// babelConfigTpl := data.LoadTpl("/generator/template/ts/babel.config.gojs")
	// readMeTpl := data.LoadTpl("/generator/template/ts/README.md")

	/**
	* Map Data: messages and service
	*/
	serviceData := vueResource{
		ClassName:    service.Name,
		DataTypeFile: strings.Title(strings.Replace(applicationName, ".proto", "", -1)),
		DataTypes:    messages,
		Functions:    service.Methods,
	}

	interfaceData := vueInterface{
		DataTypes: messages,
	}

	/** 
	* function map 
	*/
	funcMap := template.FuncMap{
		"Title": generateFuncName,
		"isGet": strings.Contains,
	}

	/**
	* create necessary templates
	*/
	tmpl, err := template.New("service").Funcs(funcMap).Parse(vueTpl)
	if err != nil {
		return nil, err
	}
	tmpl2, err := template.New("datatype").Funcs(funcMap).Parse(interfaceTpl)
	if err != nil {
		return nil, err
	}
	// tmpl3, err := template.New("index html").Funcs(funcMap).Parse(indexHtmlTpl)
	// if err != nil {
	// 	return nil, err
	// }
	// tmpl4, err := template.New("index ts").Funcs(funcMap).Parse(indexTsTpl)
	// if err != nil {
	// 	return nil, err
	// }

	/**
	* combine data with template
	*/
	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, serviceData)

	buf2 := bytes.NewBufferString("")
	err = tmpl2.Execute(buf2, interfaceData)

	// buf3 := bytes.NewBufferString("")
	// err = tmpl3.Execute(buf3, serviceData)

	// buf4 := bytes.NewBufferString("")
	// err = tmpl4.Execute(buf4, serviceData)

	serviceContent := buf.String()
	dataContent := buf2.String()
	// indexHtml := buf3.String()
	// indexTS := buf4.String()

	if err != nil {
		return nil, err
	}

	var result = make(map[string]string)

	// append generated file in result
	result[serviceFile] = serviceContent
	result[dataFile] = dataContent
	// result[indexHtmlFile] = indexHtml
	// result[indexTSFile] = indexTS
	result[helperFile] = helperTpl
	// // config files
	// result[tsConfigFile] = tsConfigTpl
	// result[webpackConfigFile] = webpackConfigTpl
	// result[packageFile] = pkgTpl
	// result[babelConfigFile] = babelConfigTpl
	// result[publicIndexFile] = publicIndexTpl
	// result[readMeFile] = readMeTpl
	
	return result, nil
}

func init() {
	data.OutputMap["ts"] = generateVueTsCode
}
