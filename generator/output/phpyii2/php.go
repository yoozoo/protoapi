package phpyii2

import (
	"strings"

	"github.com/yoozoo/protoapi/generator/data"
)

//GetPHPClassName rename class to valid php class name
func getPHPClassName(old string) string {
	if strings.ToUpper(old) == "EMPTY" {
		return "Blank"
	}
	switch old {
	case data.IntFieldType:
		return "int"
	case data.BooleanFieldType:
		return "boolean"
	case data.StringFieldType:
		return "string"
	case data.DoubleFieldType:
		return "double"
	default:
		return strings.Title(old)
	}
}
