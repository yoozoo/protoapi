package output

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/yoozoo/protoapi/generator/data"
	"github.com/yoozoo/protoapi/util"
)

var _goService *goService

// Re-use everything in echoGen, only use different template
type goGen struct {
	DataTypes []*data.MessageData
	echoGen
}

type goService struct {
	*echoService
	Gen *goGen
}

// GetGoPackageAndType convert proto data type like protoapi.common.error to common.Error
// and return its package name in go
func getGoPackageAndType(dataType string) (isFileToGenerate bool, pkg, refType string) {
	_, p := data.GetMessageProtoAndFile(dataType)
	isFileToGenerate = p.IsFileToGenerate

	pkg = p.Proto.GetOptions().GetGoPackage()

	structName := dataType[strings.LastIndex(dataType, ".")+1:]
	pkgName := pkg[strings.LastIndex(pkg, "/")+1:]

	if isFileToGenerate {
		refType = strings.Title(structName)
	} else {
		if pkgName == "" {
			refType = strings.Title(structName)
		} else {
			refType = pkgName + "." + strings.Title(structName)
		}
	}

	return
}

func appendGoImport(imports []string, dataType string) []string {
	if !strings.Contains(dataType, ".") {
		return imports
	}

	isFileToGenerate, pkg, refType := getGoPackageAndType(dataType)

	if !isFileToGenerate && !util.IsStrInSlice(`"`+pkg+`"`, imports) && pkg != "" {
		imports = append(imports, `"`+pkg+`"`)
	}

	importGoTypes[dataType] = refType

	return imports
}

func getGoImport(imports []string) (result string) {
	if len(imports) == 0 {
		return ""
	}

	sort.Slice(imports, func(i, j int) bool {
		return imports[i] > imports[j]
	})

	result = fmt.Sprintf(`import (
	%s
)
`, strings.Join(imports, "\n\t"))

	return
}

func (g *goService) Imports() (result string) {
	var imports []string

	for _, m := range g.Methods {
		imports = appendGoImport(imports, m.InputType)
		imports = appendGoImport(imports, m.OutputType)
		imports = appendGoImport(imports, m.ErrorType())
	}

	if g.HasCommonError() {
		imports = appendGoImport(imports, g.commonError())
	}

	return getGoImport(imports)
}

func (g *goService) commonError() string {
	return g.ServiceData.Options["common_error"]
}

// CommonError returns common error in go type
func (g *goService) CommonError() string {
	return wrapGoType(g.commonError())
}

func (g *goService) CommonErrorPointer() string {
	return "&" + wrapGoType(g.commonError())[1:]
}

func (g *goService) HasCommonError() bool {
	return g.ServiceData.CommonErrorType != ""
}

func (g *goService) hasCommonError(field string) bool {
	if !g.HasCommonError() {
		return false
	}

	for _, t := range g.Gen.DataTypes {
		if t.Name == g.ServiceData.CommonErrorType {
			for _, f := range t.Fields {
				if f.Name == field {
					return true
				}
			}
		}
	}
	return false
}

func (g *goService) HasCommonBindError() bool {
	return g.hasCommonError("bindError")
}

func (g *goService) HasCommonAuthError() bool {
	return g.hasCommonError("authError")
}

func (g *goService) HasCommonValidateError() bool {
	return g.hasCommonError("validateError")
}

func (g *goService) AuthRequired() bool {

	if authString, ok := g.Options["auth"]; ok {
		if authBool, err := strconv.ParseBool(authString); err == nil {
			return authBool
		}
	}
	return false
}

func (g *goService) ServicePath() string {
	return "/" + g.Name
}

func (g *goGen) genGoService(service *data.ServiceData) string {
	importGoTypes = make(map[string]string)

	buf := bytes.NewBufferString("")

	obj := newEchoService(service, g.PackageName)

	_goService = &goService{obj, g}
	err := g.serviceTpl.Execute(buf, _goService)
	if err != nil {
		util.Die(err)
	}

	return formatBuffer(buf)
}

func (g *goGen) Init(request *plugin.CodeGeneratorRequest) {
	g.echoGen.Init(request)

	g.structTpl = g.getTpl("/generator/template/go/struct.gogo")
	g.enumTpl = g.getTpl("/generator/template/go/enum.gogo")
}

func (g *goGen) Gen(applicationName string, packageName string, services []*data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options data.OptionMap) (result map[string]string, err error) {
	g.DataTypes = messages

	// Temporary hack from go server gen here
	// Should rewrite goGen completely later
	if services == nil || len(services) == 0 {
		g.serviceTpl = nil
		result, err = g.echoGen.Gen(applicationName, packageName, services, messages, enums, options)
		return
	}

	result, err = g.echoGen.Gen(applicationName, packageName, services, messages, enums, options)

	for _, service := range services {
		g.serviceTpl = g.getTpl("/generator/template/go/service.gogo")
		serviceContent := g.genGoService(service)
		serviceFilename := genEchoFileName(g.PackageName, service)
		g.serviceTpl = nil

		result[serviceFilename] = serviceContent
	}

	return
}

func init() {
	data.OutputMap["go"] = &goGen{}
}
