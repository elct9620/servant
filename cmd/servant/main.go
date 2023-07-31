package main

import (
	"fmt"

	"github.com/docker/cli/cli-plugins/manager"
	"github.com/docker/cli/cli-plugins/plugin"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
)

func main() {
	plugin.Run(func(dockerCli command.Cli) *cobra.Command {
		return &cobra.Command{
			Use: "servant",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Hello, servant!")
			},
		}
	}, manager.Metadata{
		SchemaVersion: "0.1.0",
		Vendor:        "Aotokitsuruya",
	})
}
