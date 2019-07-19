package service

import "github.com/henglory/Demo_Golang_v0.0.1/spec"

type aRequestFn func(req spec.AReq) (spec.ARes, error)
