package spec

type BReq struct {
	X int `json:"x"`
}

type BRes struct {
	StatusCode string `json:"statusCode"`
	Result     int    `json:"result"`
}
