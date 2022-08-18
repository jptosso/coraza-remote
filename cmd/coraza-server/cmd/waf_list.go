/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/jptosso/coraza-center/database"
	"github.com/jptosso/coraza-center/utils"
	"github.com/spf13/cobra"
)

// wafListCmd represents the list command
var wafListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		waf := []database.Waf{}
		database.DB.Find(&waf)
		table := &utils.Table{
			Headers: []string{"ID", "Tag"},
			Rows:    make([][]string, len(waf)),
		}
		for i, user := range waf {
			table.Rows[i] = []string{
				user.ID,
				user.Tag,
			}
		}
		fmt.Println(table.String())
	},
}

func init() {
	wafCmd.AddCommand(wafListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wafListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wafListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
