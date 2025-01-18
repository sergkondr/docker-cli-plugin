package main

import (
	"github.com/docker/cli/cli-plugins/manager"
	"github.com/docker/cli/cli-plugins/plugin"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
)

var (
	version string
)

func main() {
	plugin.Run(
		func(dockerCli command.Cli) *cobra.Command {
			cmd := &cobra.Command{
				Use:   "cli-plugin NAME[:TAG|@DIGEST] [DIRECTORY]",
				Short: "Example of cli-plugin",
				Long:  `This command does nothing`,
				Run: func(cmd *cobra.Command, args []string) {

				},
			}

			flags := cmd.Flags()
			_ = flags

			cmd.AddCommand()
			return cmd
		},

		manager.Metadata{
			SchemaVersion:    "0.1.0",
			Vendor:           "Sergei Kondrashov",
			Version:          version,
			ShortDescription: "Docker cli-plugin",
			URL:              "https://github.com/sergkondr/docker-cli-plugin",
		})
}
