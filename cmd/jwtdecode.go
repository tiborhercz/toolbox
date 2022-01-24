package cmd

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tiborhercz/cli-toolbox/pkg/jwtdecode"
	"log"
)

var (
	jwtdecodeCmd = &cobra.Command{
		Use:   "jwtdecode",
		Short: "Decode jwt token",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				logrus.Fatal("requires an argument. Example argument: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
			}

			jwtData, err := jwtdecode.Process(args[0])
			if err != nil {
				log.Fatal(err)
			}

			for _, value := range jwtData {
				prettyJson, _ := prettifyJson(value)
				logrus.Infof("\n%v", prettyJson)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(jwtdecodeCmd)
}

func prettifyJson(value []byte) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, value, "", "\t")
	if err != nil {
		return "", err
	}

	return prettyJSON.String(), nil
}
