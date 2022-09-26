package flags

import "github.com/spf13/cobra"

func AddPatternFlag(cmd *cobra.Command, p *string) {
	cmd.PersistentFlags().StringVar(p, "pattern", "", "regex pattern to look for")
}

func AddReplaceFlag(cmd *cobra.Command, p *string) {
	cmd.PersistentFlags().StringVar(p, "replace", "", "regex pattern to replace with")
}
