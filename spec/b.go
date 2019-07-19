package spec

type BReq struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type BRes struct {
	StatusCode string `json:"statusCode"`
	Result     int    `json:"result"`
}
