package app

import (
	"os"

	"mvc-go/utils/initializers"

	log "github.com/sirupsen/logrus"
)

func init() {
	initializers.LoadEnv()
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.Info("Starting logger system")
}
