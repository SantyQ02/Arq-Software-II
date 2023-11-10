package app

import (
	"mvc-go/utils/initializers"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	initializers.LoadEnv()
	log.SetOutput(os.Stdout)
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.Info("Starting logger system")
}
