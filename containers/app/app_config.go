package app

import (
	"os"

	"containers/utils/initializers"

	log "github.com/sirupsen/logrus"
)

func init() {
	initializers.LoadEnv()
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.Info("Starting logger system")
}
