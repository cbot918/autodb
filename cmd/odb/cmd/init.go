/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/cbot918/autodb/cmd/odb/cmd/pkg"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init short description",
	Long:  `init long description`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !pkg.IsFileExists(".env") {
			fmt.Println("init default .env")
			err = initProject()
			if err != nil {
				fmt.Println("init .env failed")
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

}

func initProject() error {

	// init .env
	fd, err := os.Create(".env")
	if err != nil {
		return err
	}
	_, err = fd.Write([]byte(pkg.Env()))
	if err != nil {
		return err
	}
	defer fd.Close()

	return nil
}
