package actions

import (
	"fmt"
	"task/db"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the tasks in the task list",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete! Why not take a vacation!")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(ListCmd)
}
