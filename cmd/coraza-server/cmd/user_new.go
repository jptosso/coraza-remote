/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"

	"github.com/google/uuid"
	"github.com/jptosso/coraza-center/database"
	"github.com/spf13/cobra"
)

// userNewCmd represents the new command
var userNewCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		id := uuid.New().String()
		username := cmd.Flag("username").Value.String()
		if username == "" {
			fmt.Println("Username is required")
			return
		}
		password := cmd.Flag("password").Value.String()
		if password == "" {
			// default password is a random sha512
			password = randomPassword()
			fmt.Println("No password specified, using random password:", password)
		}
		admin := cmd.Flag("admin").Value.String()
		tx := database.DB.Create(&database.User{
			ID:       id,
			UserName: username,
			Password: password,
			Admin:    admin == "true",
		})
		if tx.Error != nil {
			fmt.Println("Failed to create user:", tx.Error.Error())
			return
		}
		fmt.Println("User created with ID", id)
	},
}

func init() {
	userCmd.AddCommand(userNewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userNewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	userNewCmd.Flags().StringP("username", "u", "", "User name to create. Only alphanumeric characters and underscores are allowed")
	userNewCmd.Flags().StringP("password", "p", "", "Password to create. If not provided, a random password will be generated")
	userNewCmd.Flags().Bool("admin", false, "Create the user as an admin")
}

func randomPassword() string {
	randomBytes := make([]byte, 60)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)
}
