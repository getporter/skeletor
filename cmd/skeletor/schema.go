package main

import (
	"get.porter.sh/mixin/skeletor/pkg/skeletor"
	"github.com/spf13/cobra"
)

func buildSchemaCommand(m *skeletor.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema",
		Short: "Print the json schema for the mixin",
		Run: func(cmd *cobra.Command, args []string) {
			m.PrintSchema()
		},
	}
	return cmd
}
