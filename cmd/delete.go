package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/mookjp/yakiire/lib"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a document from a collection",
	Long:  `Delete a document from a collection with specified ID`,
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

		ctx := context.Background()
		client, err := lib.NewClient(ctx, &lib.ClientConfig{
			Credentials: cred,
			ProjectID:   projectId,
		})
		if err != nil {
			fmt.Printf("error: %+v", err)
			os.Exit(1)
		}
		delCtx, _ := context.WithCancel(ctx)
		err = client.Delete(delCtx, collectionName, docName)
		if err != nil {
			fmt.Printf("error: %+v", err)
			os.Exit(1)
		}
		ctx.Done()
	},
}

func init() {
	rootCmd.AddCommand(delCmd)

	delCmd.Flags().StringP(cmdCollectionsKey, "c", "", "The collection name to delete a document from")
	err := delCmd.MarkFlagRequired(cmdCollectionsKey)
	if err != nil {
		panic(err)
	}
}
