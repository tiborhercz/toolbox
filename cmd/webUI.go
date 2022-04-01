package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tiborhercz/cli-toolbox/internal/model"
	"github.com/tiborhercz/cli-toolbox/static"
	"net"
	"os"
)

var (
	webUIOptions model.WebUIOptions

	webUICmd = &cobra.Command{
		Use:   "webui",
		Short: "Launch the web UI for the Toolbox",
		Run: func(cmd *cobra.Command, args []string) {
			exitOnPortInUse("tcp4", webUIOptions.Port)

			static.ServeHttp(webUIOptions.Port)
		},
	}
)

func init() {
	rootCmd.AddCommand(webUICmd)
	webUICmd.Flags().StringVarP(&webUIOptions.Port, "port", "p", "8000", "Port number")
}

func exitOnPortInUse(protocol string, port string) {
	ln, err := net.Listen(protocol, "0.0.0.0:"+port)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't listen on port %q: %s\n", port, err)
		os.Exit(1)
	}

	err = ln.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't stop listening on port %q: %s\n", port, err)
		os.Exit(1)
	}
}
