package phpyii2

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/yoozoo/protoapi/generator/data"
)

// NewError return a pointer of new Error struct
func NewError(msg *data.MessageData, baseNameSpace string, enums []*data.EnumData) *Error {
	nameSpace := baseNameSpace + "\\models"
	filePath := strings.Replace(nameSpace, "\\", "/", -1)
	filePath = filePath + "/" + strings.Title(msg.Name) + ".php"
	o := &Error{msg, nameSpace, filePath, enums}
	return o
}

// Error is struct of php error class
type Error struct {
	*data.MessageData
	NameSpace string
	FilePath  string
	Enums     []*data.EnumData
}

func (p *Error) IsObject(fieldType string) bool {
	switch fieldType {
	case data.StringFieldType,
		data.DoubleFieldType,
		data.IntFieldType,
		data.BooleanFieldType:
		return false
	default:
		// check if is enum
		for _, enum := range p.Enums {
			if enum.Name == fieldType {
				return false
			}
		}
		return true
	}
}

func (p *Error) Gen(result map[string]string) error {
	buf := bytes.NewBufferString("")

	tplContent := data.LoadTpl("/generator/template/yii2/models/error.gophp")

	funcMap := template.FuncMap{
		"isObject":  p.IsObject,
		"className": getPHPClassName,
	}

	tpl, err := template.New("Error").Funcs(funcMap).Parse(tplContent)
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
