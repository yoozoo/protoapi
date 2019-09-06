package phpyii2

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/yoozoo/protoapi/generator/data"
)

// NewMessage return a pointer of new Message struct
func NewMessage(msg *data.MessageData, baseNameSpace string, enums []*data.EnumData) *Message {
	nameSpace := baseNameSpace + "\\models"
	filePath := strings.Replace(nameSpace, "\\", "/", -1)
	filePath = filePath + "/" + strings.Title(msg.Name) + ".php"

	o := &Message{msg, nameSpace, filePath, enums}
	return o
}

// Message is struct of php message class
type Message struct {
	*data.MessageData
	NameSpace string
	FilePath  string
	Enums     []*data.EnumData
}

func (p *Message) IsObject(fieldType string) bool {
	switch fieldType {
	case data.StringFieldType,
		data.DoubleFieldType,
		data.IntFieldType,
		data.BooleanFieldType,
		data.Int32FieldType,
		data.Int64FieldType:
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

func (p *Message) IsEnum(fieldType string) bool {
	for _, enum := range p.Enums {
		if enum.Name == fieldType {
			return true
		}
	}
	return false
}

func (p *Message) Gen(result map[string]string) error {
	buf := bytes.NewBufferString("")

	tplContent := data.LoadTpl("/generator/template/yii2/models/message.gophp")

	funcMap := template.FuncMap{
		"isObject":  p.IsObject,
		"isEnum":    p.IsEnum,
		"className": getPHPClassName,
	}

	tpl, err := template.New("message").Funcs(funcMap).Parse(tplContent)
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
