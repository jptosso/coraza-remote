/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/jptosso/coraza-center/server"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		opts := server.ServerOptions{
			Bind:  cmd.Flag("bind").Value.String(),
			Log:   cmd.Flag("log").Value.String(),
			Debug: cmd.Flag("debug").Value.String() == "true",
		}
		if err := server.Start(opts); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serveCmd.Flags().StringP("bind", "b", "127.0.0.1:2022", "Address and port to bind to the server")
	serveCmd.Flags().StringP("log", "l", "/dev/stdout", "Path to the log file")
	serveCmd.Flags().BoolP("debug", "d", false, "Enable debug mode")
	serveCmd.Flags().StringP("admin-acl", "a", "0.0.0.0/0,::0/0", "Admin endpoint access control list CIDRs")
}
