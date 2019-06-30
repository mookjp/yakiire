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
	"github.com/mookjp/yakiire/lib"
	"os"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

const (
	collections = "collections"
	credentials = "gcp.credentials"
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

		flags := cmd.PersistentFlags()
		collectionName, err := flags.GetString(collections)
		if err != nil {
			panic(err)
		}

		ctx := context.Background()
		cred := viper.GetString(credentials)
		client, err := lib.NewClient(ctx, &lib.ClientConfig{
			Credentials: cred,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")
	getCmd.PersistentFlags().StringP(collections, "c", "", "The collection name to get a document from")
	err := getCmd.MarkFlagRequired(collections)
	if err == nil {
		panic(err)
	}

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
