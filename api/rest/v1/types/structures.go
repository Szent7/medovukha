package types

import "github.com/docker/docker/api/types"

type ContainerBaseInfo struct {
	Id          string       `json:"Id"`
	Names       []string     `json:"Names"`
	ImageName   string       `json:"Image"`
	Ports       []types.Port `json:"Ports"`
	Created     int64        `json:"Created"`
	State       string       `json:"State"`
	IsMedovukha bool         `json:"IsMedovukha"`
}

type ImageBaseInfo struct {
	Id      string   `json:"Id"`
	Tags    []string `json:"Tags"`
	Size    int64    `json:"Size"`
	Created int64    `json:"Created"`
}

type NetworkBaseInfo struct {
	Name          string   `json:"Name"`
	Id            string   `json:"Id"`
	Driver        string   `json:"Driver"`
	EnableIPv6    bool     `json:"EnableIPv6"`
	IPAMDriver    string   `json:"IPAMDriver"`
	Subnet        []string `json:"Subnet"`
	Gateway       []string `json:"Gateway"`
	Attachable    bool     `json:"Attachable"`
	DockerNetwork bool     `json:"DockerNetwork"`
}

type VolumeBaseInfo struct {
	Name       string `json:"Name"`
	Driver     string `json:"Driver"`
	Mountpoint string `json:"Mountpoint"`
	Created    string `json:"Created"`
}
