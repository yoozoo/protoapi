package output

import (
	"github.com/yoozoo/protoapi/generator/data"
)

type springMethod struct {
	*data.Method
	ServiceName string
}

func (m *springMethod) Path() string {
	return "/" + m.ServiceName + "." + m.Name
}

func (m *springMethod) ServiceType() string {
	if servType, ok := m.Options[data.MethodOptions[data.ServiceTypeMethodOption]]; ok {
		return servType
	}

	return "POST"
}

type springService struct {
	*data.ServiceData
	Package string
	Methods []*springMethod
}

func newSpringService(msg *data.ServiceData, packageName string) *springService {
	o := &springService{
		msg,
		packageName,
		nil,
	}
	o.init()
	return o
}

func (s *springService) init() {
	s.Methods = make([]*springMethod, len(s.ServiceData.Methods))
	for i, f := range s.ServiceData.Methods {
		mtd := f
		s.Methods[i] = &springMethod{mtd, s.Name}
	}
}
