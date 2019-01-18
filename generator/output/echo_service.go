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

func (m *echoMethod) MethodPath() string {
	return "." + m.Name
}

func (m *echoMethod) Path() string {
	return "/" + m.ServiceName + "." + m.Name
}

func (m *echoMethod) ServiceType() string {
	if servType, ok := m.Options[data.MethodOptions[data.ServiceTypeMethodOption].Name]; ok {
		return servType
	}

	return "POST"
}

func (m *echoMethod) ErrorType() string {
	if errType, ok := m.Options[data.MethodOptions[data.ErrorTypeMethodOption].Name]; ok {
		return errType
	}

	return ""
}

func wrapGoType(dataType string) string {
	if val, ok := importGoTypes[dataType]; ok {
		dataType = val
	}

	if _, ok := wrapperTypes[dataType]; !ok {
		if strings.Contains(dataType, ".") {
			dataType = "*" + dataType
		} else {
			dataType = "*" + strings.Title(dataType)
		}
	}

	return dataType
}

func (m *echoMethod) ErrorGoType() string {
	return wrapGoType(m.ErrorType())
}

func (m *echoMethod) InputGoType() string {
	return wrapGoType(m.InputType)
}

func (m *echoMethod) InputGoTypeName() string {
	stmt := wrapGoType(m.InputType)
	if strings.HasPrefix(stmt, "*") {
		return stmt[1:]
	}

	return stmt
}

func (m *echoMethod) OutputGoType() string {
	return wrapGoType(m.OutputType)
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
		s.Methods[i] = &echoMethod{mtd, s.Name}
	}
}
