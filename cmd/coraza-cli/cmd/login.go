/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Save login credentials for Coraza centre",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var server string
		if remote == nil {
			fmt.Print("Coraza Center: ")
		} else {
			fmt.Printf("Coraza Center (%s): ", remote.Server)
		}
		fmt.Scan(&server)
		if server == "" {
			server = remote.Server
		}
		// Prompt for username
		fmt.Print("Username: ")
		var username string
		fmt.Scanln(&username)
		// Prompt for password
		fmt.Print("Password: ")
		password, err := term.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := login(server, username, string(password)); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Login credentials saved")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
