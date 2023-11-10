package service

import (
	"containers/client"
	"containers/dto"
	"fmt"
	e "containers/utils/errors"
	// log "github.com/sirupsen/logrus"
)

type containersService struct{}

type containersServiceInterface interface {
	GetContainersStats()(dto.ContainersStats, e.ApiError)
	CreateContainer(createDto dto.CreateContainer)(e.ApiError)
	DeleteContainer(container_id string)(e.ApiError)
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
		return e.NewInternalServerApiError("Something went wrong while creating the container", err)
	}

	return nil
}
