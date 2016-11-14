package theory

import (
	"os"
	"github.com/Sirupsen/logrus"
)

// Create a shared logger so we can control in one spot
var Log = logrus.New()

const Disabled = -1

func init() {
	// Output to stderr instead of stdout, could also be a file.
	logrus.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.DebugLevel)
}