package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "odb",
	Short: "autodb a convenient db tool",
	Long:  `autodb long description`,
}

func Execute() error {
	var err error

	err = rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
