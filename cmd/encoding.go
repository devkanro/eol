package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(encodingCmd)
}

var encodingCmd = &cobra.Command{
	Use:   "encoding",
	Short: "Get all supported encoding",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var count = 0
		for key, _ := range EncodingMap {
			count++
			print(key)

			if count >= 5 {
				print("\n")
			} else {
				print("\t")
			}
		}
	},
}
