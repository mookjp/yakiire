/*
Copyright © 2019 mookjp <mookjpy@gmail.com>

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

	"github.com/spf13/cobra"
)

const (
	cmdCredentialsKey = "credentials"
	cmdProjectIdKey   = "projectId"
	cmdCollectionsKey = "collection"
	cmdWhereKey       = "where"
	cmdLimitKey       = "limit"

	envCredentialsKey = "YAKIIRE_GOOGLE_APPLICATION_CREDENTIALS"
	envProjectIdKey   = "YAKIIRE_FIRESTORE_PROJECT_ID"
)

type cmdConfig struct {
	credentialPath string
	projectId      string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "yakiire",
	Short: "a small CLI for Google Firestore",
	Long:  `ex) yakiire get -c products ABC`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().String(cmdCredentialsKey, "", "Google Application Credential path")
	rootCmd.PersistentFlags().String(cmdProjectIdKey, "", "Firestore project ID")
}

func getConfig(cmd *cobra.Command) *cmdConfig {
	cred, err := cmd.PersistentFlags().GetString(cmdCredentialsKey)
	if err != nil {
		panic(err)
	}
	if cred == "" {
		cred = os.Getenv(envCredentialsKey)
	}
	id, err := cmd.PersistentFlags().GetString(cmdProjectIdKey)
	if err != nil {
		panic(err)
	}
	if id == "" {
		id = os.Getenv(envProjectIdKey)
	}
	return &cmdConfig{
		credentialPath: cred,
		projectId:      id,
	}
}
