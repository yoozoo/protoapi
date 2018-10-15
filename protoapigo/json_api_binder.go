package protoapigo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// JSONAPIBinder is a Binder to echo design for JSON API
type JSONAPIBinder struct {
	*echo.DefaultBinder
}

// Bind use json decoder for all context type & DisallowUnknownFields
func (b *JSONAPIBinder) Bind(i interface{}, c echo.Context) (err error) {
	req := c.Request()
	d := json.NewDecoder(req.Body)
	d.DisallowUnknownFields()
	if err = d.Decode(i); err != nil {
		if ute, ok := err.(*json.UnmarshalTypeError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, offset=%v", ute.Type, ute.Value, ute.Offset))
		} else if se, ok := err.(*json.SyntaxError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error()))
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}
	return
}
