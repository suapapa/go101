package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("## errors demo")
	errorsDemo()

	fmt.Println("## logrus demo")
	log.Logger.SetLevel(logrus.ErrorLevel)
	logrusDemo()
}
