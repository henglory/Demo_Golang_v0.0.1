package handler

import (
	"fmt"

	"github.com/henglory/Demo_Golang_v0.0.1/spec"
)

func Demo1(s interface {
	RequestToA(req spec.AReq) (spec.ARes, error)
}, req spec.Demo1Req) (res spec.Demo1Res) {
	if req.X >= 500 {
		res.StatusCode = "01"
		res.StatusDesc = "X CANNOT MORE THAN 500"
		return
	}
	if req.Z < 0 {
		res.StatusCode = "02"
		res.StatusDesc = "Z CANNOT BE NEGATIVE"
		return
	}
	ares, err := s.RequestToA(spec.AReq{
		X: req.X,
		Y: req.Y,
	})
	if err != nil {
		res.StatusCode = "03"
		res.StatusDesc = fmt.Sprintf("ERROR FROM APP A: %s", err.Error())
		return
	}
	if ares.StatusCode != "00" {
		res.StatusCode = ares.StatusCode
		res.StatusDesc = "INVALID PARAMS FROM APP A"
		return
	}

	res.StatusCode = "00"
	res.StatusDesc = "OK"
	res.Result = ares.Result
	return
}
