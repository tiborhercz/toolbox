//go:build darwin

package cmd

import (
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var dnsResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Flush the macOS DNS cache",
	Long: "Flush the macOS DNS cache by running:\n" +
		"  sudo dscacheutil -flushcache\n" +
		"  sudo killall -HUP mDNSResponder\n\n" +
		"You will be prompted for your password by sudo.",
	Run: func(cmd *cobra.Command, args []string) {
		steps := [][]string{
			{"sudo", "dscacheutil", "-flushcache"},
			{"sudo", "killall", "-HUP", "mDNSResponder"},
		}

		for _, step := range steps {
			logrus.Infof("Running: %s", strings.Join(step, " "))
			c := exec.Command(step[0], step[1:]...)
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			if err := c.Run(); err != nil {
				logrus.Fatalf("failed to run %s: %v", strings.Join(step, " "), err)
			}
		}

		logrus.Info("DNS cache flushed")
	},
}

func init() {
	dnsCmd.AddCommand(dnsResetCmd)
}
