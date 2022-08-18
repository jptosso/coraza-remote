/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the current directory to Coraza Centre",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		server := cmd.Flag("server").Value.String()
		if server != "" {
			remote.Server = server
		}
		var cw string
		var err error
		if len(args) == 0 {
			cw, err = os.Getwd()
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			cw = args[0]
		}
		fmt.Println("Uploading path:", cw)
		loadProjectDir(cw)
		if localConfig == nil {
			fmt.Println("Project not initialized")
			return
		}
		if err := validateCredentials(); err != nil {
			fmt.Println(err.Error())
			return
		}
		if err := remote.Upload(cw, localConfig.WafTag); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Deployment completed")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	deployCmd.Flags().StringP("server", "s", "", "Coraza Centre server endpoint")
}
