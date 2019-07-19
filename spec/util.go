package spec

import (
	"fmt"
	"time"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RFC3339Nano))
	return []byte(stamp), nil
}
