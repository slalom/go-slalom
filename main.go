package main

import (
	"github.com/sirupsen/logrus"
	"go-slalom/cmd"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	cmd.Execute()
}
