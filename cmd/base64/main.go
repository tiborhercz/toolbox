package main

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

type Options struct {
	Value string
	Path string
	Decode bool
	Urlencoding bool
}

var (
	options Options

    rootCmd = &cobra.Command{
		Use:   "base64",
		Short: "Encode and decode base64 strings",
		Run: func(cmd *cobra.Command, args []string) {
			if string(options.Path) == "" {
				fmt.Println("Error: required flag(s) \"path\" or \"value\" either flag should be set")
				os.Exit(1)
			}

			process(options)
		},
	}
)

func init() {
	rootCmd.Flags().StringVarP(&options.Path, "path", "p", "", "Path string")
	rootCmd.Flags().StringVarP(&options.Value, "value", "v", "", "Value string")
	rootCmd.Flags().BoolVarP(&options.Decode, "decode", "d", false, "Decode")
	rootCmd.Flags().BoolVarP(&options.Urlencoding, "urlencoding", "u", false, "URLEncoding is the alternate base64 encoding defined in RFC 4648. It is typically used in URLs and file names.")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		return
	}
}

func process(options Options) {
	var (
		data = options.Value
		processedString string
	)

	if options.Path != "" {
		data = getFileContent(options.Path)
	}

	if options.Decode == false {
		processedString = encode(data, options.Urlencoding)
	} else {
		processedString = decode(data, options.Urlencoding)
	}

	output(processedString)
}

func output (value string) {
	fmt.Println(value)
}

func getFileContent(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		os.Exit(1)
	}

	return string(data)
}

func decode(value string, urlencoding bool) string {
	var decodedString []byte

	if urlencoding {
		decodedString, _ = b64.URLEncoding.DecodeString(value)
	} else {
		decodedString, _ = b64.StdEncoding.DecodeString(value)
	}

	return string(decodedString)
}

func encode(value string, urlencoding bool) string {
	var encodedString string

	if urlencoding {
		encodedString = b64.URLEncoding.EncodeToString([]byte(value))
	} else {
		encodedString = b64.StdEncoding.EncodeToString([]byte(value))
	}

	return encodedString
}