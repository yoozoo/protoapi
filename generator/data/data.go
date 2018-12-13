// Package data data structure , constant values and global variables used in protoconf generator related code
package data

import (
	"os"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/yoozoo/protoapi/generator/data/tpl"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

const (
	//IntFieldType datatype string for interge, it is assumed to be of signed and at least 64 bit
	IntFieldType = "int"
	//Int32FieldType datatype string for interge 32 bit
	Int32FieldType = "int32"
	//Int64FieldType datatype string for interge 64 bit
	Int64FieldType = "int64"
	// BooleanFieldType datatype string for boolean
	BooleanFieldType = "bool"
	// StringFieldType datatype string for string
	StringFieldType = "string"
	// DoubleFieldType datatype string for floating point, it is assumed to be at least of ieee double precision
	DoubleFieldType = "double"

	// PathSeparator the path seperator used to form the full key (ie, key/sub_key )
	PathSeparator = "/"
	// FieldRepeatedLabel is the label for repeated data type
	FieldRepeatedLabel = "LABEL_REPEATED"
	// FieldOptionalLabel is the label for optional data type
	FieldOptionalLabel = "LABEL_OPTIONAL"
	// FieldRequiredLabel is the label for required data type
	FieldRequiredLabel = "LABEL_REQUIRED"
	// JavaPackageOption is Java package option constant
	JavaPackageOption = "javaPackageOption"
	// ServiceCommonErrorOption is service common_error option
	ServiceCommonErrorOption = 51008
	// ServiceTypeMethodOption is service method option
	ServiceTypeMethodOption = 51006
	// ErrorTypeMethodOption is error return type option
	ErrorTypeMethodOption = 51007
	// FormatFieldOption is the field type validation field option
	FormatFieldOption = 51002
	// RequiredFieldOption is the required type validation field option
	RequiredFieldOption = 51003

	// ComErrMsgName  is common error message name
	ComErrMsgName = "CommonError"
)

// ServiceOptions is the map of field number and field name in service options
var ServiceOptions = map[int32]string{
	ServiceCommonErrorOption: "common_error",
}

// MethodOptions is the map of field number and field name in method options
var MethodOptions = map[int32]string{
	ServiceTypeMethodOption: "service_method",
	ErrorTypeMethodOption:   "error",
}

// FieldOptions is the map of field number and field name in field options
var FieldOptions = map[int32]string{
	FormatFieldOption:   "val_format",
	RequiredFieldOption: "val_required",
}

var debugTpl = os.Getenv("debugTpl") == "true"

// LoadTpl is the function to load template file as string
// It loads file content from esc embed by default
// Set environment variable debugTpl to "true" to load template from disk directly
func LoadTpl(tplPath string) string {
	//useLocal is true, the filesystem's contents are instead used.
	return tpl.FSMustString(debugTpl, tplPath)
}

// EnumField a enum entry for enum datatype
type EnumField struct {
	Name  string // enum entry name
	Value int32  // enum entry value
}

// EnumData a structure to represent a enum datatype
type EnumData struct {
	Name   string      // enum type name
	Fields []EnumField // enum entries
}

// MessageField a field for the defined message.
type MessageField struct {
	Name     string // message variable name
	DataType string // message variable type
	Key      string // coresponding key name for the variable, default is the same as variable name
	Label    string
	Options  OptionMap
}

// MessageData a structure to represent a message datatype
type MessageData struct {
	File   string         // file where this message is defined
	Name   string         // name of the message (class, struct)
	Fields []MessageField // message members
}

type Method struct {
	Name       string
	InputType  string
	OutputType string
	HttpMtd    string
	URI        string
	Options    OptionMap // service method option (default is GET and POST)
}

type ServiceData struct {
	Name            string
	Methods         []*Method
	Options         OptionMap
	CommonErrorType string
	Service         *descriptor.ServiceDescriptorProto
}

// Option is a structure represents the option declared in a proto file
type OptionMap map[string]string

type CodeGenerator interface {
	Init(request *plugin.CodeGeneratorRequest)
	Gen(applicationName string, packageName string, services *ServiceData, messages []*MessageData, enums []*EnumData, options OptionMap) (map[string]string, error)
}

// OutputMap the registra for output code type and its associated output plugin
var OutputMap = make(map[string]CodeGenerator)
