// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// KeyListRequest
type KeyListRequest struct {
	Service_id int `json:"service_id"`
	Env_id     int `json:"env_id"`
}

func (r KeyListRequest) Validate() *ValidateError {
	errs := []*FieldError{}
	if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}
