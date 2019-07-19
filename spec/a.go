package spec

type AReq struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type ARes struct {
	StatusCode string `json:"statusCode"`
	Result     int    `json:"result"`
}

type LogAReq struct {
	LogTime JSONTime `json:"logTime"`
	Info    string   `json:"info"`
	Req     string   `json:"requestBody"`
}

type LogARes struct {
	LogTime  JSONTime `json:"logTime"`
	Info     string   `json:"info"`
	Res      string   `json:"responseBody"`
	Overhead int64    `json:"overhead"`
}
