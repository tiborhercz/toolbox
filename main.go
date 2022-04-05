package main

import (
	"github.com/tiborhercz/toolbox/cmd"
	"github.com/tiborhercz/toolbox/internal/logrus"
)

func main() {
	logrus.SetOptions()
	cmd.Execute()
}
