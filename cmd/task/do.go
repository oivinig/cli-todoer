package task

import "github.com/spf13/cobra"

var doCmd = &cobra.Command{
	Use: "do",
	Short: "Completes a task on TODO list",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: implement function
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}