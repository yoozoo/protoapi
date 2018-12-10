package phpyii2

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/yoozoo/protoapi/util"

	"github.com/yoozoo/protoapi/generator/data"
)

// NewController return a pointer of new controller struct
func NewController(nameSpace string, methods []*data.Method) *Controller {

	fileDir := strings.Replace(nameSpace, "\\", "/", -1)
	filePath := fileDir + "/controllers/ApiController.php"

	o := &Controller{nameSpace, filePath, methods}
	return o
}

// Controller is struct of php Controller class
type Controller struct {
	NameSpace string
	FilePath  string
	Methods   []*data.Method
}

func (p *Controller) escape(s string) string {
	return strings.Replace(s, `\`, `\\`, -1)
}

func (p *Controller) Gen(result map[string]string) error {
	buf := bytes.NewBufferString("")

	tplContent := data.LoadTpl("/generator/template/yii2/controllers/ApiController.gophp")
	funcMap := template.FuncMap{
		"escape":    p.escape,
		"className": util.GetPHPClassName,
		"title":     strings.Title,
	}
	tpl, err := template.New("Controller").Funcs(funcMap).Parse(tplContent)
	if err != nil {
		return err
	}
	err = tpl.Execute(buf, p)
	if err != nil {
		return err
	}
	result[p.FilePath] = buf.String()

	return nil
}
