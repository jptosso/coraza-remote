/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes directory as a Coraza project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		p := cmd.Flag("path").Value.String()
		if p == "." {
			p, err = os.Getwd()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		tag := cmd.Flag("tag").Value.String()
		if tag == "" {
			tag = path.Base(p)
			fmt.Println("No --tag defined, using", tag)
		}
		if err := initDirectory(tag, p); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Initialized Coraza project in", p+"/.coraza")
	},
}

func initDirectory(id string, cwd string) error {
	cdir := path.Join(cwd, ".coraza")
	if _, err := os.Stat(cdir); os.IsNotExist(err) {
		if err := os.Mkdir(cdir, 0755); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Coraza project already exists in %s", cdir)
	}
	// now we create the ID file
	idFile := path.Join(cdir, "TAG")
	if err := os.WriteFile(idFile, []byte(id), 0644); err != nil {
		return err
	}
	tsFile := path.Join(cdir, "TIMESTAMP")
	if err := os.WriteFile(tsFile, []byte(strconv.FormatInt(time.Now().Unix(), 10)), 0644); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().StringP("tag", "t", "", "Tag of the remote WAF instance")
	initCmd.Flags().StringP("path", "p", ".", "Path of the WAF configuration")
}
