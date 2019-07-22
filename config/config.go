package config

import (
	"fmt"
	"os"
	"time"
)

const (
	ActionAUrl = "https://demo-service-stubby.herokuapp.com/api/actionA"
	ActionBUrl = "https://demo-service-stubby.herokuapp.com/api/actionB"
	ATimeout   = 10 * time.Second
	BTimeout   = 10 * time.Second
)

var (
	ServicePort = fmt.Sprintf(":%s", os.Getenv("PORT"))
)
