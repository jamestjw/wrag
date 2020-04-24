package wrag

import (
	"fmt"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func intToDuration(i int) time.Duration {
	return time.Duration(i) * time.Second
}

func bearerToken(t string) string {
	return fmt.Sprintf("bearer %s", t)
}
