package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func init() {
	logrus.SetFormatter(&Formatter{
		TimestampFormat:  time.Stamp,
		NoUppercaseLevel: false,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)
}
