/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/cbot918/autodb/cmd/odb/cmd/pkg"
	"github.com/spf13/cobra"
)

// sqlcCmd represents the sqlc command
var sqlcCmd = &cobra.Command{
	Use:   "sqlc",
	Short: "sqlc short",
	Long:  `sqlc long`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := Sqlc(); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(sqlcCmd)

}

func Sqlc() error {
	return pkg.SqlcInit()
}
