package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a Task to the list",

	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		addTask(task)
	},
}

func addTask(task string) {
	file, err := os.OpenFile("tasks.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		println("error opening file")
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	timestamp := time.Now().Format(time.RFC3339)
	err = writer.Write([]string{task, timestamp})
	if err != nil {
		print("error writing to file")
		return
	}

	fmt.Println("Task added")
}

func init() {
	rootCmd.AddCommand(addCmd)
}
