package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendorf/dory/client"
)

func init() {
	rootCmd.AddCommand(setCommand)
	setCommand.Flags().StringP("name", "n", "", "Name of memory to set")
	setCommand.Flags().StringP("value", "v", "", "Value of memory")
	setCommand.Flags().DurationP("duration", "d", time.Nanosecond, "Lifetime of memory (e.g. 1h5m2s)")
	setCommand.MarkFlagRequired("name")
	setCommand.MarkFlagRequired("value")
	setCommand.MarkFlagRequired("duration")
}

var setCommand = &cobra.Command{
	Use:   "set",
	Short: "Set a memory",
	Long:  "Set a memory",
	Run: func(cmd *cobra.Command, args []string) {
		duration, err := cmd.Flags().GetDuration("duration")
		if err != nil {
			log.Fatal(err)
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}

		value, err := cmd.Flags().GetString("value")
		if err != nil {
			log.Fatal(err)
		}

		socketPath, ok := os.LookupEnv("DORY_SOCK")
		if !ok {
			log.Fatalln("DORY_SOCK must be specified")
		}

		doryClient, err := client.NewClient(socketPath)
		if err != nil {
			log.Fatalf("could not create client: %v\n", err)
		}

		if err := doryClient.CreateMemory(name, value, duration); err != nil {
			log.Fatalf("could not talk to server: %v\n", err)
		}

		fmt.Printf("Successfully set %s\n", name)
	},
}
