package client

import (
	"errors"
	"containers/dto"
	"os/exec"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type containersClient struct{}

type containersClientInterface interface {
	GetContainersStats()(dto.ContainersStats, error)
	CreateContainer(service string, quantity uint)(error)
	DeleteContainer(container_id string)(error)
	GetContainersStatsByService(service string)(dto.ContainersStats, error)
}

var (
	ContainersClient containersClientInterface
)

func init() {
	ContainersClient = &containersClient{}
}

func (c *containersClient) GetContainersStats()(dto.ContainersStats, error) {

	cmd := exec.Command("bash", "client/bash/get_stats.sh")
	output, err := cmd.Output()
    if err != nil {
		log.Error(err.Error())
		return dto.ContainersStats{}, err
	}
	var containersStats dto.ContainersStats

	// Decodificar la salida JSON en la estructura
	err = json.Unmarshal(output, &containersStats)
	if err != nil {
		log.Error(err.Error())
		return dto.ContainersStats{}, err
	}

	return containersStats, nil
}

func (c *containersClient) GetContainersStatsByService(service string)(dto.ContainersStats, error) {

	cmd := exec.Command("bash", "-c", fmt.Sprintf("client/bash/get_stats_by_service.sh %s", service))
	output, err := cmd.Output()
    if err != nil {
		log.Error(err.Error())
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode := exitErr.ExitCode()
			log.Error(fmt.Sprintf("exitCode: %d", exitCode))
			switch exitCode {
			case 1:
				return dto.ContainersStats{}, errors.New(fmt.Sprintf("Service with name: '%s' does not exist", service))
			default:
				return dto.ContainersStats{}, err
			}
		} else {
			return dto.ContainersStats{}, err
		}
	}
	var containersStats dto.ContainersStats

	// Decodificar la salida JSON en la estructura
	err = json.Unmarshal(output, &containersStats)
	if err != nil {
		log.Error(err.Error())
		return dto.ContainersStats{}, err
	}

	return containersStats, nil
}

func (c *containersClient) CreateContainer(service string, quantity uint)(error) {

	cmd := exec.Command("bash", "-c", fmt.Sprintf("client/bash/scale_service.sh %s %d", service, quantity))
	err := cmd.Run()
	if err != nil {
		log.Error(err.Error())
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode := exitErr.ExitCode()
			log.Error(fmt.Sprintf("exitCode: %d", exitCode))
			switch exitCode {
			case 1:
				return errors.New(fmt.Sprintf("Service with name: '%s' does not exist", service))
			default:
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func (c *containersClient) DeleteContainer(container_id string)(error) {

	cmd := exec.Command("bash", "-c", fmt.Sprintf("client/bash/delete_container.sh %s", container_id))
	err := cmd.Run()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}