# protoapi 验证

protoapi 验证处理

## 常用验证

对大多数情况来说，当用户提交申请， 申请数据必须经过验证才能进入业务逻辑中。因此我们需要提供一个常用验证处理为各种框架和服务所使用。并且多数验证错误是被提前定义并共享给不同的服务的。举用户注册为例：

* 用户名是必填项
* 邮箱必须是以邮件的格式

protoapi 因此把验证错误当作共同错误处理。

## 验证错误的定义

验证错误定义在 `protoapi_common.proto`中被其他protobuf文件引用。

```protobuf

extend google.protobuf.FieldOptions {
    string format = 51002;
    bool required = 51003;
}

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

`ValidationError` 作为一个常用错误类被定义在 `CommonError` message。 当使用 `protoapi`来生成代码后， `ValidationError` 会被生成为一个struct。 该struct包含了一个 `FieldError` 数组。 `FieldError`包括了对应需要填空的名字和错误类型。 所有验证错误类型被预定义在 `ValidationErrorType` 这个enum中。这些错误类型包括 `Invalid email format` 和 `required field` 等等。

## 验证错误代码生成

对应 `ValidationError` 的handler也会同时生成。 如果使用protoapi来生成echo[echo](https://github.com/labstack/echo)代码， 验证的部分会是这样：

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

请在proto文件中定义需要验证的选项。

## 定义验证选项

请把 ``[(validation_option) = "selected_option"]`` 置于需要验证的值域后。 比如：

```protobuf
import "protoapi_common.proto"

message ServiceSearchRequest{
    repeated int32 tag_ids = 1; // optional, for filter
    string prefix = 2 [(val_required) = true, (val_format) = "email"];
    int32 env_id = 3;
    int32 offset = 4;
    int32 limit = 5;
}
```

以上proto文件表明 `ServiceSearchRequest` 中的 `prefix` 值域是必须的，并且 `必须是电子邮件格式` 。

protoapi生成的代码：

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

`Validate()` 是生成的自带的验证方法，它会检验所有被定义需要验证的值域并且返回对应的错误。

## 验证错误处理

在生成的API代码中，我们会在发送请求前进行验证。如果验证失败，错误信息会伴随HTTP Code 420被返回。

```go
    if valErr := in.Validate(); valErr != nil {
        resp := CommonError{ValidateError: valErr}
        return c.JSON(420, resp)
    }
```

对于前端，如果返回的HTTP code是420， 则需要去显示返回的验证错误。
