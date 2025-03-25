package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "complete a task",
	Args:  cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		markTaskCompleted(task)
	},
}

func markTaskCompleted(task string) {
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
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	taskFound := false
	for _, record := range records {
		if record[0] == task && record[1] == "Pending" {
			record[1] = "Completed"
			taskFound = true
		}
		writer.Write(record)
	}

	if taskFound {
		fmt.Println("Task marked as completed")
	} else {
		fmt.Println("Task not found")
	}
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
