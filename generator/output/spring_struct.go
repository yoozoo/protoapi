package output

import (
	"protoapi/generator/data"
	"strings"
)

type springStruct struct {
	*data.MessageData
}

func (s *springStruct) ContructParam() string {
	return strings.Join(nil, ", ")
}
