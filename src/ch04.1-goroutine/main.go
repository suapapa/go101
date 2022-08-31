package main

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Entry
	wg  sync.WaitGroup
)

func main() {
	log = logrus.New().WithFields(logrus.Fields{
		"program": "dragon-ball",
	})
	log.Logger.SetLevel(logrus.DebugLevel)
	defer log.Debugf("exit dragon-ball")

	attackCh := make(chan string)
	wg.Add(2) // sync.WaitGroup
	go vegitor(attackCh)
	go goku(attackCh)

	wg.Wait()
}

func goku(ach chan string) {
	log := log.WithField("name", "🔥 손오공")
	defer log.Debug("out")

	log.Infof("받아랏!")
	ach <- "에너지파"

	log.Infof("지구인들아 나에게 힘을 모아줘!")
	time.Sleep(3 * time.Second)
	ach <- "원기옥"

	time.Sleep(1 * time.Second)

	wg.Done()
}

func vegitor(ach <-chan string) {
	log := log.WithField("name", "🔥 베지터")
	defer log.Debug("out")

	for a := range ach {
		if a == "원기옥" {
			log.Errorf("으악! %s", a)
			break
		}
		log.Warnf("%s 시시하군", a)
	}

	wg.Done()
}
