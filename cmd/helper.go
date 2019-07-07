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

// SetCommandFlag sets a flag to a command
func SetCommandFlag(cmd *cobra.Command, f *Flag, required bool) {
	switch v := f.value.(type) {
	case string:
		cmd.Flags().StringP(f.key, f.shortKey, v, f.description)
	case int:
		cmd.Flags().IntP(f.key, f.shortKey, v, f.description)
	case bool:
		cmd.Flags().BoolP(f.key, f.shortKey, v, f.description)
	case []string:
		cmd.Flags().StringArrayP(f.key, f.shortKey, v, f.description)
	default:
		panic(errors.New("Failed to infer flag type for command " + f.key))
	}

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

// GetFlag gets the value of a flag
func GetFlag(cmd *cobra.Command, f *Flag, panicOnFail bool) interface{} {
	flags := cmd.Flags()

	var value interface{}
	var err error

	switch f.value.(type) {
	case string:
		value, err = flags.GetString(f.key)
	case int:
		value, err = flags.GetInt(f.key)
	case bool:
		value, err = flags.GetBool(f.key)
	case []string:
		value, err = flags.GetStringArray(f.key)
	default:
		panic(errors.New("Failed to infer flag type for command " + f.key))
	}

	if err != nil {
		if panicOnFail {
			panic(err)
		}
		return f.value
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
	projectID := config.projectID

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
