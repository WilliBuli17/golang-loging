package golang_loging

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("Hello Logger")
}

func TestLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel) // ini contoh untuk kita mau ambil log dari level mana, karena secara default itu info -- opsional
	//logger.SetLevel(logrus.WarnLevel)

	logger.Trace("Trace")
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")
	//logger.Fatal("Fatal")
	//logger.Panic("Panic")
}

func TestOutput(t *testing.T) {
	logger := logrus.New()
	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.Panic(err)
	}

	logger.SetOutput(file)

	logger.Info("Hallo File Logger")
	logger.Warn("Hallo File Logger")
	logger.Error("Hallo File Logger")
}

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		logger.Panic(err)
	}

	logger.SetOutput(file)

	logger.Info("Hallo File Logger")
	logger.Warn("Hallo File Logger")
	logger.Error("Hallo File Logger")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		logger.Panic(err)
	}

	logger.SetOutput(file)

	logger.WithFields(logrus.Fields{
		"username": "wlb",
		"name":     "Willi Buli",
	}).Infof("Hallo logger")

	logger.WithFields(logrus.Fields{
		"username": "wlb2",
		"name":     "Willi Buli 2",
	}).Infof("Hallo logger 2")
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})

	logger.Warn("Hallo logger")
}
