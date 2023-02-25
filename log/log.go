package log

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&Formatter{
		TimestampFormat:  time.Stamp,
		NoUppercaseLevel: false,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)
}
