package handler

import (
	"fmt"

	"github.com/henglory/Demo_Golang_v0.0.1/spec"
)

func Demo2(s interface {
	RequestToA(req spec.AReq) (spec.ARes, error)
	RequestToB(req spec.BReq) (spec.BRes, error)
}, req spec.Demo2Req) (res spec.Demo2Res) {
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

	bres, err := s.RequestToB(mappingToBReq(ares))
	if err != nil {
		res.StatusCode = "03"
		res.StatusDesc = fmt.Sprintf("ERROR FROM APP B: %s", err.Error())
		return
	}
	if bres.StatusCode != "00" {
		res.StatusCode = bres.StatusCode
		res.StatusDesc = "INVALID PARAMS FROM APP B"
		return
	}

	res.StatusCode = "00"
	res.StatusDesc = "OK"
	res.Result = bres.Result
	return
}

func mappingToBReq(a spec.ARes) (b spec.BReq) {
	b.X = a.Result
	return
}
