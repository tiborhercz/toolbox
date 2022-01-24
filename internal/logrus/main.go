package logrus

import "github.com/sirupsen/logrus"

func SetOptions() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})
}
