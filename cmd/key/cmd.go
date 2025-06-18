package key

import (
	"github.com/agentio/tokens/cmd/key/generate"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "key",
		Short: "create and manage JWKs",
	}
	cmd.AddCommand(generate.Cmd())
	return cmd
}
