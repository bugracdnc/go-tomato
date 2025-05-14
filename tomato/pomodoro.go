package tomato

import (
	"fmt"
	"go-tomato/timer"
	"time"
)

// Line to be put between complete whole sessions
const line = "##############################"

// Pomodoro struct typedef, to be initialized by cobra flags in cli/root.go
type Pomodoro struct {
	StudyDuration, BreakDuration, LongBreakDuration, LongBreakIntervals int
	Title                                                               string
	DoLongBreaks                                                        bool
	start, end                                                          time.Time
}

// Start the timer initialized in cli/root.go by creating timers for study and breaks with the set values in an endless loop
func (p *Pomodoro) StartTimer() {
	//i to count the sessions
	i := 1
	for {
		// Print starting message
		fmt.Printf("Starting %s #%d\n\r", p.Title, i)
		// Start the study session timer
		timer.TimerFor(time.Duration(p.StudyDuration))
		// Indicate the end of study session
		fmt.Printf("\n\rStudy session %s #%d ended!\r\n\n", p.Title, i)

		// Deviate if long breaks are enabled (default)
		// Check if current session is a long break session (i.e., if it is time for a long break)
		if p.DoLongBreaks && i%p.LongBreakIntervals == 0 {

			// Print long break start
			fmt.Printf("Starting long break #%d\n\r", i/p.LongBreakIntervals)
			// Start the long break timer
			timer.TimerFor(time.Duration(p.LongBreakDuration))
			// Indicate long break end
			fmt.Printf("\n\rLong break session %s #%d ended!\r\n\n", p.Title, i/p.LongBreakIntervals)

		} else { // If it is time for a normal break or long breaks are disabled altogether

			// Print break start
			fmt.Printf("Starting break #%d\n\r", i)
			// Start the break timer
			timer.TimerFor(time.Duration(p.BreakDuration))
			// Indicate break end
			fmt.Printf("\n\rBreak session %s #%d ended!\r\n\n", p.Title, i)
		}

		// Increase the session counter
		i++
		fmt.Printf("%s\r\n", line)
	}

}
