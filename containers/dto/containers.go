package dto

type ContainersStats struct {
	ContainersStats []ContainerStats `json:"containers_stats"`
}

type ContainerStats struct {
	ContainerID string `json:"container_id"`
	Name        string `json:"name"`
	CPU         string `json:"cpu"`
	MemoryUsage string `json:"memory_usage"`
	MemoryLimit string `json:"memory_limit"`
	Memory      string `json:"memory"`
	NetI        string `json:"net_i"`
	NetO        string `json:"net_o"`
	BlockI      string `json:"block_i"`
	BlockO      string `json:"block_o"`
}

type CreateContainer struct {
	Service string `json:"service" binding:"required"`
	Quantity uint `json:"quantity,omitempty" binding:"omitempty,min=1"`
}