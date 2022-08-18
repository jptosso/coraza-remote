/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a WAF tag configuration to the current directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		tag := cmd.Flag("tag").Value.String()
		if tag == "" {
			fmt.Println("TAG is required")
			os.Exit(1)
		}
		fmt.Println("Downloading remote package")
		cw, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := initDirectory(tag, cw); err != nil {
			fmt.Println(err)
			return
		}
		if err := remote.Download(tag, cw); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	downloadCmd.Flags().String("tag", "", "TAG of the WAF config to download")
}
