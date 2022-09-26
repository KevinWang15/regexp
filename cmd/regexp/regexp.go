package main

import (
	"github.com/KevinWang15/regexp/cmd/regexp/cmds/match"
	"github.com/KevinWang15/regexp/cmd/regexp/cmds/match-replace"
	"github.com/KevinWang15/regexp/cmd/regexp/cmds/replace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "regexp",
}

func init() {
	rootCmd.AddCommand(match.MatchCmd)
	rootCmd.AddCommand(replace.ReplaceCmd)
	rootCmd.AddCommand(match_replace.MatchReplaceCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
