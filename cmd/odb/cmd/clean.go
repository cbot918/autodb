/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"

	"github.com/cbot918/autodb/cmd/odb/cmd/pkg"
	"github.com/cbot918/autodb/internal"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "short clean description",
	Long:  `long clean description`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, db, cfgErr, dbErr := pkg.Init()
		if cfgErr != nil {
			fmt.Println("load config error")
			return
		}
		if dbErr != nil {
			fmt.Println("db open error")
			return
		}
		defer db.Close()

		if err := clean(cfg, db); err != nil {
			fmt.Println("clean failed")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

func clean(cfg *internal.Config, db *sql.DB) error {
	return internal.Clean(cfg, db)
}
