package task

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "task",
		Short: "A TODO manager using command line.",
		Long: `Task is a CLI task manager that helps you keep track of your TODO lists.`,
	}
)
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}