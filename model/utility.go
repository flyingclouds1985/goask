package model

import (
	"time"
)

func UnixTime() time.Time {
	t := time.Now().Unix()
	return time.Unix(t, 0)
}
