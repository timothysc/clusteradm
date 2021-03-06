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

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "removes clusteradm operator and associated components",
	Long:  `removes clusteradm operator and associated components`,
	Run:   runReset,
}

func init() {
	rootCmd.AddCommand(resetCmd)
	//TODO - determine if we should remove the local KIND cluster.
}

func runReset(cmd *cobra.Command, args []string) {
	fmt.Println("performing config...")
	klog.V(2).Infoln("calling interface ClusteradmClient.Reset()")
	cc, _ := client.NewClusteradmClient()
	cc.Reset()
}
