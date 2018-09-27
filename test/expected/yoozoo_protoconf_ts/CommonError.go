// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// CommonError
type CommonError struct {
	GenericError  *GenericError  `json:"genericError"`
	AuthError     *AuthError     `json:"authError"`
	ValidateError *ValidateError `json:"validateError"`
	BindError     *BindError     `json:"bindError"`
}

func (r CommonError) Validate() *ValidateError {
	errs := []*FieldError{}
	if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}
