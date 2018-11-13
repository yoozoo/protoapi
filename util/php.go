package util

import (
	"strings"
)

//GetPHPClassName rename class to valid php class name
func GetPHPClassName(old string) string {
	if strings.ToUpper(old) == "EMPTY" {
		return "Blank"
	}
	return strings.Title(old)
}
