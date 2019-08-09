package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(crlfCmd)
	crlfCmd.Flags().StringVarP(&encodingString, "encoding", "e", "utf8",
		`input text encoding, could be utf8, utf16, gbk, etc..
use 'encoding'' command to get all supported encoding`)
}

var crlfCmd = &cobra.Command{
	Use:   "crlf",
	Short: "Convert all end of line to crlf(\\r\\n)",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return handleArgs(args, []rune{'\r', '\n'}, encodingString)
	},
}
