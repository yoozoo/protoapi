package output

import (
	"strings"

	"github.com/yoozoo/protoapi/generator/data"
)

type echoEnumField struct {
	data.EnumField
}

func newEchoEnum(enum *data.EnumData, packageName string) *echoEnum {
	ss := strings.Split(packageName, "/")
	s := ss[len(ss)-1]
	o := &echoEnum{
		enum,
		s,
		nil,
	}
	o.init()
	return o
}

type echoEnum struct {
	*data.EnumData
	Package string
	Fields  []*echoEnumField
}

func (s *echoEnum) init() {
	s.Fields = make([]*echoEnumField, len(s.EnumData.Fields))
	for i, f := range s.EnumData.Fields {
		s.Fields[i] = &echoEnumField{f}
	}
}
