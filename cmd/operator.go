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
	"runtime"

	"k8s.io/klog"

	"github.com/spf13/cobra"
)

// operatorCmd represents the operator command
var operatorCmd = &cobra.Command{
	Use:   "operator",
	Short: "operator starts a clusteradm operator",
	Long:  `operator starts a clusteradm operator`,
	Run: runOperator,
}

func init() {
	rootCmd.AddCommand(operatorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// operatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// operatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func runOperator(cmd *cobra.Command, args []string) {
	fmt.Println("operator called")
	printVersion()
}

func printVersion() {
	klog.V(0).Infof(fmt.Sprintf("Go Version: %s", runtime.Version()))
	klog.V(0).Infof(fmt.Sprintf("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH))
}
