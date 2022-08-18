/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jptosso/coraza-center/database"
	"github.com/spf13/cobra"
)

// wafNewCmd represents the new command
var wafNewCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tag := cmd.Flag("tag").Value.String()
		if tag == "" {
			fmt.Println("Tag is required")
			return
		}
		tx := database.DB.Create(&database.Waf{
			ID:  uuid.New().String(),
			Tag: tag,
		})
		if tx.Error != nil {
			fmt.Println("Failed to create WAF:", tx.Error.Error())
			return
		}
		if tx.RowsAffected == 0 {
			fmt.Println("Failed to create WAF: no rows affected")
			return
		}
		fmt.Println("WAF created")
	},
}

func init() {
	wafCmd.AddCommand(wafNewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wafNewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	wafNewCmd.Flags().StringP("tag", "t", "", "Tag to use for the new WAF")
}
