package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tiborhercz/cli-toolbox/internal/model"
	"github.com/tiborhercz/cli-toolbox/pkg/base64"
	"io/ioutil"
	"os"
)

var (
	base64Options model.Base64Options

	base64Cmd = &cobra.Command{
		Use:   "base64",
		Short: "Encode and decode base64 strings",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 && string(base64Options.Path) == "" {
				logrus.Fatal("requires an argument or 'path' flag")
			}

			var value string

			if base64Options.Path != "" {
				value = getFileContent(base64Options.Path)
			} else {
				value = args[0]
			}

			v, err := base64.Process(value, base64Options.Decode, base64Options.Urlencoding)
			if err != nil {
				logrus.Fatal(err)
			}

			fmt.Println(v)
		},
	}
)

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.Flags().StringVarP(&base64Options.Path, "path", "p", "", "Path string")
	base64Cmd.Flags().BoolVarP(&base64Options.Decode, "decode", "d", false, "Decode")
	base64Cmd.Flags().BoolVarP(&base64Options.Urlencoding, "urlencoding", "u", false, "URLEncoding is the alternate base64 encoding defined in RFC 4648. It is typically used in URLs and file names.")
}

func getFileContent(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		os.Exit(1)
	}

	return string(data)
}
