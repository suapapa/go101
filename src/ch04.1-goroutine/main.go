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
	log := log.WithField("name", "π₯ μμ€κ³΅")
	defer log.Debug("out")

	log.Infof("λ°μλ!")
	ach <- "μλμ§ν"

	log.Infof("μ§κ΅¬μΈλ€μ λμκ² νμ λͺ¨μμ€!")
	time.Sleep(3 * time.Second)
	ach <- "μκΈ°μ₯"

	time.Sleep(1 * time.Second)

	wg.Done()
}

func vegitor(ach <-chan string) {
	log := log.WithField("name", "π₯ λ² μ§ν°")
	defer log.Debug("out")

	for a := range ach {
		if a == "μκΈ°μ₯" {
			log.Errorf("μΌμ! %s", a)
			break
		}
		log.Warnf("%s μμνκ΅°", a)
	}

	wg.Done()
}
