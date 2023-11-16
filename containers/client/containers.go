package client

import (
	"errors"
	"strings"
	"fmt"
	"context"
	"os"
	"os/exec"

	"github.com/docker/docker/api/types/container"
	log "github.com/sirupsen/logrus"

	"containers/dto"
	d "containers/docker"
)

type containersClient struct{}

type containersClientInterface interface {
	GetContainersStats()(dto.ContainersStats, error)
	CreateContainer(service string, quantity uint)(error)
	StartContainer(container_id string)(error)
	StopContainer(container_id string)(error)
	DeleteContainer(container_id string)(error)
	RestartContainer(container_id string)(error)
	GetContainersStatsByService(service string)(dto.ContainersStats, error)
}

var (
	ContainersClient containersClientInterface
)

func init() {
	ContainersClient = &containersClient{}
}

func (c *containersClient) GetContainersStats()(dto.ContainersStats, error) {

	containerIDsCmd := exec.Command("docker-compose", "-f", os.Getenv("DOCKER_FILE_PATH"), "ps", "-q", "-a")
	containerIDsOutput, err := containerIDsCmd.Output()
	if err != nil {
		log.Error("aca 1")
		return dto.ContainersStats{}, err
	}

	args := getArgsForDockerStats(containerIDsOutput)

	statsCmd := exec.Command("docker", args...)
	statsOutput, err := statsCmd.Output()
	if err != nil {
		log.Error(err)
		return dto.ContainersStats{}, err
	}

	return parseStats(statsOutput), nil
}

func (c *containersClient) GetContainersStatsByService(service string)(dto.ContainersStats, error) {

	serviceExistsCmd := exec.Command("docker-compose", "-f", os.Getenv("DOCKER_FILE_PATH"), "config", "--services")
    serviceExistsOutput, err := serviceExistsCmd.Output()
    if err != nil {
        log.Error(err.Error())
        return dto.ContainersStats{}, err
    }

    if !serviceExists(service, serviceExistsOutput) {
        return dto.ContainersStats{}, errors.New(fmt.Sprintf("Service with name: '%s' does not exist", service))
    }

    containerIDsCmd := exec.Command("docker-compose", "-f", os.Getenv("DOCKER_FILE_PATH"), "ps", "-q", service)
    containerIDsOutput, err := containerIDsCmd.Output()
    if err != nil {
        log.Error(err.Error())
        return dto.ContainersStats{}, err
    }

    containerIDs := strings.TrimSpace(string(containerIDsOutput))
    if containerIDs == "" {
        return dto.ContainersStats{}, nil 
    }

	args := getArgsForDockerStats(containerIDsOutput)

    // Obtener estadísticas de los contenedores del servicio
    statsCmd := exec.Command("docker", args...)
    statsOutput, err := statsCmd.Output()
    if err != nil {
        log.Error(err.Error())
        return dto.ContainersStats{}, err
    }


    return parseStats(statsOutput), nil
}

func (c *containersClient) CreateContainer(service string, quantity uint)(error) {

	serviceExistsCmd := exec.Command("docker-compose", "-f", os.Getenv("DOCKER_FILE_PATH"), "config", "--services")
    serviceExistsOutput, err := serviceExistsCmd.Output()
    if err != nil {
        log.Error("Error while searching the service name")
        return err
    }

    if !serviceExists(service, serviceExistsOutput) {
        return errors.New(fmt.Sprintf("Service with name: '%s' does not exist", service))
    }

	createServicecmd := exec.Command("docker-compose", "-f", os.Getenv("DOCKER_FILE_PATH"), "up", "--scale", fmt.Sprintf("%s=%d", service, quantity), "-d")
	if err := createServicecmd.Run(); err != nil {
		log.Error("Error while creating a container")
	}

	restartServiceCmd := exec.Command("docker-compose", "-f", os.Getenv("DOCKER_FILE_PATH"), "restart", fmt.Sprintf("nginx_%s", service))
	if err := restartServiceCmd.Run(); err != nil {
		log.Error("Error while restarting nginx")
	}

	return nil
}

