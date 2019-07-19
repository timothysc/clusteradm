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
	"os"
	gort "runtime"

	"github.com/spf13/cobra"

	clusteradmv1 "github.com/timothysc/clusteradm/api/v1"
	"github.com/timothysc/clusteradm/controllers"
	"k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	// +kubebuilder:scaffold:imports
)

// operatorCmd represents the operator command
var operatorCmd = &cobra.Command{
	Use:   "operator",
	Short: "operator starts a clusteradm operator",
	Long:  `operator starts a clusteradm operator`,
	Run:   runOperator,
}

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
	version  = "v0.0.0" // TODO: set up a version file
)

func init() {
	rootCmd.AddCommand(operatorCmd)

	operatorCmd.Flags().Bool("enable-leader-election", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
	operatorCmd.Flags().String("metrics-addr", ":8080", "The address the metric endpoint binds to.")

	clusteradmv1.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

func runOperator(cmd *cobra.Command, args []string) {
	banner := fmt.Sprintf("Clusteradm operator %s (Go Version: %s running on %s/%s)", version, gort.Version(), gort.GOOS, gort.GOARCH)

	// TODO: error checking
	metricsAddr, _ := cmd.Flags().GetString("metrics-addr")
	enableLeaderElection, _ := cmd.Flags().GetBool("enable-leader-election")

	ctrl.SetLogger(zap.Logger(true))

	// set up logging for operator
	logging := ctrl.Log.WithName("controllers").WithName("Provider")

	logging.Info(banner)
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		LeaderElection:     enableLeaderElection,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	err = (&controllers.ProviderReconciler{
		Client: mgr.GetClient(),
		Log:    logging,
	}).SetupWithManager(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Provider")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
