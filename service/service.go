package service

import "github.com/henglory/Demo_Golang_v0.0.1/spec"

type Service struct {
	ARequestFn aRequestFn
}

func (s Service) RequestToA(req spec.AReq) (spec.ARes, error) {
	return s.ARequestFn(req)
}
