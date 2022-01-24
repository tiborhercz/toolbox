package main

import (
	"github.com/tiborhercz/cli-toolbox/cmd"
	"github.com/tiborhercz/cli-toolbox/internal/logrus"
)

func main() {
	logrus.SetOptions()
	cmd.Execute()
}
