package task

import "github.com/spf13/cobra"

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Add a task to TODO list",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: implement function
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}