package phpyii2

import (
	"bytes"
	"errors"
	"strings"
	"text/template"

	"github.com/yoozoo/protoapi/util"

	"github.com/yoozoo/protoapi/generator/data"
)

// NewModule return a pointer of new module struct
func NewModule(NameSpace string, ServiceName string, methods []data.Method) *Module {
	fileDir := strings.Replace(NameSpace, "\\", "/", -1)
	temp := strings.SplitN(fileDir, "/", 3)
	if len(temp) != 3 {
		util.Die(errors.New("Invalid namespace : " + NameSpace))
	}
	name := temp[2]

	o := &Module{NameSpace, name, ServiceName, methods}
	return o
}

// Module is struct of php Module class
type Module struct {
	NameSpace   string
	Name        string
	ServiceName string
	Methods     []data.Method
}

func (p *Module) routeFromUrl(name string) string {
	return p.ServiceName + "." + name
}

func (p *Module) routeToUrl(name string) string {
	return p.Name + "/api/" + name
}

func (p *Module) Gen(result map[string]string) error {
	filePath := strings.Replace(p.NameSpace, "\\", "/", -1)

	buf := bytes.NewBufferString("")
	tplContent := data.LoadTpl("/generator/template/yii2/Module.gophp")
	funcMap := template.FuncMap{
		"routeFromUrl": p.routeFromUrl,
		"routeToUrl":   p.routeToUrl,
	}
	tpl, err := template.New("Module").Funcs(funcMap).Parse(tplContent)
	if err != nil {
		return err
	}
	err = tpl.Execute(buf, p)
	if err != nil {
		return err
	}
	result[filePath+"/Module.php"] = buf.String()

	buf = bytes.NewBufferString("")
	tplContent = data.LoadTpl("/generator/template/yii2/RequestHandler.gophp")
	funcMap = template.FuncMap{
		"className": util.GetPHPClassName,
	}
	tpl, err = template.New("Module").Funcs(funcMap).Parse(tplContent)
	if err != nil {
		return err
	}
	err = tpl.Execute(buf, p)
	if err != nil {
		return err
	}
	result[filePath+"/RequestHandler.php"] = buf.String()

	return nil
}
