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

		cfg, db, cfgErr, dbErr := internal.Init()
		if cfgErr != nil {
			fmt.Println("load config error")
			return
		}
		if dbErr != nil {
			fmt.Println("db open error")
			return
		}

		defer db.Close()
		if err := createdb(cfg); err != nil {
			fmt.Println("createdb failed: ", err)
			return
		}
		fmt.Println("createdb success")
	},
}

func init() {
	rootCmd.AddCommand(createdbCmd)

}

func createdb(cfg *internal.Config) error {
	// createdb
	err := internal.CreateDB(cfg)
	if err != nil {
		return err
	}
	return nil
}
