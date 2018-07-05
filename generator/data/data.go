// Package data data structure , constant values and global variables used in protoconf generator related code
package data

import (
	"os"
	"protoapi/generator/data/tpl"
)

const (
	//IntFieldType datatype string for interge, it is assumed to be of signed and at least 64 bit
	IntFieldType = "int"
	// BooleanFieldType datatype string for boolean
	BooleanFieldType = "bool"
	// StringFieldType datatype string for string
	StringFieldType = "string"
	// DoubleFieldType datatype string for floating point, it is assumed to be at least of ieee double precision
	DoubleFieldType = "double"

	// PathSeparator the path seperator used to form the full key (ie, key/sub_key )
	PathSeparator = "/"
)

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
}

type ServiceData struct {
	Name    string
	Methods []Method
}

// OutputFunc the code output plugin prototype
type OutputFunc func(applicationName string, packageName string, services *ServiceData, messages []*MessageData, enums []*EnumData) (map[string]string, error)

// OutputMap the registra for output code type and its associated output plugin
var OutputMap = make(map[string]OutputFunc)
