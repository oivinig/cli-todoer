package task

import "github.com/spf13/cobra"

var deleteCmd = &cobra.Command{
	Use: "delete",
	Aliases: []string{"del"},
	Short: "Deletes a task from TODO list",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: implement function
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}