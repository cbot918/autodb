/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"

	"github.com/cbot918/autodb/internal"
	"github.com/cbot918/autodb/internal/gentype"
	"github.com/spf13/cobra"
)

var gentypeCmd = &cobra.Command{
	Use:   "gentype",
	Short: "gentype short",
	Long:  `gentype long`,
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

		if err := _gentype(cfg, db); err != nil {
			fmt.Println("gentype failed")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(gentypeCmd)

}

func _gentype(cfg *internal.Config, db *sql.DB) error {

	// initial DBMetadata: get result dbm
	dbm, err := internal.NewDBMetadata(cfg, db)
	if err != nil {
		return err
	}

	internal.PrintJSON(dbm)

	gt := gentype.NewGentype(cfg, dbm)

	err = gt.InitContent()
	if err != nil {
		return err
	}

	err = gt.Create()
	if err != nil {
		return err
	}

	return nil
}
