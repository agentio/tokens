package cmd

import (
	"github.com/agentio/tokens/cmd/key"
	"github.com/agentio/tokens/cmd/serve"
	"github.com/agentio/tokens/cmd/token"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tokens",
		Short: "the agentio/tokens tool",
	}

	cmd.AddCommand(key.Cmd())
	cmd.AddCommand(token.Cmd())
	cmd.AddCommand(serve.Cmd())
	return cmd
}
