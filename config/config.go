package config

import (
	"fmt"
	"os"
	"time"
)

const (
	ActionAUrl = "https://demo-service-stubby.herokuapp.com/api/actionA"
	ATimeout   = 10 * time.Second
)

var (
	ServicePort = fmt.Sprintf(":%s", os.Getenv("PORT"))
)
