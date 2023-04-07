package task

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List all the tasks from TODO list",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: implement function
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}