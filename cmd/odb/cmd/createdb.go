/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/cbot918/autodb/internal"
	"github.com/spf13/cobra"
)

// createdbCmd represents the createdb command
var createdbCmd = &cobra.Command{
	Use:   "createdb",
	Short: "createdb short description",
	Long:  `createdb long description`,
	Run: func(cmd *cobra.Command, args []string) {
		defer DB.Close()
		if err := createdb(); err != nil {
			fmt.Println("createdb failed")
			return
		}
		fmt.Println("createdb success")
	},
}

func init() {
	rootCmd.AddCommand(createdbCmd)

}

func createdb() error {
	// createdb
	err := internal.CreateDB(Cfg)
	if err != nil {
		return err
	}
	return nil
}
