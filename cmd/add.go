package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a document to a collection",
	Long:  `Add a single document to a collection with a random ID`,
	Run: func(cmd *cobra.Command, args []string) {
		documentStr := GetArgument(args, 0, "Document Content", true)
		doc := Unmarshal(documentStr)

		collectionName := GetFlag(cmd, cmdCollection, true).(string)

		ctx := context.Background()
		client := GetClient(ctx, cmd)

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
	SetCommandFlag(addCmd, cmdCollection, true)
}
