package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/wendorf/dory/server"
)

func init() {
	rootCmd.AddCommand(serverCommand)
}

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Launch dory server",
	Long:  `Launch dory server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Launched")
		socketPath, ok := os.LookupEnv("DORY_SOCK")
		if !ok {
			log.Fatalln("DORY_SOCK must be specified")
		}

		fmt.Println("listening...")
		doryServer := server.NewDoryServer(socketPath)
		if err := doryServer.Run(); err != nil {
			log.Fatalf("server failed: %v\n", err)
		}
	},
}
