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

	"github.com/timothysc/clusteradm/pkg/client"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize the Management Cluster",
	Long:  `Runs clusteradm operator on the Management Cluster`,
	Run:   runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringSlice("providers", nil, "providers to initialize")
	initCmd.Flags().String("bootstrap", "", "provider used to bootstrap")
	initCmd.MarkFlagRequired("providers")
	initCmd.Flags().Lookup("bootstrap").NoOptDefVal = "kind"
	// TODO - determine bootstrap/pivot scenario
}

func runInit(cmd *cobra.Command, args []string) {

	// TODO preflight checks to determine if KUBECONFIG exists for local=false.
	// if not, print long help thing.
	// 1. If it's already running exit with note on the version running.
	fmt.Println("performing init...")

	config := client.ClusteradmCfg{}

	bootstrap, _ := cmd.Flags().GetString("bootstrap")
	config.Bootstrap = bootstrap

	cc, _ := client.NewClusteradmClient()
	providers, _ := cmd.Flags().GetStringSlice("providers")
	config.Providers = providers
	cc.Init(config)
}
