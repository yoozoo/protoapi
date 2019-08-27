package phpyii2

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/yoozoo/protoapi/generator/data"
)

// NewError return a pointer of new Error struct
func NewEnum(enum *data.EnumData, baseNameSpace string) *Enum {
	nameSpace := baseNameSpace + "\\models"
	filePath := strings.Replace(nameSpace, "\\", "/", -1)
	filePath = filePath + "/" + strings.Title(enum.Name) + ".php"
	o := &Enum{enum, nameSpace, filePath}
	return o
}

// Error is struct of php error class
type Enum struct {
	*data.EnumData
	NameSpace string
	FilePath  string
}

func (p *Enum) Gen(result map[string]string) error {
	buf := bytes.NewBufferString("")

	tplContent := data.LoadTpl("/generator/template/yii2/models/enum.gophp")

	funcMap := template.FuncMap{
		"className": getPHPClassName,
	}

	tpl, err := template.New("Enum").Funcs(funcMap).Parse(tplContent)
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
