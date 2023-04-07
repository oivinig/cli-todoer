package task

import (
	"fmt"
	"log"

	"github.com/oivinig/cli-todoer/internal/app/infrastructure"
	"github.com/oivinig/cli-todoer/internal/app/interfaces"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use: "do",
	Short: "Completes a task on TODO list",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := infrastructure.Connect()
		if err != nil {
			log.Fatal(err)
		}
		cmdr := interfaces.NewTaskCommander(db)
		err = cmdr.TaskInteractor.Do(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Item done.")
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}