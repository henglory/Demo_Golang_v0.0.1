package aclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/henglory/Demo_Golang_v0.0.1/spec"
)

type AClient struct {
	url       string
	loggingFn func(info interface{}) error
	client    interface {
		Do(req *http.Request) (*http.Response, error)
	}
}

func New(url string, timeout time.Duration, loggingFn func(info interface{}) error) *AClient {
	client := &http.Client{
		Timeout: timeout,
	}
	c := &AClient{
		url:       url,
		loggingFn: loggingFn,
		client:    client,
	}
	return c
}

func (c *AClient) Request(req spec.AReq) (spec.ARes, error) {
	var res spec.ARes

	b, err := json.Marshal(req)
	if err != nil {
		return res, err
	}
	body := bytes.NewReader(b)
	hreq, err := http.NewRequest("POST", c.url, body)
	if err != nil {
		return res, err
	}
	defer hreq.Body.Close()
	hreq.Header.Set("Content-Type", "text/json;charset=UTF-8")
	c.loggingFn(spec.LogAReq{
		LogTime: spec.JSONTime(time.Now()),
		Info:    "Start",
		Req:     string(b),
	})
	st := time.Now()

	hresp, err := c.client.Do(hreq)
	if err != nil {
		return res, err
	}
	resp, err := ioutil.ReadAll(hresp.Body)
	c.loggingFn(spec.LogARes{
		LogTime:  spec.JSONTime(time.Now()),
		Info:     "End",
		Res:      string(resp),
		Overhead: time.Since(st).Nanoseconds() / 1000,
	})
	if err != nil {
		return res, err
	}
	defer hresp.Body.Close()
	err = json.Unmarshal(resp, &res)
	return res, err
}
