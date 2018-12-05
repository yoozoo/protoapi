package output

import (
	"strings"

	"github.com/yoozoo/protoapi/generator/data"
)

type echoField struct {
	data.MessageField
	isEnum bool
}

func (s *echoField) Title() string {
	return strings.Title(s.Name)
}

func (s *echoField) Type() string {
	// if not primary type return data type and ignore the . in the data type
	dataType := s.DataType
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

func (s *echoStruct) ClassName() string {
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
