package cmd

import (
	"fmt"
	"strconv"

	"acakp.task/db"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks task(s) as complete",
	Run: func(cmd *cobra.Command, args []string) {
		toDo := make([]int, len(args))
		for i, arg := range args {
			var err error
			toDo[i], err = strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("%v is not a number", toDo[i])
			}
		}
		db.Do(toDo)
		fmt.Printf("%v task(s) has been marked as complete", len(args))
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
