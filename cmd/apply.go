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

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply new configuration to a cluster",
	Long:  `Apply new configuration to a cluster`,
	Run:   runApply,
}

func init() {
	rootCmd.AddCommand(applyCmd)
}

func runApply(cmd *cobra.Command, args []string) {
	fmt.Println("performing apply...")
	klog.V(2).Infoln("calling interface ClusteradmClient.Apply()")
	cc, _ := client.NewClusteradmClient()
	cc.Apply()
}
