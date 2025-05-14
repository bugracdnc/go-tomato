package cli

import (
	"go-tomato/tomato"
	"log"

	"github.com/spf13/cobra"
)

const long_break_duration_default = 30
const long_break_interval_default = 5

var pomodoro = tomato.Pomodoro{}

var rootCmd = &cobra.Command{
	Use:   "tomato",
	Short: "Keep track of intervals using pomodoro method",
	Run: func(cmd *cobra.Command, args []string) {
		if pomodoro.LongBreakDuration != long_break_duration_default || pomodoro.LongBreakIntervals != long_break_interval_default {
			pomodoro.DoLongBreaks = true
		}
		pomodoro.StartTimer()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&pomodoro.StudyDuration, "study", "s", 25, "Set study duration in minutes")
	rootCmd.Flags().IntVarP(&pomodoro.BreakDuration, "break", "b", 5, "Set break duration in minutes")
	rootCmd.Flags().IntVarP(&pomodoro.LongBreakIntervals, "long-break-intervals", "i", long_break_interval_default, "Set intervals for long break")
	rootCmd.Flags().IntVarP(&pomodoro.LongBreakDuration, "long-break-duration", "l", long_break_duration_default, "Set long break duration")
	rootCmd.Flags().BoolVarP(&pomodoro.DoLongBreaks, "enable-long-breaks", "e", false, "Enable long breaks")
	rootCmd.Flags().StringVarP(&pomodoro.Title, "title", "t", "tomato", "Set a title")
}
