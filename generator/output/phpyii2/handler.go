package phpyii2

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/yoozoo/protoapi/generator/data"
)

// NewHandler return a pointer of new handler struct
func NewHandler(methods []*data.Method, baseNameSpace string) *Handler {
	nameSpace := baseNameSpace + "\\handlers"
	o := &Handler{methods, nameSpace, baseNameSpace}
	return o
}

// Handler is struct of php handler class
type Handler struct {
	Methods       []*data.Method
	NameSpace     string
	BaseNameSpace string
}

func (p *Handler) Gen(result map[string]string) error {
	filePath := strings.Replace(p.NameSpace, "\\", "/", -1)
	funcMap := template.FuncMap{
		"className": getPHPClassName,
	}

	buf := bytes.NewBufferString("")
	errorHandlerTplContent := data.LoadTpl("/generator/template/yii2/handlers/ErrorHandler.gophp")
	tpl, err := template.New("error handler").Funcs(funcMap).Parse(errorHandlerTplContent)
	if err != nil {
		return err
	}
	err = tpl.Execute(buf, p)
	if err != nil {
		return err
	}
	result[filePath+"/ErrorHandler.php"] = buf.String()

	buf = bytes.NewBufferString("")
	requestHandlerTplContent := data.LoadTpl("/generator/template/yii2/handlers/RequestHandler.gophp")
	tpl, err = template.New("request handler").Funcs(funcMap).Parse(requestHandlerTplContent)
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
