/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/cbot918/autodb/cmd/odb/cmd/pkg"
	"github.com/cbot918/autodb/internal"
	"github.com/spf13/cobra"
)

var (
	Cfg *internal.Config
	DB  *sql.DB
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "odb",
	Short: "autodb a convenient db tool",
	Long:  `autodb long description`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	var err error
	if !pkg.IsFileExists(".env") {
		fmt.Println("init default .env")
		err = initProject()
		if err != nil {
			return err
		}
	}

	// init config
	Cfg, err = internal.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// initial db pool
	DB, err = internal.NewDB(Cfg)
	if err != nil {
		log.Fatal(err)
	}
	err = rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.odb.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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