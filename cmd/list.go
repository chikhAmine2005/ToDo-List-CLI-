package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
	"time"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "shows the tasks added",

	Run: func(cmd *cobra.Command, args []string) {
		listTasks()
	},
}

func listTasks() {
	file, err := os.Open("tasks.csv")
	if err != nil {
		println("error opening file")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		println("error reading tasks")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintln(writer, "Tasks\tStatus\tAdded")

	for _, record := range records {

		task, status, timestamp := record[0], record[1], record[2]

		parsedTime, err := time.Parse(time.RFC3339, timestamp)
		if err != nil {
			fmt.Println("Error parsing time for:", task, "Skipping...")
			continue
		}

		timeAgo := timediff.TimeDiff(parsedTime)
		fmt.Fprintf(writer, "%s\t%s\t%s\n", task, status, timeAgo)
	}

	writer.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)
}
