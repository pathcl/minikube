/*
Copyright 2015 The Kubernetes Authors All rights reserved.
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

	"github.com/docker/machine/libmachine/drivers/plugin/localbinary"
	"github.com/kubernetes/minikube/cli/machine"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Minikube is a tool for managing local Kubernetes clusters.",
	Long: `Minikube is a CLI tool that provisions and manages single-node Kubernetes
clusters optimized for development workflows.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv(localbinary.PluginEnvKey) == localbinary.PluginEnvVal {
			driverName := os.Getenv(localbinary.PluginEnvDriverName)
			machine.StartDriver(driverName)
			return
		}
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if os.Getenv(localbinary.PluginEnvKey) != localbinary.PluginEnvVal {
			localbinary.CurrentBinaryIsDockerMachine = true
		}
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}
