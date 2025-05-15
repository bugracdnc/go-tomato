package timer

import (
	"fmt"
	"syscall"
	"time"
)

// Enum type for return values of the timer
type TimerReturn int

// Timer can either finish by itself or can be skipped, represent these two states
const (
	TimerDone TimerReturn = iota
	TimerSkipped
)

// Start a timer for given duration
func TimerFor(duration time.Duration, ch chan byte) TimerReturn {
	// Calculate the end time by adding the duration to time.Now()
	end := time.Now().Add(time.Minute * duration)

	for {
		// If a key is pressed, detect it, else pass on
		select {
		case stdin := <-ch:
			switch stdin {
			case 'q':
				// Exit by sending an interrupt signal to be detected by the input listener in pomodoro and exit from there
				if err := syscall.Kill(syscall.Getpid(), syscall.SIGQUIT); err != nil {
					panic(err)
				}
			case 's':
				// Skip to the next timer by ending the current one
				fmt.Printf("\r\nSkipping session...\r\n")
				return TimerSkipped
			case '+':
				// Add a minute to the timer
				end = end.Add(time.Minute)
			case '-':
				// Remove a minute from the timer
				if end.Minute() > 1 {
					end = end.Add(-time.Minute)
				}
			}
		default:
		}

		// If the time now is after the end, return gracefully (i.e., detect if the timer should stop)
		if time.Now().After(end) {
			return TimerDone
		}
		// Else, calculate the remaining time and print formatted
		remaining := time.Time{}.Add(end.Sub(time.Now()))
		fmt.Printf("\r     \r%s ", remaining.Format("04:05")) // Clear the line with the empty line string and print back over with the time
		time.Sleep(1000)                                      // Wait for a second since the remaining time output will not change between seconds
	}
}
