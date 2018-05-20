package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().BoolP("version", "", false, "Display version")
}

var rootCmd = &cobra.Command{
	Use:   "dory",
	Short: "Dory is a very forgetful fish",
	Long:  `A simple store for data you want to save for a short bit, like a password during your work day.`,
	Run: func(cmd *cobra.Command, args []string) {
		version, err := cmd.Flags().GetBool("version")
		if err != nil {
			log.Fatal(err)
		}

		if version {
			fmt.Println("0.0.1")
			os.Exit(0)
		}

		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
