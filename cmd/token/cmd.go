package token

import (
	"github.com/agentio/tokens/cmd/token/check"
	"github.com/agentio/tokens/cmd/token/generate"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token",
		Short: "create and manage JWTs",
	}
	cmd.AddCommand(check.Cmd())
	cmd.AddCommand(generate.Cmd())

	return cmd
}
