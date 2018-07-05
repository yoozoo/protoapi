package output

import (
	"protoapi/generator/data"
	"strings"
)

type springField struct {
	data.MessageField
}

func (s *springField) Title() string {
	return strings.Title(s.Name)
}

func (s *springField) JavaType() string {
	return javaTypes[s.MessageField.DataType]
}

func newSpringStruct(msg *data.MessageData, packageName string) *springStruct {
	o := &springStruct{
		msg,
		packageName,
		nil,
	}
	o.init()
	return o
}

type springStruct struct {
	*data.MessageData
	Package string
	Fields  []*springField
}

func (s *springStruct) init() {
	s.Fields = make([]*springField, len(s.MessageData.Fields))
	for i, f := range s.MessageData.Fields {
		s.Fields[i] = &springField{f}
	}
}

func (s *springStruct) ContructParam() string {
	params := make([]string, len(s.Fields))
	for i, f := range s.Fields {
		params[i] = f.JavaType() + " " + f.Name
	}
	return strings.Join(params, ", ")
}

func (s *springStruct) ClassName() string {
	return s.Name
}
