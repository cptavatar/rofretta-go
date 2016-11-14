package theory

import (
	"os"
	"github.com/Sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	// Output to stderr instead of stdout, could also be a file.
	logrus.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.DebugLevel)
}