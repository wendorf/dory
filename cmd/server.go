package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/wendorf/dory/config"
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

		doryServer := server.NewDoryServer(config.SocketPath())
		if err := doryServer.Run(); err != nil {
			log.Fatalf("server failed: %v\n", err)
		}
	},
}
