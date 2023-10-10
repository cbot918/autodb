/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"

	"github.com/cbot918/autodb/internal"
	"github.com/spf13/cobra"
)

// verifydbCmd represents the verifydb command
var verifydbCmd = &cobra.Command{
	Use:   "verifydb",
	Short: "verifydb short description",
	Long:  `verifydb short description`,
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
		num, err := tableNumber(cfg, db)
		if err != nil {
			fmt.Println("query table count failed")
			return
		}
		fmt.Println("number of tables: ", num)
	},
}

func init() {
	rootCmd.AddCommand(verifydbCmd)
}

func tableNumber(cfg *internal.Config, db *sql.DB) (int64, error) {
	var q string
	if cfg.DB_DRIVER == "mysql" {
		q = fmt.Sprintf(`SELECT COUNT(table_name)
		FROM information_schema.tables
		WHERE table_schema = '%s'`, cfg.DB_NAME)
	} else if cfg.DB_DRIVER == "postgres" {
		q = fmt.Sprintf(`SELECT COUNT(*)
		FROM information_schema.tables
		WHERE table_schema = 'public' AND table_catalog = '%s';`, cfg.DB_NAME)
	}

	rows, err := db.Query(q)
	if err != nil {
		return 0, err
	}

	var count int64
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}
