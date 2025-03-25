package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "removes a task from the list",
	Args:  cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		removeTask(task)
	},
}

func removeTask(task string) {
	file, err := os.Open("tasks.csv")
	if err != nil {
		println("error opening file")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error reading tasks")
		return
	}

	file.Close()
	file, err = os.Create("tasks.csv")
	if err != nil {
		println("error opening file")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		if record[0] != task {
			writer.Write(record)
		}
	}

	fmt.Println("Task removed")
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
