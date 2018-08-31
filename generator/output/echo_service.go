package output

import (
	"strings"

	"version.uuzu.com/Merlion/protoapi/generator/data"
)

type echoMethod struct {
	*data.Method
	ServiceName string
	ServiceType string
	ErrorType   string
}

func (m *echoMethod) Title() string {
	return strings.Title(m.Name)
}

func (m *echoMethod) Path() string {
	return "/" + m.ServiceName + "." + m.Name
}

type echoService struct {
	*data.ServiceData
	Package string
	Methods []*echoMethod
}

func newEchoService(msg *data.ServiceData, packageName string) *echoService {
	ss := strings.Split(packageName, ".")
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
		s.Methods[i] = &echoMethod{&mtd, s.Name, "POST", ""}
		for _, option := range f.Options {
			if option.Name == MethodOptions["ServiceType"] {
				s.Methods[i].ServiceType = option.Value
			}
			if option.Name == MethodOptions["ErrorType"] {
				s.Methods[i].ErrorType = option.Value
			}
		}
	}
}
