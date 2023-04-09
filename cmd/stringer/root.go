package stringer

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{

	Use:   "stringer",
	Short: "Stringer - a simple cli for strings",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops - There was an error while executing your cli, %s", err)
        os.Exit(1)
	}
}
