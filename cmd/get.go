package cmd

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/wendorf/dory/client"
	"github.com/wendorf/dory/config"
	"golang.org/x/crypto/ssh/terminal"
)

func init() {
	rootCmd.AddCommand(getCommand)
	getCommand.Flags().StringP("name", "n", "", "Name of memory to get")
	getCommand.MarkFlagRequired("name")
}

var getCommand = &cobra.Command{
	Use:   "get",
	Short: "Get a memory",
	Long:  "Get a memory",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}

		doryClient, err := client.NewClient(config.SocketPath())
		if err != nil {
			log.Fatalf("could not create client: %v\n", err)
		}

		value, err := doryClient.GetMemory(name)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		os.Stdout.Write(value)
		if terminal.IsTerminal(int(syscall.Stdout)) {
			fmt.Print("\n")
		}
	},
}
