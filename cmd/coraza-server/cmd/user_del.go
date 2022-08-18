/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/jptosso/coraza-center/database"
	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		id := cmd.Flag("id").Value.String()
		username := cmd.Flag("username").Value.String()
		if id == "" && username == "" {
			fmt.Println("ID or username is required")
			return
		}
		tx := database.DB.Delete(&database.User{}, "id = ? OR user_name = ?", id, username)
		if tx.Error != nil {
			fmt.Println("Failed to delete user:", tx.Error.Error())
			return
		}
		if tx.RowsAffected > 0 {
			fmt.Println("User deleted")
		} else {
			fmt.Println("User not found")
		}

	},
}

func init() {
	userCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	delCmd.Flags().String("id", "", "ID of the user to delete")
	delCmd.Flags().String("username", "", "name of the user to delete")
}
