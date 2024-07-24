package commands

import (
	"github.com/Vaelatern/monitoror/cli"
	initCmd "github.com/Vaelatern/monitoror/cli/commands/init"
	"github.com/Vaelatern/monitoror/cli/commands/version"
)

func AddCommands(cli *cli.MonitororCli) {
	cli.RootCmd.AddCommand(
		// INIT
		initCmd.NewInitCommand(cli),
		// VERSION
		version.NewVersionCommand(cli),
	)
}
