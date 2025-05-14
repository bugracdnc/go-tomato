package timer

import (
	"fmt"
	"time"
)

const clear_line = "              "

func TimerFor(duration time.Duration) int {
	end := time.Now().Add(time.Minute * duration)

	for true {
		if time.Now().After(end) {
			return 0
		}
		remaining := time.Time{}.Add(end.Sub(time.Now()))
		fmt.Printf("\r%s\r%s ", clear_line, remaining.Format("04:05"))
		time.Sleep(1000)
	}

	return 0
}
