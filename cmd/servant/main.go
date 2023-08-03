package main

import (
	"github.com/docker/cli/cli-plugins/manager"
	"github.com/docker/cli/cli-plugins/plugin"
	"github.com/docker/cli/cli/command"
	"github.com/elct9620/servant"
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
			installCmd(dockerCli),
			uninstallCmd(dockerCli),
		)

		return rootCmd
	}, manager.Metadata{
		SchemaVersion: "0.1.0",
		Vendor:        "Aotokitsuruya",
	})
}

func installCmd(dockerCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install the servantd on the swarm manager",
		RunE: func(cmd *cobra.Command, args []string) error {
			installer := &servant.Installer{}
			return installer.Execute(cmd.Context(), dockerCli.Client())
		},
	}

	return cmd
}

func uninstallCmd(dockerCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstall the servantd on the swarm manager",
		RunE: func(cmd *cobra.Command, args []string) error {
			uninstaller := &servant.Uninstaller{}
			return uninstaller.Execute(cmd.Context(), dockerCli.Client())
		},
	}

	return cmd
}
