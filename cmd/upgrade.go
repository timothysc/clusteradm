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

// upgradeCmd represents the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "perform an upgrade of the clusteradm operator on the Management cluster",
	Long:  `perform an upgrade of the clusteradm operator on the Management cluster`,
	Run:   runUpgrade,
}

// TODO - should we support the entire CRUD workflow and enable upgrades of target cluster.
// QOTD - Fuse mount but for clusters.
// Root -- Your root kubeconfig would need to know all the mgmt clusters.
//  MC1
//    Target1
// 	MC2
//   Target2
//
// Should we repl this thing?

func init() {
	rootCmd.AddCommand(upgradeCmd)
}

func runUpgrade(cmd *cobra.Command, args []string) {
	fmt.Println("performing upgrade...")
	klog.V(2).Infoln("calling interface ClusteradmClient.Upgrade()")
	cc, _ := client.NewClusteradmClient()
	cc.Upgrade()
}
