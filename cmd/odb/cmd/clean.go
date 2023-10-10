/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/cbot918/autodb/internal"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "short clean description",
	Long:  `long clean description`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := clean(); err != nil {
			fmt.Println("clean failed")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

func clean() error {
	return internal.Clean(Cfg, DB)
}
