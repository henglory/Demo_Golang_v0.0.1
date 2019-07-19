package spec

type CReq struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type CRes struct {
	StatusCode string `json:"statusCode"`
	Result     int    `json:"result"`
}
