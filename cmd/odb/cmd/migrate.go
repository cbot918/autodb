/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/cbot918/autodb/internal"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate short description",
	Long:  `migrate short description long discription`,
	Run: func(cmd *cobra.Command, args []string) {

		if err := migrate(); err != nil {
			fmt.Println("migrate failed")
			return
		}
		fmt.Println("migrate success")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func isSQLFile(file string) bool {
	_, err := os.Stat(file)

	return err == nil
}

func migrate() error {

	err := internal.Migrate(Cfg, DB)
	if err != nil {
		return err
	}

	fmt.Println("3秒後開始 read metadata")

	time.Sleep(3 * time.Second)

	return nil
}
