package replace

import (
	"fmt"
	"github.com/KevinWang15/regexp/pkg/flags"
	"github.com/KevinWang15/regexp/pkg/stdin"
	"github.com/spf13/cobra"
	"regexp"
)

var ReplaceCmd = &cobra.Command{
	Use: "replace",
	Run: func(cmd *cobra.Command, args []string) {
		re, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}

		if err := stdin.ReadLines(func(b []byte) error {
			found := re.ReplaceAll(b, []byte(replace))
			fmt.Println(string(found))

			return nil
		}); err != nil {
			panic(err)
		}
	},
}

var pattern string
var replace string

func init() {
	flags.AddPatternFlag(ReplaceCmd, &pattern)
	flags.AddReplaceFlag(ReplaceCmd, &replace)
}
