package utils

import (
	"containers/client"
	"strings"
	"strconv"
	"time"
	log "github.com/sirupsen/logrus"
)

var service_auto_scaling = "search"

func AutoScaling(){
	for {
		time.Sleep(20 * time.Second)

		containers_stats, err := client.ContainersClient.GetContainersStatsByService(service_auto_scaling)

        instances := len(containers_stats.ContainersStats)
		if instances == 0 {
			continue
		}
        cpuStr := containers_stats.ContainersStats[0].CPU
        cpuStr = strings.TrimRight(cpuStr, "%")
        cpu, err := strconv.ParseFloat(cpuStr, 64)
        if err != nil {
            log.Error("Error parsing CPU usage: ", err)
            return
        }
        if cpu > 5 {
			log.Warning("CPU Usage: ", cpu)
            err = client.ContainersClient.CreateContainer(service_auto_scaling, uint(instances) + uint(1))
			if err != nil {
				log.Error("Error while creating a new instance of search service")
				time.Sleep(2 * time.Second)
				continue
			}

            log.Info("New instance of search service was created. Total: ", instances+1)
        }
		if cpu < 2.5 && instances > 2 {
			var container_id = containers_stats.ContainersStats[instances - 1].ContainerID
			err := client.ContainersClient.DeleteContainer(container_id)
			if err != nil {
				log.Error("Error while deleting as instance of search service")
				time.Sleep(2 * time.Second)
				continue
			}

            log.Info("An instance of search service was deleted. Total: ", instances - 1)
		}
    }
}
