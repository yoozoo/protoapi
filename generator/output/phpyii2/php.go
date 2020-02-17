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
	case data.IntFieldType, data.Int32FieldType, data.Int64FieldType:
		return "int"
	case data.BooleanFieldType:
		return "bool"
	case data.StringFieldType:
		return "string"
	case data.DoubleFieldType:
		return "double"
	default:
		return strings.Title(old)
	}
}
