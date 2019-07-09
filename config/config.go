package config

import (
	"fmt"
	"os"
)

var (
	ServicePort = fmt.Sprintf(":%s", os.Getenv("PORT"))
)
