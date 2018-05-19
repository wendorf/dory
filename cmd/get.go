package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/wendorf/dory/client"
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
		socketPath, ok := os.LookupEnv("DORY_SOCK")
		if !ok {
			log.Fatalln("DORY_SOCK must be specified")
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}

		doryClient, err := client.NewClient(socketPath)
		if err != nil {
			log.Fatalf("could not create client: %v\n", err)
		}

		value, err := doryClient.GetMemory(name)
		if err != nil {
			log.Fatalf("could not get memory %s: %v\n", name, err)
		}

		fmt.Println(value)
	},
}
