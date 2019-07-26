package client

import (
	"k8s.io/klog"
)

func (c *ClusteradmClient) Init(cfg ClusteradmCfg) error {
	if cfg.Bootstrap != "" {
		klog.V(2).Infof("bootstrapping the cluster with: %s", cfg.Bootstrap)
		// do the bootstrap and fetch kubeconfig
	}

	// get kubeconfig for bootstrap cluster

	for _, p := range cfg.Providers {
		klog.V(2).Infof("calling interface ClusteradmClient.Init() with provider: %s\n", p)
	}

	// do the init

	return nil
}
