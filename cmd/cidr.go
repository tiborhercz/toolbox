package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tiborhercz/cli-toolbox/internal/base64"
	"os"
)

var (
	base64Cmd = &cobra.Command{
		Use:   "base64",
		Short: "Encode and decode base64 strings",
		Run: func(cmd *cobra.Command, args []string) {
			if string(options.Path) == "" && string(options.Value) == "" {
				fmt.Println("Error: required flag(s) \"path\" or \"value\" either flag should be set")
				os.Exit(1)
			}

			fmt.Println(base64.Process(options.Value, options.Path, options.Decode, options.Urlencoding))
		},
	}
)

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.Flags().StringVarP(&options.Path, "path", "p", "", "Path string")
	base64Cmd.Flags().StringVarP(&options.Value, "value", "v", "", "Value string")
	base64Cmd.Flags().BoolVarP(&options.Decode, "decode", "d", false, "Decode")
	base64Cmd.Flags().BoolVarP(&options.Urlencoding, "urlencoding", "u", false, "URLEncoding is the alternate base64 encoding defined in RFC 4648. It is typically used in URLs and file names.")
}
