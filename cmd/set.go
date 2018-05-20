package cmd

import (
	"bufio"
	"fmt"
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

		var value string
		if cmd.Flags().Changed("value") {
			value, err = cmd.Flags().GetString("value")

			if err != nil {
				log.Fatal(err)
			}
		} else if !terminal.IsTerminal(int(syscall.Stdin)) {
			value = ""
			s := bufio.NewScanner(os.Stdin)
			for s.Scan() {
				value+=s.Text()
			}
		} else {
			fmt.Print("Enter value: ")
			byteValue, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print("\n")
			value = string(byteValue)
		}

		doryClient, err := client.NewClient(config.SocketPath())
		if err != nil {
			log.Fatalf("could not create client: %v\n", err)
		}

		if err := doryClient.CreateMemory(name, value, duration); err != nil {
			log.Fatalf("could not talk to server: %v\n", err)
		}

		fmt.Printf("Successfully set %s\n", name)
	},
}
