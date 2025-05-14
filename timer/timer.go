package timer

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

// A string const filled with space to clear the line before printing the remaining time
const clear_line = "                                          "

// Start a timer for given duration
func TimerFor(duration time.Duration) int {
	// Calculate the end time by adding the duration to time.Now()
	end := time.Now().Add(time.Minute * duration)

	//Here we use go routines to get input from the user during the timer
	ch := make(chan byte)
	go func(ch chan byte) {
		// Disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		// Do not display entered characters on the screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

		// Catch interrupt signal (^C) to be able to set the terminal back to sane
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
		// Execute if interrupt signal is detected
		go func() {
			<-sigs
			// Set terminal back to sane
			exec.Command("stty", "-F", "/dev/tty", "sane").Run()
			// Manually print the signal: interrupt message
			fmt.Println("signal: interrupt")
			// Exit gracefully
			os.Exit(0)
		}()

		// byte array with a length of 1 to capture user input
		var b []byte = make([]byte, 1)
		for {
			os.Stdin.Read(b) // capture user input
			ch <- b[0]       // suppy the channel
		}
	}(ch)

	for {
		// If a key is pressed, detect it, else pass on
		select {
		case stdin := <-ch:
			fmt.Printf("        %q", stdin)
		default:
		}

		// If the time now is after the end, return gracefully (i.e., detect if the timer should stop)
		if time.Now().After(end) {
			return 0
		}
		// Else, calculate the remaining time and print formatted
		remaining := time.Time{}.Add(end.Sub(time.Now()))
		fmt.Printf("\r%s\r%s ", clear_line, remaining.Format("04:05")) // Clear the line with the empty line string and print back over with the time
		time.Sleep(1000)                                               // Wait for a second since the remaining time output will not change between seconds
	}
}
