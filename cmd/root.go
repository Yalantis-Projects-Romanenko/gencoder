package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gencrypt",
		Short: "text encoder/decoder. supports piglatin and numbers",
		Long:  `this is the text encoder decoder tool that works with different encodings`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
