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
		rootCmd := &cobra.Command{
			Use:   "servant",
			Short: "Servant is a plugin for docker swarm to manage the service",
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
		}

		rootCmd.AddCommand(
			installCmd(),
			uninstallCmd(),
		)

		return rootCmd
	}, manager.Metadata{
		SchemaVersion: "0.1.0",
		Vendor:        "Aotokitsuruya",
	})
}

func installCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install the servantd on the swarm manager",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Work in progress")
			return nil
		},
	}

	return cmd
}

func uninstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstall the servantd on the swarm manager",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Work in progress")
			return nil
		},
	}

	return cmd
}