func (c *containersClient) DeleteContainer(container_id string)(error) {

	serviceNameCmd := exec.Command("docker", "inspect", "--format", "{{ index .Config.Labels \"com.docker.compose.service\" }}", container_id)
	serviceNameOutput, err := serviceNameCmd.Output()
	if err != nil {
		log.Error("Error while searching the service name")
	}
	serviceName := strings.TrimRight(string(serviceNameOutput), "\n")

	deleteContainerCmd := exec.Command("docker", "rm", "-f", container_id)
	if err := deleteContainerCmd.Run(); err != nil {
		log.Error("Error while deleting the container")
	}

	restartServiceCmd := exec.Command("docker-compose", "-f", os.Getenv("DOCKER_FILE_PATH"), "restart", fmt.Sprintf("nginx_%s", serviceName))
	if err := restartServiceCmd.Run(); err != nil {
		log.Error("Error while restarting nginx")
	}

	return nil
}

func (c *containersClient) StartContainer(container_id string)(error) {

	serviceNameCmd := exec.Command("docker", "inspect", "--format", "{{ index .Config.Labels \"com.docker.compose.service\" }}", container_id)
	serviceNameOutput, err := serviceNameCmd.Output()
	if err != nil {
		log.Error("Error while searching the service name")
	}
	serviceName := strings.TrimRight(string(serviceNameOutput), "\n")

	deleteContainerCmd := exec.Command("docker", "start", container_id)
	if err := deleteContainerCmd.Run(); err != nil {
		log.Error("Error while starting the container")
	}

	restartServiceCmd := exec.Command("docker-compose", "-f", os.Getenv("DOCKER_FILE_PATH"), "restart", fmt.Sprintf("nginx_%s", serviceName))
	if err := restartServiceCmd.Run(); err != nil {
		log.Error("Error while restarting nginx")
	}

	return nil
}

func (c *containersClient) StopContainer(container_id string)(error) {

	serviceNameCmd := exec.Command("docker", "inspect", "--format", "{{ index .Config.Labels \"com.docker.compose.service\" }}", container_id)
	serviceNameOutput, err := serviceNameCmd.Output()
	if err != nil {
		log.Error("Error while searching the service name")
	}
	serviceName := strings.TrimRight(string(serviceNameOutput), "\n")

	deleteContainerCmd := exec.Command("docker", "stop", container_id)
	if err := deleteContainerCmd.Run(); err != nil {
		log.Error("Error while starting the container")
	}

	restartServiceCmd := exec.Command("docker-compose", "-f", os.Getenv("DOCKER_FILE_PATH"), "restart", fmt.Sprintf("nginx_%s", serviceName))
	if err := restartServiceCmd.Run(); err != nil {
		log.Error("Error while restarting nginx")
	}

	return nil
}

func (c *containersClient) RestartContainer(container_id string)(error) {

	err := d.DockerClient.ContainerRestart(context.Background(), container_id, container.StopOptions{})
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil

	// cmd := exec.Command("docker", "restart", container_id)
	// err := cmd.Run()
	// if err != nil {
	// 	log.Error(err.Error())
	// 	return err
	// }

	// return nil
}

func getArgsForDockerStats(containerIDsOutput []byte) []string {
	containerIDs := strings.TrimSpace(string(containerIDsOutput))
	ids := strings.Split(containerIDs, "\n")

	args := append([]string{"stats", "--no-stream"}, ids...)
	return args
}

func serviceExists(service string, serviceExistsOutput []byte) bool {
	services := strings.Fields(string(serviceExistsOutput))
    serviceExists := false
    for _, s := range services {
        if s == service {
            serviceExists = true
            break
        }
    }
	return serviceExists
}

func parseStats(statsOutput []byte) dto.ContainersStats{
	var containersStats dto.ContainersStats

	lines := strings.Split(strings.TrimSpace(string(statsOutput)), "\n")
	for _, line := range lines[1:] { // Omitir la primera línea (encabezado)
		fields := strings.Fields(line)
		if len(fields) < 11 {
			continue
		}

		containerStats := dto.ContainerStats{
			ContainerID: fields[0],
			Name:        fields[1],
			CPU:         fields[2],
			MemoryUsage: fields[3],
			MemoryLimit: fields[5],
			Memory:      fields[6],
			NetI:        fields[7],
			NetO:        fields[9],
			BlockI:      fields[10],
			BlockO:      fields[12],
		}

		containersStats.ContainersStats = append(containersStats.ContainersStats, containerStats)
	}

	return containersStats
}