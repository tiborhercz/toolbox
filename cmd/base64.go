package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tiborhercz/cli-toolbox/internal/base64"
	"github.com/tiborhercz/cli-toolbox/internal/model"
	"os"
)

var (
	base64Options model.Base64Options

	base64Cmd = &cobra.Command{
		Use:   "base64",
		Short: "Encode and decode base64 strings",
		Run: func(cmd *cobra.Command, args []string) {
			if string(base64Options.Path) == "" && string(base64Options.Value) == "" {
				fmt.Println("Error: required flag(s) \"path\" or \"value\" either flag should be set")
				os.Exit(1)
			}

			fmt.Println(base64.Process(base64Options.Value, base64Options.Path, base64Options.Decode, base64Options.Urlencoding))
		},
	}
)

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.Flags().StringVarP(&base64Options.Path, "path", "p", "", "Path string")
	base64Cmd.Flags().StringVarP(&base64Options.Value, "value", "v", "", "Value string")
	base64Cmd.Flags().BoolVarP(&base64Options.Decode, "decode", "d", false, "Decode")
	base64Cmd.Flags().BoolVarP(&base64Options.Urlencoding, "urlencoding", "u", false, "URLEncoding is the alternate base64 encoding defined in RFC 4648. It is typically used in URLs and file names.")
}
