package app

import (
	"github.com/sirupsen/logrus"
	"log-hooks"
	"sync"
)

const appName = "MyTarget Collector"

var (
	logOnce  sync.Once
	logEntry *logrus.Entry
)

func Log() *logrus.Entry {
	logOnce.Do(func() {
		log := logrus.New()
		err := log_hooks.UsefulSetupLogrus(
			log,
			Config().MailServerAddr,
			Config().LogFormat,
			Config().LogLevel,
			appName,
			Config().ErrMailFrom,
			Config().ErrMailTo,
		)
		if err != nil {
			panic(err)
		}

		logEntry = logrus.NewEntry(log).WithField("app", appName)
	})
	return logEntry
}
