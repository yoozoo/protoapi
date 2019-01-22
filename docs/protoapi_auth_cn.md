# protoapi 认证

protoapi 认证

## 认证的定义

认证选项定义在 `protoapi_common.proto`中被其他protobuf文件引用。`auth`是一个在service层面的选项，默认是`false`，当设置为`true`时，该service下的所有rpc都将支持认证。protobuf文件支持单文件多service，必要时请按照有无认证和认证方式的不同来划分service. 在设置了`common_error`选项的情况下，auth方法在错误时会返回`common_error`给客户端。

```protobuf
syntax = "proto3";

package example;
option go_package = "go_example";

import "protoapi_common.proto";

message AppRequest {
	string app_name = 1;
}

message AuthedAppRequest {
	string app_name = 1;
}

message BizError {
	string message = 1;
}

service AppAuthService {
	option (common_error) = "CommonError";
	option (auth) = true;
	rpc getAuthedApp(AuthedAppRequest) returns (Empty) {
		option (error) = "BizError";
	}
}

service AppNoAuthService {
	option (common_error) = "CommonError";
	rpc getApp(AppRequest) returns (Empty) {
		option (error) = "BizError";
	}
}
```


## 验证错误代码生成

如果使用protoapi来生成go [echo](https://github.com/labstack/echo)代码， AppAuthService的部分会是这样：

```go
type AppAuthService interface {
	AppAuthServiceAuth(c echo.Context) (err error)
	getAuthedApp(c echo.Context, req *AppRequest) (resp *Empty, bizError *BizError, err error)
}

func _AppAuthServiceAuth_Handler(srv AppAuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			err = srv.AppAuthServiceAuth(c)
			if err != nil {
				return c.String(500, err.Error())
			}

			return next(c)
		}
	}
}

func _getAuthedApp_Handler(srv AppAuthService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// bind data
		req := new(AppRequest)
		if err = c.Bind(req); err != nil {
			resp := CommonError{BindError: &BindError{err.Error()}}
			return c.JSON(420, resp)
		}

		resp, bizError, err := srv.getAuthedApp(c, req)
		if err != nil {
			if e, ok := err.(*CommonError); ok {
				return c.JSON(420, e)
			}
			return c.String(500, err.Error())
		}
		if bizError != nil {
			return c.JSON(400, bizError)
		}
		return c.JSON(200, resp)
	}
}

func RegisterAppAuthService(e *echo.Echo, srv AppService) {
	RegisterAppAuthServiceWithPrefix(e, srv, "")
}

func RegisterAppAuthServiceWithPrefix(e *echo.Echo, srv AppAuthService, prefix string) {
	// switch to strict JSONAPIBinder, if using echo's DefaultBinder
	if _, ok := e.Binder.(*echo.DefaultBinder); ok {
		e.Binder = new(protoapigo.JSONAPIBinder)
	}
	g := e.Group(prefix+"/AppService", _AppAuthServiceAuth_Handler(srv))
	g.POST(".getAuthedApp", _getAuthedAppAuthService(srv))
}
```

在实现interface的时候需要额外实现AppAuthServiceAuth方法，这个方法是认证方法，参数为echo.Context, 可以把必要的信息存储在context中方便后续api获取。
