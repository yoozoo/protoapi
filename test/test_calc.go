package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yoozoo/protoapi/test/result/go/calcsvr"
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

func (s *calcService) Minus(c echo.Context, req *calcsvr.AddReq) (resp *calcsvr.AddResp, bizError *calcsvr.AddError, err error) {
	resp = new(calcsvr.AddResp)
	if req.X > 100 {
		bizError = new(calcsvr.AddError)
		bizError.Req = req
		bizError.Error = "x overflow"
		return
	}
	resp.Result = req.X - req.Y

	return
}

func (s *calcService) CalcServiceAuth(c echo.Context) (err error) {
	/** get token from header **/
	header := c.Request().Header
	token := header.Get("Authorization")

	if token == "" {
		err = echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
		return
	}
	return
}

func main() {
	e := echo.New()

	srv := &calcService{}
	calcsvr.RegisterCalcService(e, srv)

	e.Logger.Fatal(e.Start(":8080"))
}
