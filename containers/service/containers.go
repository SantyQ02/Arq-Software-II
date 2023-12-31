package service

import (
	"containers/client"
	"containers/dto"
	"fmt"
	"time"
	"strings"
	"strconv"
	e "containers/utils/errors"
	log "github.com/sirupsen/logrus"
)

type containersService struct{}

type containersServiceInterface interface {
	GetContainersStats()(dto.ContainersStats, e.ApiError)
	CreateContainer(createDto dto.CreateContainer)(e.ApiError)
	StartContainer(container_id string)(e.ApiError)
	StopContainer(container_id string)(e.ApiError)
	DeleteContainer(container_id string)(e.ApiError)
	RestartContainer(container_id string)(e.ApiError)
}

var (
	ContainersService containersServiceInterface
)

func init() {
	ContainersService = &containersService{}
}

func (s *containersService) GetContainersStats()(dto.ContainersStats, e.ApiError) {

	containers_stats, err := client.ContainersClient.GetContainersStats()
	if err != nil {
		return dto.ContainersStats{}, e.NewInternalServerApiError("Something went wrong getting stats", err)
	}

	return containers_stats, nil
}

func (s *containersService) CreateContainer(createDto dto.CreateContainer)(e.ApiError) {

	var err error

	if createDto.Quantity != 0 {
		err = client.ContainersClient.CreateContainer(createDto.Service, createDto.Quantity)
	} else {
		containers_stats, err := client.ContainersClient.GetContainersStatsByService(createDto.Service)
		if err != nil {
			return e.NewInternalServerApiError(fmt.Sprintf("Something went wrong while creating the containers: %s", err.Error()), err)
		}
		err = client.ContainersClient.CreateContainer(createDto.Service, uint(len(containers_stats.ContainersStats)) + uint(1))
	}

	if err != nil {
		return e.NewInternalServerApiError(fmt.Sprintf("Something went wrong while creating the containers: %s", err.Error()), err)
	}

	return nil
}

func (s *containersService) DeleteContainer(container_id string)(e.ApiError) {

	err := client.ContainersClient.DeleteContainer(container_id)
	if err != nil {
		return e.NewInternalServerApiError("Something went wrong while deleting the container", err)
	}

	return nil
}

func (s *containersService) StartContainer(container_id string)(e.ApiError) {

	err := client.ContainersClient.StartContainer(container_id)
	if err != nil {
		return e.NewInternalServerApiError("Something went wrong while starting the container", err)
	}

	return nil
}

func (s *containersService) StopContainer(container_id string)(e.ApiError) {

	err := client.ContainersClient.StopContainer(container_id)
	if err != nil {
		return e.NewInternalServerApiError("Something went wrong while stopping the container", err)
	}

	return nil
}

func (s *containersService) RestartContainer(container_id string)(e.ApiError) {

	err := client.ContainersClient.RestartContainer(container_id)
	if err != nil {
		return e.NewInternalServerApiError("Something went wrong while restarting the container", err)
	}

	return nil
}

func AutoScaling(service_auto_scaling string){
	for {

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
        if cpu > 50 {
			log.Warning(fmt.Sprintf("%s: CPU Usage: ", service_auto_scaling), cpu)
            err = client.ContainersClient.CreateContainer(service_auto_scaling, uint(instances) + uint(1))
			if err != nil {
				log.Error(fmt.Sprintf("Error while creating a new instance of %s service", service_auto_scaling))
				time.Sleep(2 * time.Second)
				continue
			}

            log.Info(fmt.Sprintf("New instance of %s service was created. Total: ", service_auto_scaling), instances+1)
        }
		if cpu < 10 && instances > 2 {
			var container_id = containers_stats.ContainersStats[instances - 1].ContainerID
			err := client.ContainersClient.DeleteContainer(container_id)
			if err != nil {
				log.Error(fmt.Sprintf("Error while deleting as instance of %s service", service_auto_scaling))
				time.Sleep(2 * time.Second)
				continue
			}

            log.Info(fmt.Sprintf("An instance of %s service was deleted. Total: ", service_auto_scaling), instances - 1)
		}

		time.Sleep(15 * time.Second)
    }
}