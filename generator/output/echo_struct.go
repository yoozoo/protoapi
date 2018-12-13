package output

import (
	"strings"

	"github.com/yoozoo/protoapi/generator/data"
)

type echoField struct {
	data.MessageField
	isEnum bool
}

// this is ugly, should rely on proto_structs later
var importGoTypes map[string]string

func (s *echoField) Title() string {
	return strings.Title(s.Name)
}

func (s *echoField) Type() string {
	// if not primary type return data type and ignore the . in the data type
	dataType := s.DataType
	if val, ok := importGoTypes[dataType]; ok {
		dataType = val
	}

	if _, ok := wrapperTypes[dataType]; !ok && !s.isEnum {
		dataType = "*" + dataType
	}

	// check if the field is repeated
	if s.Label == data.FieldRepeatedLabel {
		dataType = "[]" + dataType
	}

	return dataType
}

func (s *echoField) ValidateRequired() bool {
	if _, ok := s.Options[data.FieldOptions[data.RequiredFieldOption]]; ok {
		return true
	}
	return false
}

func (s *echoField) ValidateFormat() string {
	if format, ok := s.Options[data.FieldOptions[data.FormatFieldOption]]; ok {
		return format
	}
	return ""
}

func newEchoStruct(msg *data.MessageData, packageName string, enums []*data.EnumData) *echoStruct {
	ss := strings.Split(packageName, "/")
	s := ss[len(ss)-1]
	o := &echoStruct{
		msg,
		s,
		nil,
	}
	o.init(enums)
	return o
}

type echoStruct struct {
	*data.MessageData
	Package string
	Fields  []*echoField
}

func (s *echoStruct) init(enums []*data.EnumData) {
	importGoTypes = make(map[string]string)
	s.Fields = make([]*echoField, len(s.MessageData.Fields))
	for i, f := range s.MessageData.Fields {
		isEnum := false
		for _, enum := range enums {
			if enum.Name == f.DataType {
				isEnum = true
				break
			}
		}
		s.Fields[i] = &echoField{f, isEnum}
	}
}

func (s *echoStruct) Imports() (result string) {
	var imports []string

	for _, f := range s.MessageData.Fields {
		imports = appendGoImport(imports, f.DataType)
	}

	return getGoImport(imports)
}

func (s *echoStruct) ClassName() string {
	pos := strings.LastIndex(s.Name, ".")
	if pos > 0 {
		return strings.Title(s.Name[pos+1:])
	}

	return strings.Title(s.Name)
}

func (s *echoStruct) IsCommonErrorStruct() bool {
	if _goService == nil {
		return false
	}

	return _goService.CommonError() == s.ClassName()
}

func (s *echoStruct) ValidateRequired() bool {
	for _, f := range s.Fields {
		if f.ValidateRequired() {
			return true
		}
	}

	return false
}
