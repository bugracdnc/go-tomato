package cli

import (
	"go-tomato/tomato"
	"log"

	"github.com/spf13/cobra"
)

// Defaults for variables as const because of being accessed in multiple locations
const long_break_duration_default = 30
const long_break_interval_default = 5

// Create a Pomodoro object in the tomato package
var pomodoro = tomato.Pomodoro{}

// cobra root command, used for argument parsing (flags)
var rootCmd = &cobra.Command{
	Use:   "tomato",
	Short: "Keep track of intervals using pomodoro method",
	Run: func(cmd *cobra.Command, args []string) {
		// Set DoLongBreaks value to true if either long break duration or interval is changed
		if pomodoro.LongBreakDuration != long_break_duration_default || pomodoro.LongBreakIntervals != long_break_interval_default {
			pomodoro.DoLongBreaks = true
		}

		//Start the pomodoro timer with the loop of sessions: study -> long break | break
		pomodoro.StartTimer()
	},
}

// Access point from outside the package
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// init the flags and their defaults
func init() {
	rootCmd.Flags().IntVarP(&pomodoro.StudyDuration, "duration", "d", 25, "Set study duration in minutes")
	rootCmd.Flags().IntVarP(&pomodoro.BreakDuration, "break", "b", 5, "Set break duration in minutes")
	rootCmd.Flags().IntVarP(&pomodoro.LongBreakIntervals, "intervals", "i", long_break_interval_default, "Set intervals for long break")
	rootCmd.Flags().IntVar(&pomodoro.LongBreakDuration, "lb-duration", long_break_duration_default, "Set long break duration")
	rootCmd.Flags().BoolVar(&pomodoro.DoLongBreaks, "lb-disable", true, "Disable long breaks")
	rootCmd.Flags().StringVarP(&pomodoro.Title, "title", "t", "tomato", "Set a title")
}
