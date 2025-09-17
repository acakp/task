package cmd

import (
	"fmt"
	"strconv"

	"acakp.task/db"
	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Deletes selected task(s)",
	Long: `task del <task number> [<task number...>]

Example:
  task del 1 4
  Deletes 1st and 4th tasks in task list
		`,
	Run: func(cmd *cobra.Command, args []string) {
		toDel := make([]int, len(args))
		for i, arg := range args {
			var err error
			toDel[i], err = strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("%v is not a number", toDel[i])
			}
		}
		db.Del(toDel)
		fmt.Printf("%v task(s) has been deleted", len(args))
	},
}

func init() {
	rootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
