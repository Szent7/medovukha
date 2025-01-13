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

/*
type Ports struct {
	Ip          string `json:"Ip"`
	PrivatePort uint16 `json:"PrivatePort"`
	PublicPort  uint16 `json:"PublicPort"`
	Type        string `json:"Type"`
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
}
*/
