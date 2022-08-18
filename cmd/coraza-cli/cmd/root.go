/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/jptosso/coraza-center/client"
	"github.com/spf13/cobra"
)

var prefixPath string
var remote *client.Remote

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "coraza-cli",
	Short: "Coraza Centre CLI tool for managing Coraza WAF projects.\nAuthor Juan Pablo Tosso <jptosso@gmail.com>. https://github.com/jptosso/coraza-centre",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var err error
	prefixPath, err = os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	prefixPath = path.Join(prefixPath, ".coraza")
	if _, err := os.Stat(prefixPath); os.IsNotExist(err) {
		if err := os.Mkdir(prefixPath, 0755); err != nil {
			fmt.Println("Failed to create config directory:", err)
			os.Exit(1)
		}
	}
	bts, _ := os.ReadFile(path.Join(prefixPath, "server"))
	remote.Server = string(bts)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.client.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
