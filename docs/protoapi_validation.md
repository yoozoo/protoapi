# protoapi Validation

protoapi validation handling

## Common Validation

When submitting a request, the request data must be validated before entering business logic because invalid data will cause error. This practice is common for most of the request, therefore we often have a common handler for validation such as many validation frameworks. And many comon validation error are predefined and shared among services. Take user registration as example:

* username field is required
* email field require email format

protoapi therefore defines validation error as a common error and handles it properly.

## Validation Error Definition

Validation error is defined in `protoapi_common.proto` which is common used proto file.

```protobuf
namespace ?

message CommonError {
    GenericError genericError = 1;
    AuthError authError = 2;
    ValidateError validateError = 3;
    BindError bindError = 4;
}

message ValidateError {
    repeated FieldError errors = 1;
}

message FieldError {
    string fieldName = 1;
    ValidateErrorType errorType = 2;
}

enum ValidateErrorType {
    INVALID_EMAIL = 0;
    FIELD_REQUIRED = 1;
    ...
}
```

`ValidationError` as a common error type defined in `CommonError` message. After running `protoapi` code generation, the `ValidationError` will be generated as a struct which contains an array of `FieldError`s In which you will find a field name and validation error type. `ValidationErrorType` is an predefined enum contains all the possible validation error type such as `Invalid email format` and `required field`, etc.

## Validation Error Code Generation

With the struct of `ValidationError`, a handler will also be generated. If we use protoapi to generate code for webframework [echo](https://github.com/labstack/echo) code, the validation part will be like:

```go
type CommonError struct {
    ValidateError *ValidateError `json:"validateError"`
}

type ValidateError struct {
    Errors []*FieldError `json:"errors"`
}

type ValidateErrorType int

const (
    INVALID_EMAIL ValidateErrorType = 0
    FIELD_REQUIRED ValidateErrorType = 1
)

func (code ValidateErrorType) String() string {
	names := map[ValidateErrorType]string{
        INVALID_EMAIL: "INVALID_EMAIL",
        FIELD_REQUIRED: "FIELD_REQUIRED",
	}

	return names[code]
}

func (code ValidateErrorType) Code() int {
	return (int)(code)
}
```

And for the request data which requires validation, we need to define the option of the specific field of the request object in the API protobuf file.

## Define Validation Field

To define field validation option, we put ``[(validation_option) = "selected_option"]`` after the field. For example:

```protobuf
// service list
import "protoapi_common.proto"

message ServiceSearchRequest{
    repeated int32 tag_ids = 1; // optional, for filter
    string prefix = 2 [(required) = 0, (format) = "email"];
    int32 env_id = 3;
    int32 offset = 4;
    int32 limit = 5;
}
```

It means the `prefix` field in `ServiceSearchRequest` is **required** to have and the `format must be email`.

Generated code by protoapi will be like:

```go
type ServiceSearchRequest struct {
    Tag_ids []int `json:"tag_ids"`
    Prefix string `json:"prefix"`
    Env_id int `json:"env_id"`
    Offset int `json:"offset"`
    Limit int `json:"limit"`
}

func (r ServiceSearchRequest) Validate() *ValidateError {
    errs := []*FieldError{}
	if r.Prefix == "" {
        e := FIELD_REQUIRED
		errs = append(errs, &FieldError{Field_name: r.Prefix, Error_type: &e})
	}
	if !rxEmail.MatchString(r.Prefix) {
        e := INVALID_EMAIL
		errs = append(errs, &FieldError{Field_name: r.Prefix, Error_type: &e})
	}
    if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}
```

The `Validate()` is the handler to validate all the field with validate option. You can run this method to check all the fields and it will return ValidateError if exists.

## handling validation error

In our generated API handler code, we will check the validation before sending request to the service. If validation failed, the validation error returned will be put in response and send to the client with HTTP code 420.

```go
    if valErr := in.Validate(); valErr != nil {
        resp := ResponseInternal{Error: CommonError{ValidateError: valErr}}
        return c.JSON(420, resp)
    }
```

So frontend need to check HTTP code. If it is 420, then it needs to check the existence of validation error and output it.
