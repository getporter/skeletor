package main

import (
	"github.com/getporter/skeletor/pkg/skeletor"
	"github.com/spf13/cobra"
)

func buildLintCommand(m *skeletor.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lint",
		Short: "Execute the lint functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.PrintLintResults(cmd.Context())
		},
	}
	return cmd
}
