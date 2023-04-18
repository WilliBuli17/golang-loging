package golang_loging

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type SampleHook struct {
}

func (s *SampleHook) Levels() []logrus.Level {
	loggerHook := []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
	return loggerHook
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Sample Hook", entry.Level, entry.Message)
	return nil // implementasi disini jika ingin kirim log error ke email atau wa atau tele
}
