package main

import (
	"github.com/labstack/echo"
	"version.uuzu.com/Merlion/protoapi/test/result/go/calcsvr"
)

type calcService struct{}

// Echo just return req
func (s *calcService) Add(c echo.Context, req *calcsvr.AddReq) (resp *calcsvr.AddResp, bizError *calcsvr.AddError, err error) {
	resp = new(calcsvr.AddResp)
	if req.X > 100 {
		bizError = new(calcsvr.AddError)
		bizError.Req = req
		bizError.Error = "x overflow"
		return
	}
	resp.Result = req.X + req.Y

	return
}

func main() {
	e := echo.New()

	srv := &calcService{}
	calcsvr.RegisterCalcService(e, srv)

	e.Logger.Fatal(e.Start(":8080"))
}
