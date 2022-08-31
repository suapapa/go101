package main

import (
	"github.com/sirupsen/logrus"
)

var (
	log         *logrus.Entry
	programName = "thirdparty_pkg_demos"
	programVer  = "0.0.1"
)

func init() {
	initLogger()
}

func initLogger() {
	log = logrus.New().WithFields(logrus.Fields{
		"program": programName,
		"ver":     programVer,
	})

	// if !flagDebug {
	// 	log.Logger.SetLevel(logrus.InfoLevel)
	// }

	// // ES hook
	// esClient, err := elastic.NewSimpleClient(elastic.SetURL(esAddr))
	// if err != nil {
	// 	panic(err)
	// }
	// esHook, err := elogrus.NewAsyncElasticHook(esClient, hostname, logrus.InfoLevel, esIdx)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Logger.AddHook(esHook)
}

func logrusDemo() {
	log.Debug("debug msg")
	log.Info("info msg")
	log.Warn("warn msg")
	log.Error("err msg")
	log.Fatal("fatal msg")
}
