/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// verifydbCmd represents the verifydb command
var verifydbCmd = &cobra.Command{
	Use:   "verifydb",
	Short: "verifydb short description",
	Long:  `verifydb short description`,
	Run: func(cmd *cobra.Command, args []string) {
		num, err := tableNumber()
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

func tableNumber() (int64, error) {
	q := fmt.Sprintf(`SELECT COUNT(table_name)
	FROM information_schema.tables
	WHERE table_schema = '%s'`, Cfg.DB_NAME)

	rows, err := DB.Query(q)
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
