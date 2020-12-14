package cmd

import (
	"fmt"
	"github.com/fdistorted/gencoder/encodings"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	var Encoding string
	var encode = &cobra.Command{
		Use:   "decode [text]",
		Short: "decodes text according to selected encoding",
		RunE: func(cmd *cobra.Command, args []string) error {
			text := strings.Join(args, " ")
			decodeFunc, ok := encodings.DecodeFunctions[Encoding]
			if ok {
				fmt.Println("Encoding: " + Encoding)
				fmt.Println("Text: " + text)
				fmt.Println("Result: " + decodeFunc(text))
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
