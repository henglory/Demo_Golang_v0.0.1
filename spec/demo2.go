package spec

type Demo2Req struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Demo2Res struct {
	StatusCode string `json:"statusCode"`
	StatusDesc string `json:"desc"`
	Result     int    `json:"result"`
}
