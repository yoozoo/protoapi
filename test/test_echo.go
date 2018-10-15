package main

import (
	"github.com/labstack/echo"
	"github.com/yoozoo/protoapi/protoapigo"
	"github.com/yoozoo/protoapi/test/result/go/echosvr"
)

type echoService struct{}

// Echo just return req
func (s *echoService) Echo(c echo.Context, req *echosvr.Msg) (resp *echosvr.Msg, err error) {
	resp = req

	return
}

func main() {
	e := echo.New()
	e.Binder = new(protoapigo.JSONAPIBinder)

	srv := &echoService{}
	echosvr.RegisterEchoService(e, srv)

	e.Logger.Fatal(e.Start(":8080"))
}
