package output

import (
	"strings"

	"github.com/yoozoo/protoapi/generator/data"
)

type echoMethod struct {
	*data.Method
	ServiceName string
}

func (m *echoMethod) Title() string {
	return strings.Title(m.Name)
}

func (m *echoMethod) Path() string {
	return "/" + m.ServiceName + "." + m.Name
}

func (m *echoMethod) ServiceType() string {
	if servType, ok := m.Options[data.MethodOptions[data.ServiceTypeMethodOption]]; ok {
		return servType
	}

	return "POST"
}

func (m *echoMethod) ErrorType() string {
	if errType, ok := m.Options[data.MethodOptions[data.ErrorTypeMethodOption]]; ok {
		return errType
	}

	return ""
}

type echoService struct {
	*data.ServiceData
	Package string
	Methods []*echoMethod
}

func newEchoService(msg *data.ServiceData, packageName string) *echoService {
	ss := strings.Split(packageName, "/")
	s := ss[len(ss)-1]
	o := &echoService{
		msg,
		s,
		nil,
	}
	o.init()
	return o
}

func (s *echoService) init() {
	s.Methods = make([]*echoMethod, len(s.ServiceData.Methods))
	for i, f := range s.ServiceData.Methods {
		mtd := f
		mtd.InputType = strings.Title(mtd.InputType)
		mtd.OutputType = strings.Title(mtd.OutputType)
		s.Methods[i] = &echoMethod{mtd, s.Name}
	}
}
