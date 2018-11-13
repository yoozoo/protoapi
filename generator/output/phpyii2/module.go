package phpyii2

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/yoozoo/protoapi/util"

	"github.com/yoozoo/protoapi/generator/data"
)

// NewModule return a pointer of new module struct
func NewModule(NameSpace string, service *data.ServiceData) *Module {
	o := &Module{NameSpace, service}
	return o
}

// Module is struct of php Module class
type Module struct {
	NameSpace string
	Service   *data.ServiceData
}

func (p *Module) Gen(result map[string]string) error {
	filePath := strings.Replace(p.NameSpace, "\\", "/", -1)

	buf := bytes.NewBufferString("")
	tplContent := data.LoadTpl("/generator/template/yii2/Module.gophp")
	tpl, err := template.New("Module").Parse(tplContent)
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
	funcMap := template.FuncMap{
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
