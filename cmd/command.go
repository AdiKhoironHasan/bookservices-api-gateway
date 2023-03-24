package cmd

import (
	"github.com/AdiKhoironHasan/bookservices/config"
	"github.com/AdiKhoironHasan/bookservices/domain/service"
	"github.com/urfave/cli/v2"
)

// Command is a struct
type Command struct {
	conf *config.Config
	repo *service.Repositories

	CLI []*cli.Command
}

// NewCommand is a constructor
func NewCommand(options ...CommandOption) *Command {
	cmd := &Command{}

	for _, op := range options {
		op(cmd)
	}

	return cmd
}

// registerCLI is a function
func (cmd *Command) registerCLI(cmdCLI *cli.Command) {
	cmd.CLI = append(cmd.CLI, cmdCLI)
}
