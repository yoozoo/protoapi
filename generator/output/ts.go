package output

import (
	"bytes"
	"strings"
	"text/template"

	"version.uuzu.com/Merlion/protoapi/generator/data"
)

/**
*  Map go type to ts types
 */
var tsTypes = map[string]string{
	"int":      "number",
	"double":   "number",
	"float":    "number",
	"int32":    "number",
	"int64":    "number",
	"uint32":   "number",
	"uint64":   "number",
	"sint32":   "number",
	"sint64":   "number",
	"fixed32":  "number",
	"fixed64":  "number",
	"sfixed32": "number",
	"sfixed64": "number",
	"bool":     "boolean",
	"string":   "string",
}

type tsGen struct {
	vueResourceFile string
	axiosFile       string
	dataFile        string
	helperFile      string
	vueResourceTpl  *template.Template
	axiosTpl        *template.Template
	dataTpl         *template.Template
	helperTpl       *template.Template
}

type tsStruct struct {
	ClassName string
	DataTypes []*data.MessageData
	Enums     []*data.EnumData
	Functions []data.Method
}

func toTypeScriptType(dataType string) string {
	if primaryType, ok := tsTypes[dataType]; ok {
		return primaryType
	}
	return dataType
}

func getErrorType(options data.OptionMap) string {
	if errType, ok := options["error"]; ok {
		return errType
	}

	return ""
}

func getServiceMtd(options data.OptionMap) string {
	if servMtd, ok := options["service_method"]; ok {
		return servMtd
	}

	return "POST"
}

func getImportDataTypes(mtds []data.Method) map[string]bool {
	res := make(map[string]bool)
	res["Error"] = true
	for _, mtd := range mtds {
		_, exist := res[mtd.InputType]
		if !exist {
			res[mtd.InputType] = true
		}
		_, exist = res[mtd.OutputType]
		if !exist {
			res[mtd.OutputType] = true
		}
	}
	return res
}

func genFileName(packageName string, fileName string) string {
	return fileName + ".ts"
}

/**
* Get TEMPLATE
 */
func (g *tsGen) loadTpl() {
	g.vueResourceTpl = g.getTpl("/generator/template/ts/ts_service.govue")
	g.axiosTpl = g.getTpl("/generator/template/ts/ts_service.gots")
	g.dataTpl = g.getTpl("/generator/template/ts/data.gots")
	g.helperTpl = g.getTpl("/generator/template/ts/helper.gots")
}

/**
* Parse TEMPLATE
 */
func (g *tsGen) getTpl(path string) *template.Template {
	var funcs = template.FuncMap{
		"tsType":             toTypeScriptType,
		"toLower":            strings.ToLower,
		"getErrorType":       getErrorType,
		"getServiceMtd":      getServiceMtd,
		"getImportDataTypes": getImportDataTypes,
	}
	var err error
	tpl := template.New("tpl").Funcs(funcs)
	tplStr := data.LoadTpl(path)
	result, err := tpl.Parse(tplStr)
	if err != nil {
		panic(err)
	}
	return result
}

/**
* load CONTENT into TEMPLATE
 */
func (g *tsGen) genContent(tpl *template.Template, data tsStruct) string {
	buf := bytes.NewBufferString("")
	err := tpl.Execute(buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

/**
* init filename with path
 */
func initFiles(packageName string, service *data.ServiceData) *tsGen {
	gen := &tsGen{
		vueResourceFile: genFileName(packageName, service.Name),
		axiosFile:       genFileName(packageName, "api"),
		dataFile:        genFileName(packageName, "data"),
		helperFile:      genFileName(packageName, "helper"),
	}
	return gen
}

func generateVueTsCode(applicationName string, packageName string, service *data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options data.OptionMap) (map[string]string, error) {
	/**
	* name files
	 */
	gen := initFiles(packageName, service)

	/**
	* prep template
	 */
	gen.loadTpl()

	/**
	* Map Data: messages and service
	 */
	dataMap := tsStruct{
		ClassName: service.Name,
		DataTypes: messages,
		Enums:     enums,
		Functions: service.Methods,
	}

	/**
	* combine data with template
	 */
	var result = make(map[string]string)
	result[gen.vueResourceFile] = gen.genContent(gen.vueResourceTpl, dataMap)
	result[gen.axiosFile] = gen.genContent(gen.axiosTpl, dataMap)
	result[gen.dataFile] = gen.genContent(gen.dataTpl, dataMap)
	result[gen.helperFile] = data.LoadTpl("/generator/template/ts/helper.gots")
	return result, nil
}

func init() {
	data.OutputMap["ts"] = generateVueTsCode
}
