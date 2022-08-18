/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"path"

	"github.com/jptosso/coraza-center/database"
	"github.com/spf13/cobra"
)

var PrefixPath = ""

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "coraza-center",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	setup()
	if err := database.Connect(path.Join(PrefixPath, "database.sqlite")); err != nil {
		panic(err)
	}
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	home, _ := os.UserHomeDir()
	defaultDataPath := path.Join(home, ".coraza-center")
	rootCmd.PersistentFlags().StringVar(&PrefixPath, "prefix", defaultDataPath, "Data path for database and rule files (default is $HOME/.coraza-center/)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

func setup() {
	p := []string{
		"data/revisions", // Used to store revision files
		"data/waf",       // Used to store waf profiles
	}
	for _, dir := range p {
		dir = path.Join(PrefixPath, dir)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				panic(err)
			}
		}
	}
}
