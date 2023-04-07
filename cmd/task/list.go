package task

import (
	"fmt"
	"log"

	"github.com/oivinig/cli-todoer/internal/app/infrastructure"
	"github.com/oivinig/cli-todoer/internal/app/interfaces"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List all the tasks from TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := infrastructure.Connect()
		if err != nil {
			log.Fatal(err)
		}
		cmdr := interfaces.NewTaskCommander(db)
		values, err := cmdr.TaskInteractor.List()
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range values {
			fmt.Println(v)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}