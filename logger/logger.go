package logger

import (
	"encoding/json"
	"fmt"
)

func WriteLog(msg interface{}) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
