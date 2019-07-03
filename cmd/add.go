package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mookjp/yakiire/lib"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a document to a collection",
	Long:  `Add a single document to a collection with a random ID`,
	Run: func(cmd *cobra.Command, args []string) {
		documentStr := args[0]
		if documentStr == "" {
			fmt.Printf("The document you entered seems to be empty!")
			os.Exit(1)
		}

		var doc interface{}
		err := json.Unmarshal([]byte(documentStr), &doc)
		if err != nil {
			fmt.Printf("Failed to unmarshal JSON with error: %s", err)
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

		ctx := context.Background()
		client, err := lib.NewClient(ctx, &lib.ClientConfig{
			Credentials: cred,
			ProjectID:   projectId,
		})
		if err != nil {
			fmt.Printf("error: %+v", err)
			os.Exit(1)
		}

		addCtx, _ := context.WithCancel(ctx)
		res, err := client.Add(addCtx, collectionName, doc)
		if err != nil {
			fmt.Printf("error: %+v", err)
			os.Exit(1)
		}
		ctx.Done()

		fmt.Printf("%s", res)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP(cmdCollectionsKey, "c", "", "The collection name to add a document to")
	err := addCmd.MarkFlagRequired(cmdCollectionsKey)
	if err != nil {
		panic(err)
	}
}
