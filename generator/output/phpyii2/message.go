package phpyii2

import (
	"strings"

	"github.com/yoozoo/protoapi/generator/data"
)

// NewMessage return a pointer of new Message struct
func NewMessage(msg *data.MessageData, baseNameSpace string) *Message {
	nameSpace := baseNameSpace + "\\models"
	filePath := strings.Replace(nameSpace, "\\", "/", -1)
	filePath = filePath + "/" + strings.Title(msg.Name) + ".php"
	o := &Message{msg, nameSpace, filePath}
	return o
}

// Message is struct of php message class
type Message struct {
	*data.MessageData
	NameSpace string
	FilePath  string
}

func (p *Message) Title(s string) string {
	return strings.Title(s)
}
