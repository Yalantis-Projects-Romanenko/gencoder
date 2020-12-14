package cmd

import (
	"fmt"
	"strings"

	"github.com/fdistorted/gencoder/encodings"
	"github.com/spf13/cobra"
)

func init() {
	var Encoding string
	var encode = &cobra.Command{
		Use:   "encode [text]",
		Short: "encodes text according to selected encoding",
		RunE: func(cmd *cobra.Command, args []string) error {
			text := strings.Join(args, " ")
			encodeFunc, ok := encodings.EncodeFunctions[Encoding]
			if ok {
				fmt.Println("Encoding: " + Encoding)
				fmt.Println("Text: " + text)
				fmt.Println("Result: " + encodeFunc(text))
			} else {
				fmt.Printf("Unsupported encoding %q\n", Encoding)
			}

			return nil

		},
	}

	encode.Flags().StringVarP(&Encoding, "encoding", "e", "", "one of the supported encoding")
	encode.MarkFlagRequired("encoding")

	rootCmd.AddCommand(encode)

}
