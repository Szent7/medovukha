package types

import "github.com/docker/docker/api/types"

type ContainerBaseInfo struct {
	Id          string       `json:"id"`
	Names       []string     `json:"Names"`
	ImageName   string       `json:"Image"`
	Ports       []types.Port `json:"Ports"`
	Created     int64        `json:"Created"`
	State       string       `json:"State"`
	IsMedovukha bool         `json:"IsMedovukha"`
}

/*
type Ports struct {
	Ip          string `json:"Ip"`
	PrivatePort uint16 `json:"PrivatePort"`
	PublicPort  uint16 `json:"PublicPort"`
	Type        string `json:"Type"`
}
*/
