package tomato

import (
	"fmt"
	"go-tomato/timer"
	"time"
)

const line = "##############################"

type Pomodoro struct {
	StudyDuration, BreakDuration, LongBreakDuration, LongBreakIntervals int
	Title                                                               string
	DoLongBreaks                                                        bool
	start, end                                                          time.Time
}

func (p *Pomodoro) StartTimer() {
	i := 1
	for true {
		fmt.Printf("Starting %s #%d\n\r", p.Title, i)
		timer.TimerFor(time.Duration(p.StudyDuration))
		fmt.Printf("\n\rStudy session %s #%d ended!\r\n\n", p.Title, i)

		if p.DoLongBreaks {
			if i%p.LongBreakIntervals == 0 {
				fmt.Printf("Starting long break #%d\n\r", i/p.LongBreakIntervals)
				timer.TimerFor(time.Duration(p.LongBreakDuration))
				fmt.Printf("\n\rLong break session %s #%d ended!\r\n\n", p.Title, i/p.LongBreakIntervals)
				i++
				continue
			}
		}
		fmt.Printf("Starting break #%d\n\r", i)
		timer.TimerFor(time.Duration(p.BreakDuration))
		fmt.Printf("\n\rBreak session %s #%d ended!\r\n\n", p.Title, i)
		i++
		fmt.Printf("%s\r\n", line)
	}

}
