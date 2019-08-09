package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(crCmd)
	crCmd.Flags().StringVarP(&encodingString, "encoding", "e", "utf8",
		`input text encoding, could be utf8, utf16, gbk, etc..
use 'encoding'' command to get all supported encoding`)
}

var crCmd = &cobra.Command{
	Use:   "cr",
	Short: "Convert all end of line to cr(\\r)",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return handleArgs(args, []rune{'\r'}, encodingString)
	},
}
