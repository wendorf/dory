package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendorf/dory/client"
	"github.com/wendorf/dory/config"
	"golang.org/x/crypto/ssh/terminal"
)

func init() {
	rootCmd.AddCommand(setCommand)
	setCommand.Flags().StringP("name", "n", "", "Name of memory to set")
	setCommand.Flags().StringP("value", "v", "", "Value of memory")
	setCommand.Flags().DurationP("duration", "d", time.Nanosecond, "Lifetime of memory (e.g. 1h5m2s)")
	setCommand.MarkFlagRequired("name")
	setCommand.MarkFlagRequired("duration")
}

var setCommand = &cobra.Command{
	Use:   "set",
	Short: "Set a memory",
	Long:  "Set a memory. Supports piping value through stdin, or interactively asks if no value is provided.",
	Run: func(cmd *cobra.Command, args []string) {
		duration, err := cmd.Flags().GetDuration("duration")
		if err != nil {
			log.Fatal(err)
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}

		var value []byte
		if cmd.Flags().Changed("value") {
			valueParam, err := cmd.Flags().GetString("value")

			if err != nil {
				log.Fatal(err)
			}

			value = []byte(valueParam)
		} else if !terminal.IsTerminal(int(syscall.Stdin)) {
			value, err = ioutil.ReadAll(os.Stdin)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Print("Enter value: ")
			value, err = terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print("\n")
		}

		doryClient, err := client.NewClient(config.SocketPath())
		if err != nil {
			log.Fatalf("could not create client: %v\n", err)
		}

		if err := doryClient.CreateMemory(name, value, duration); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Printf("Successfully set %s\n", name)
	},
}
