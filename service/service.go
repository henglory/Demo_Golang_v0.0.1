package service

import "github.com/henglory/Demo_Golang_v0.0.1/spec"

type Service struct {
	ARequestFn aRequestFn
	BRequestFn bRequestFn
}

func (s Service) RequestToA(req spec.AReq) (spec.ARes, error) {
	return s.ARequestFn(req)
}

func (s Service) RequestToB(req spec.BReq) (spec.BRes, error) {
	return s.BRequestFn(req)
}
