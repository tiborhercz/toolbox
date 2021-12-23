package cmd

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tiborhercz/cli-toolbox/internal/model"
	"github.com/tiborhercz/cli-toolbox/pkg/jwtdecode"
	"log"
)

var (
	jwtOptions model.JwtOptions

	jwtdecodeCmd = &cobra.Command{
		Use:   "jwtdecode",
		Short: "Decode jwt token",
		Run: func(cmd *cobra.Command, args []string) {
			jwtData, err := jwtdecode.Process(jwtOptions.Value)
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
	jwtdecodeCmd.Flags().StringVarP(&jwtOptions.Value, "value", "v", "", "Value string")
}

func prettifyJson(value []byte) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, value, "", "\t")
	if err != nil {
		return "", err
	}

	return prettyJSON.String(), nil
}
