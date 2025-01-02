package types

import "errors"

var (
	ErrContainerNotFound    = errors.New("container not found err")
	ErrContainerIsMedovukha = errors.New("container is medovukha err")
)
