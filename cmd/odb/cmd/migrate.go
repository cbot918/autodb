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

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate short description",
	Long:  `migrate short description long discription`,
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

		if err := migrate(cfg, db); err != nil {
			fmt.Println("migrate failed: ", err)
			return
		}
		fmt.Println("migrate success")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migrate(cfg *internal.Config, db *sql.DB) error {

	err := internal.Migrate(cfg, db)
	if err != nil {
		return err
	}

	return nil
}
