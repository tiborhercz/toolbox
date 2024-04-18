package cmd

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tiborhercz/toolbox/internal/model"
	"github.com/tiborhercz/toolbox/static"
)

var (
	webUIOptions model.WebUIOptions

	webUICmd = &cobra.Command{
		Use:   "webui",
		Short: "Launch the web UI for the Toolbox",
		Run: func(cmd *cobra.Command, args []string) {
			port, err := getAvailablePort("tcp4", webUIOptions.Port, 5)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Can't listen on port %q: %s\n", port, err)
				os.Exit(1)
			}

			static.ServeHttp(port)
		},
	}
)

func init() {
	rootCmd.AddCommand(webUICmd)
	webUICmd.Flags().StringVarP(&webUIOptions.Port, "port", "p", "8000", "Port number")
}

func getAvailablePort(protocol string, port string, limit int) (string, error) {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println(fmt.Errorf("error %s", err))
		return "", err
	}
	fmt.Println(portInt)

	for try := 0; try <= limit; try++ {
		fmt.Println("0.0.0.0:" + strconv.Itoa(portInt))
		ln, err := net.Listen(protocol, "0.0.0.0:"+strconv.Itoa(portInt))

		if err != nil && strings.Contains(err.Error(), "address already in use") {
			portInt++
			fmt.Println(portInt)
			continue
		}

		_ = ln.Close()

		break
	}

	return strconv.Itoa(portInt), nil
}
