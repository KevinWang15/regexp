package match

import (
	"fmt"
	"github.com/KevinWang15/regexp/pkg/flags"
	"github.com/KevinWang15/regexp/pkg/stdin"
	"github.com/spf13/cobra"
	"regexp"
)

var MatchCmd = &cobra.Command{
	Use: "match",
	Run: func(cmd *cobra.Command, args []string) {
		re, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}

		if err := stdin.ReadLines(func(b []byte) error {
			found := re.FindAll(b, -1)
			for i := range found {
				fmt.Println(string(found[i]))
			}
			return nil
		}); err != nil {
			panic(err)
		}
	},
}

var pattern string

func init() {
	flags.AddPatternFlag(MatchCmd, &pattern)
}
