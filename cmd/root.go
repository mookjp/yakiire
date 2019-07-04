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

// Flag defines one flag for a command
type Flag struct {
	key         string
	shortKey    string
	description string
	value       interface{}
}

var cmdVersion = &Flag{"version", "v", "Get current version", false}
var cmdCredentials = &Flag{"credentials", "cred", "Set credentials path for firebase login", ""}
var cmdProjectID = &Flag{"projectId", "prj", "Set the project to work on", ""}
var cmdCollection = &Flag{"collection", "c", "Set the collection to work on", ""}
var cmdWhere = &Flag{"where", "w", "Where condition to search documents", []string{}}
var cmdLimit = &Flag{"limit", "l", "Limit the number of results", 20}
var cmdDocumentID = &Flag{"document-id", "id", "Set the document ID to work on", ""}

const (
	envCredentialsKey = "YAKIIRE_GOOGLE_APPLICATION_CREDENTIALS"
	envProjectIdKey   = "YAKIIRE_FIRESTORE_PROJECT_ID"

	version = "0.0.1-alpha"
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
	Run: func(cmd *cobra.Command, args []string) {
		if v, err := cmd.Flags().GetBool(cmdVersion.key); err == nil {
			if v {
				fmt.Print(version)
			}
			os.Exit(0)
		} else {
			panic("wrong version key")
		}
	},
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
	SetCommandFlag(rootCmd, cmdVersion, false)

	for _, f := range []*Flag{cmdCredentials, cmdProjectID} {
		rootCmd.PersistentFlags().StringP(f.key, f.shortKey, f.value.(string), f.description)
	}
}

func getConfig(cmd *cobra.Command) *cmdConfig {
	cred, _ := GetFlagString(cmd, cmdCredentials, false)
	if cred == "" {
		cred = os.Getenv(envCredentialsKey)
	}
	id, _ := GetFlagString(cmd, cmdProjectID, false)
	if id == "" {
		id = os.Getenv(envProjectIdKey)
	}

	return &cmdConfig{
		credentialPath: cred,
		projectId:      id,
	}
}
