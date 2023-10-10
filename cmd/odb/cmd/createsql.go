/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/cbot918/autodb/internal"
	"github.com/spf13/cobra"
)

// createsqlCmd represents the createsql command
var createsqlCmd = &cobra.Command{
	Use:   "createsql",
	Short: "createsql short description",
	Long:  `createsql short description`,
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

		if isSQLFile(cfg.SQL_FILE) {

			defaultSQL := "https://raw.githubusercontent.com/cbot918/autodb/main/sample.sql"
			fmt.Println("download: ", defaultSQL)

			if err := downloadSQL(defaultSQL); err != nil {
				fmt.Println("downloadSQL failed")
				return
			}
			fmt.Println("createSQL success")
			return
		}

		fmt.Println(cfg.SQL_FILE + " ready, go next step!")

	},
}

func init() {
	rootCmd.AddCommand(createsqlCmd)

}

func isSQLFile(path string) bool {
	_, err := os.Stat(path)

	return err != nil
}

func downloadSQL(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("HTTP Status Code:", resp.StatusCode)
		return err
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	err = ioutil.WriteFile("sample.sql", body, 0644)
	if err != nil {
		fmt.Println("Error writing to sample.sql:", err)
		return err
	}

	return nil
}
