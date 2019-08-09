package cmd

import (
	"github.com/spf13/cobra"
)

var encodingString string

func init() {
	rootCmd.AddCommand(lfCmd)
	lfCmd.Flags().StringVarP(&encodingString, "encoding", "e", "utf8",
		`input text encoding, could be utf8, utf16, gbk, etc..
use 'encoding'' command to get all supported encoding`)
}

var lfCmd = &cobra.Command{
	Use:   "lf",
	Short: "Convert all end of line to lf(\\n)",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return handleArgs(args, []rune{'\n'}, encodingString)
	},
}
