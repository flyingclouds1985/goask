package model

import (
	"fmt"
	"strconv"
	"time"
)

func UnixTime() time.Time {
	t := time.Now().Unix()
	i, err := strconv.ParseInt(strconv.FormatInt(t, 10), 10, 64)
	if err != nil {
		fmt.Println("error in unix time parsing. ", err)
	}

	return time.Unix(i, 0)
}
