package client

import (
	"k8s.io/client-go/kubernetes"
)

type ClusteradmClient struct {
	kclient kubernetes.Interface
}

func NewClusteradmClient() (*ClusteradmClient, error) {
	cc := &ClusteradmClient{
		kclient: nil,
	}
	return cc, nil
}

var _ Interface = &ClusteradmClient{}

type Interface interface {
	Init() error
	GenerateConfig() error
	Upgrade() error
	Reset() error
	Apply() error
	PreflightChecks() error
}
