package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("## errors demo")
	errorsDemo()

	fmt.Println("## errgroup demo")
	log.Logger.SetLevel(logrus.DebugLevel)
	errgroupDemo()

	fmt.Println("## logrus demo")
	log.Logger.SetLevel(logrus.ErrorLevel)
	logrusDemo()
}
