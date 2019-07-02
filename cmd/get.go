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
	"context"
	"fmt"
	"os"

	"github.com/mookjp/yakiire/lib"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a document by document ID",
	Long:  `Get a document by document ID`,
	Run: func(cmd *cobra.Command, args []string) {
		docName := args[0]
		if docName == "" {
			fmt.Printf("doc name is required")
			os.Exit(1)
		}

		flags := cmd.Flags()
		collectionName, err := flags.GetString(cmdCollectionsKey)
		if err != nil {
			panic(err)
		}

		config := getConfig(cmd.Root())
		cred := config.credentialPath
		projectId := config.projectId
		fmt.Printf("credential path: %s\n", cred)
		fmt.Printf("project id: %s\n", projectId)

		ctx := context.Background()
		client, err := lib.NewClient(ctx, &lib.ClientConfig{
			Credentials: cred,
			ProjectID:   projectId,
		})
		if err != nil {
			fmt.Printf("error: %+v", err)
			os.Exit(1)
		}
		getCtx, _ := context.WithCancel(ctx)
		res, err := client.Get(getCtx, collectionName, docName)
		if err != nil {
			fmt.Printf("error: %+v", err)
			os.Exit(1)
		}
		ctx.Done()

		fmt.Printf("%s", res)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringP(cmdCollectionsKey, "c", "", "The collection name to get a document from")
	err := getCmd.MarkFlagRequired(cmdCollectionsKey)
	if err != nil {
		panic(err)
	}
}
