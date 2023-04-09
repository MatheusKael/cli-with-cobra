package stringer

import (
	"fmt"

	"example.com/cobra/pkg/stringer"
	"github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
	Use:   "reverse",
	Short: "reverse string",
	Run: func(cmd *cobra.Command, args []string) {

		res := stringer.Reverse(args[0])
		fmt.Println(res)

	},
}

func init() {

	rootCmd.AddCommand(reverseCmd)
}
