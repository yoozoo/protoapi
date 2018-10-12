package output

import (
	"bytes"
	"strings"
	"text/template"

	"version.uuzu.com/Merlion/protoapi/generator/data"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
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
	DataTypes []*data.MessageData
	Lib       tsLibs

	objsFile   string
	helperFile string

	axiosFile string
	fetchFile string

	objsTpl   *template.Template
	helperTpl *template.Template

	axiosTpl *template.Template
	fetchTpl *template.Template

	service *data.ServiceData
}

type tsStruct struct {
	ClassName string
	DataTypes []*data.MessageData
	Enums     []*data.EnumData
	Functions []data.Method
	Gen       *tsGen
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
	g.axiosTpl = g.getTpl("/generator/template/ts/service_axios.gots")
	g.fetchTpl = g.getTpl("/generator/template/ts/service_fetch.gots")
	g.objsTpl = g.getTpl("/generator/template/ts/objs.gots")
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

func (g *tsGen) CommonError() string {
	return g.service.Options["common_error"]
}

func (g *tsGen) CommonErrorSubTypes() string {
	var fieldTypes []string
	for _, f := range g.GetCommoneErrorFields() {
		subType := toTypeScriptType(f.DataType)
		fieldTypes = append(fieldTypes, " | "+subType)
	}

	return strings.Join(fieldTypes, "")
}

func (g *tsGen) GetCommoneErrorFields() []data.MessageField {
	commonErrorType := g.service.Options["common_error"]
	for _, t := range g.DataTypes {
		if t.Name == commonErrorType {
			return t.Fields
		}
	}
	return nil
}

func (g *tsGen) HasCommonError() bool {
	_, ok := g.service.Options["common_error"]
	return ok
}

/**
* init filename with path
 */
func (g *tsGen) initFiles(packageName string, service *data.ServiceData) {
	g.axiosFile = genFileName(packageName, service.Name)
	g.fetchFile = genFileName(packageName, service.Name)
	g.objsFile = genFileName(packageName, service.Name+"Objs")
	g.helperFile = genFileName(packageName, "helper")
	g.service = service
}

type tsLibs int

const (
	tsLibFetch tsLibs = iota
	tsLibAxios
)

func (g *tsGen) Init(request *plugin.CodeGeneratorRequest) {
	g.loadTpl()
}

func (g *tsGen) Gen(applicationName string, packageName string, svr *data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options data.OptionMap) (map[string]string, error) {
	g.initFiles(packageName, svr)
	g.DataTypes = messages

	/**
	* Map Data: messages and service
	 */
	dataMap := tsStruct{
		ClassName: svr.Name,
		DataTypes: messages,
		Enums:     enums,
		Functions: svr.Methods,
		Gen:       g,
	}

	var result = make(map[string]string)
	switch g.Lib {
	case tsLibAxios:
		result[g.axiosFile] = g.genContent(g.axiosTpl, dataMap)
	default:
		result[g.fetchFile] = g.genContent(g.fetchTpl, dataMap)
	}

	result[g.objsFile] = g.genContent(g.objsTpl, dataMap)
	result[g.helperFile] = g.genContent(g.helperTpl, dataMap)

	return result, nil
}

func getTSgen(lib tsLibs) *tsGen {
	g := new(tsGen)
	g.Lib = lib
	return g
}

func init() {
	fetch := getTSgen(tsLibFetch)
	data.OutputMap["ts"] = fetch
	data.OutputMap["ts-fetch"] = fetch
	data.OutputMap["ts-axios"] = getTSgen(tsLibAxios)
}
