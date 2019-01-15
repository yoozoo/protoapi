### Example Proto
```protobuf
syntax = "proto3";

package example;
option go_package = "go_example";

import "protoapi_common.proto";

message AppRequest {
	string app_name = 1;
}

message BizError {
	string message = 1;
}

service AppService {
	option (common_error) = "CommonError";
	option (auth) = true;
	rpc getApp(AppRequest) returns (Empty) {
		option (error) = "BizError";
	}
}
```

### Generated Code
```go
// AppService is the interface contains all the controllers
type AppService interface {
	AppServiceAuth(c echo.Context) (err error)
	GetApp(c echo.Context, req *AppRequest) (resp *Empty, bizError *BizError, err error)
}

func _AppServiceAuth_Handler(srv AppService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			err = srv.AppServiceAuth(c)
			if err != nil {
				return c.String(500, err.Error())
			}

			return next(c)
		}
	}
}

func _getApp_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// bind data
		req := new(AppRequest)
		if err = c.Bind(req); err != nil {
			resp := CommonError{BindError: &BindError{err.Error()}}
			return c.JSON(420, resp)
		}

		resp, bizError, err := srv.GetApp(c, req)
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

// RegisterAppService is used to bind routers
func RegisterAppService(e *echo.Echo, srv AppService) {
	RegisterAppServiceWithPrefix(e, srv, "")
}

// RegisterAppServiceWithPrefix is used to bind routers with custom prefix
func RegisterAppServiceWithPrefix(e *echo.Echo, srv AppService, prefix string) {
	// switch to strict JSONAPIBinder, if using echo's DefaultBinder
	if _, ok := e.Binder.(*echo.DefaultBinder); ok {
		e.Binder = new(protoapigo.JSONAPIBinder)
	}
	g := e.Group(prefix+"/AppService", _AppServiceAuth_Handler(srv))
	g.POST(".getApp", _getApp_Handler(srv))
}
```
