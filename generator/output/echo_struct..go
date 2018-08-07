package output

import (
	"protoapi/generator/data"
	"strings"
)

type echoField struct {
	data.MessageField
}

func (s *echoField) Title() string {
	return strings.Title(s.Name)
}

func (s *echoField) Type() string {
	return toGoType(s.MessageField.DataType, s.MessageField.Label)
}

func newEchoStruct(msg *data.MessageData, packageName string) *echoStruct {
	ss := strings.Split(packageName, ".")
	s := ss[len(ss)-1]
	o := &echoStruct{
		msg,
		s,
		nil,
	}
	o.init()
	return o
}

type echoStruct struct {
	*data.MessageData
	Package string
	Fields  []*echoField
}

func (s *echoStruct) init() {
	s.Fields = make([]*echoField, len(s.MessageData.Fields))
	for i, f := range s.MessageData.Fields {
		s.Fields[i] = &echoField{f}
	}
}

func (s *echoStruct) ClassName() string {
	return s.Name
}
