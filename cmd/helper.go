package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/mookjp/yakiire/lib"
	"github.com/spf13/cobra"
)

// SetStringCommandFlag sets a string flag to a command
func SetStringCommandFlag(cmd *cobra.Command, f *Flag, required bool) {
	cmd.Flags().StringP(f.key, f.shortKey, f.value.(string), f.description)
	if required {
		MarkFlagRequired(cmd, f)
	}
}

// SetBoolCommandFlag sets a toggle flag to a command
func SetBoolCommandFlag(cmd *cobra.Command, f *Flag, required bool) {
	cmd.Flags().BoolP(f.key, f.shortKey, f.value.(bool), f.description)
	if required {
		MarkFlagRequired(cmd, f)
	}
}

// SetIntCommandFlag sets a integer flag to a command
func SetIntCommandFlag(cmd *cobra.Command, f *Flag, required bool) {
	cmd.Flags().IntP(f.key, f.shortKey, f.value.(int), f.description)
	if required {
		MarkFlagRequired(cmd, f)
	}
}

// MarkFlagRequired marks a flag as required
func MarkFlagRequired(cmd *cobra.Command, f *Flag) {
	err := cmd.MarkFlagRequired(f.key)
	if err != nil {
		panic(err)
	}
}

// GetFlagString gets the input string from a flag
func GetFlagString(cmd *cobra.Command, f *Flag, panicOnFail bool) (string, error) {
	flags := cmd.Flags()
	value, err := flags.GetString(f.key)
	if err != nil {
		if panicOnFail {
			panic(err)
		}
		return "", err
	}
	return value, nil
}

// GetFlagInt gets input integer from a flag
func GetFlagInt(cmd *cobra.Command, f *Flag, panicOnFail bool) int {
	flags := cmd.Flags()
	value, err := flags.GetInt(f.key)
	if err != nil {
		if panicOnFail {
			panic(err)
		}
		return 0
	}
	return value
}

// GetArgument gets a directly passed argument with index
func GetArgument(args []string, index int, name string, required bool) string {
	if len(args) <= index || args[index] == "" {
		if required {
			panic(errors.New("Required argument not found: " + name))
		}
		return ""
	}
	return args[index]
}

// GetClient gets a new lib/client for firestore calls
func GetClient(ctx context.Context, cmd *cobra.Command) *lib.Client {
	config := getConfig(cmd.Root())
	cred := config.credentialPath
	projectID := config.projectId

	client, err := lib.NewClient(ctx, &lib.ClientConfig{
		Credentials: cred,
		ProjectID:   projectID,
	})
	if err != nil {
		fmt.Printf("error: %+v", err)
		os.Exit(1)
	}
	return client
}

// Unmarshal creates a JSON interface from a string document
func Unmarshal(jsonStr string) map[string]interface{} {
	var doc map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &doc)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON with error: %s", err)
		os.Exit(1)
	}
	return doc
}
