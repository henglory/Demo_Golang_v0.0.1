package bclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/henglory/Demo_Golang_v0.0.1/spec"
)

type BClient struct {
	url       string
	loggingFn func(info interface{}) error
	client    interface {
		Do(req *http.Request) (*http.Response, error)
	}
}

func New(url string, timeout time.Duration, loggingFn func(info interface{}) error) *BClient {
	client := &http.Client{
		Timeout: timeout,
	}
	c := &BClient{
		url:       url,
		loggingFn: loggingFn,
		client:    client,
	}
	return c
}

func (c *BClient) Request(req spec.BReq) (spec.BRes, error) {
	var res spec.BRes

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
	hreq.Header.Set("Content-Type", "application/json;charset=UTF-8")
	c.loggingFn(spec.LogAReq{
		LogTime: spec.JSONTime(time.Now()),
		Info:    "StartCallB",
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
		Info:     "EndCallB",
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
