# protoapi Authentication

protoapi Authentication

## Defining Authentication

The authentication option definition is referenced by other protobuf files in `protoapi_common.proto`. `auth` is an option at the service level. The default is `false`. When set to `true`, all rpc under the service will support authentication. The protobuf file supports single file with multiple service. If necessary, please group the service according to the different authentication methods and whether it needs authentication. When the `common_error` option is set, the auth method will return `common_error` to the client when an error occurs.

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


## Verify Error code generation

If you use protoapi to generate the go [echo](https://github.com/labstack/echo) code， the AppAuthService would look like this:

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

In the implementation of the interface, user neeed to implement the AppAuthServiceAuth method, which is an authentication method. The parameter is echo.Context. You can store the necessary information in the context to facilitate subsequent API acquisition.
