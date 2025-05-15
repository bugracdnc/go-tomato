package tomato

import (
	"fmt"
	"go-tomato/timer"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func disableInputBuf() {
	// Disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// Do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func enableInputBuf() {
	// Set terminal back to sane
	exec.Command("stty", "-F", "/dev/tty", "sane").Run()
}

func startInputListener(ch chan byte) {
	//Function to disable input buffering and disable echo
	disableInputBuf()

	// Catch interrupt signal (^C) to be able to set the terminal back to sane
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	// Execute if interrupt signal is detected
	go func() {
		sig := <-sigs
		// Set terminal back to sane
		enableInputBuf()

		// Exit normally if SIGQUIT
		if sig == syscall.SIGQUIT {
			fmt.Printf("\r\nExiting...\r\n")
			os.Exit(0)
		} else { // exit with the signal
			fmt.Printf("signal: %s\r\n", sig.String())
			os.Exit(128 + int(sig.(syscall.Signal)))
		}
	}()
	go func() {
		// byte array with a length of 1 to capture user input
		var b []byte = make([]byte, 1)
		for {
			os.Stdin.Read(b) // capture user input
			ch <- b[0]       // suppy the channel
		}
	}()

}

// Prints specific user inputs to interract with the program
func printInputUsage() {
	fmt.Printf("(s)kip, (+) add minute, (-) remove minute, (q)uit\r\n")
}

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
	// Declare a chan to capture key presses in a goroutine
	ch := make(chan byte)

	// Start the input listener
	startInputListener(ch)

	//i to count the sessions
	i := 1
	for {
		// Print starting message
		fmt.Printf("Starting %s #%d\n\r", p.Title, i)
		printInputUsage()
		// Start the study session timer
		timer.TimerFor(time.Duration(p.StudyDuration), ch)
		// Indicate the end of study session
		fmt.Printf("\n\rStudy session %s #%d ended!\r\n\n", p.Title, i)

		// Deviate if long breaks are enabled (default)
		// Check if current session is a long break session (i.e., if it is time for a long break)
		if p.DoLongBreaks && i%p.LongBreakIntervals == 0 {

			// Print long break start
			fmt.Printf("Starting long break #%d\n\r", i/p.LongBreakIntervals)
			printInputUsage()
			// Start the long break timer
			timer.TimerFor(time.Duration(p.LongBreakDuration), ch)
			// Indicate long break end
			fmt.Printf("\n\rLong break session %s #%d ended!\r\n\n", p.Title, i/p.LongBreakIntervals)

		} else { // If it is time for a normal break or long breaks are disabled altogether

			// Print break start
			fmt.Printf("Starting break #%d\n\r", i)
			printInputUsage()
			// Start the break timer
			timer.TimerFor(time.Duration(p.BreakDuration), ch)
			// Indicate break end
			fmt.Printf("\n\rBreak session %s #%d ended!\r\n\n", p.Title, i)
		}

		// Increase the session counter
		i++
		fmt.Printf("%s\r\n", line)
	}

}
