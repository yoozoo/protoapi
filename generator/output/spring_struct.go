package output

import (
	"protoapi/generator/data"
	"strings"
)

type springStruct struct {
	*data.MessageData
	Package string
}

func (s *springStruct) ContructParam() string {
	return strings.Join(nil, ", ")
}

func (s *springStruct) ClassName() string {
	return s.Name
}
