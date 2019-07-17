/*
Copyright © 2019 The Kubernetes Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"k8s.io/klog"

	"github.com/timothysc/clusteradm/pkg/client"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a Kubernetes cluster",
	Long:  `initialize a Kubernetes cluster`,
	Run:   runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringSlice("providers", nil, "providers to initialize")
	initCmd.MarkFlagRequired("providers")
	initCmd.Flags().Bool("local", false, "bootstrap a local cluster with kind")
}

func runInit(cmd *cobra.Command, args []string) {
	fmt.Println("performing init...")
	cc, _ := client.NewClusteradmClient()
	providers, _ := cmd.Flags().GetStringSlice("providers")
	for _, p := range providers {
		klog.V(2).Infof("calling interface ClusteradmClient.Init() with provider: %s\n", p)
	}
	cc.Init()
}