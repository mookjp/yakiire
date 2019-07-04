/*
Copyright © 2019 mookjp <mookjpy@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mookjp/yakiire/lib"

	"github.com/spf13/cobra"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		collectionName := GetFlag(cmd, cmdCollection, true).(string)
		limit := GetFlag(cmd, cmdLimit, true).(int)

		flags := cmd.Flags()
		jsons, err := flags.GetStringArray(cmdWhere.key)
		if err != nil {
			panic(err)
		}
		conds, err := parseJSONs(jsons)
		if err != nil {
			fmt.Printf("wrong JSON: %+v", err)
			os.Exit(1)
		}

		ctx := context.Background()
		client := GetClient(ctx, cmd)
		queryCtx, _ := context.WithCancel(ctx)
		res, err := client.Query(queryCtx, collectionName, conds, limit)
		if err != nil {
			fmt.Printf("error: %+v", err)
			os.Exit(1)
		}
		ctx.Done()

		for _, r := range res {
			fmt.Printf("%s\n", r)
		}
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
	SetCommandFlag(queryCmd, cmdCollection, true)
	SetCommandFlag(queryCmd, cmdWhere, true)
	SetCommandFlag(queryCmd, cmdLimit, false)
}

func parseJSONs(jsons []string) ([]*lib.Condition, error) {
	if len(jsons) == 0 {
		panic("no items")
	}
	res := make([]*lib.Condition, 0)
	for _, j := range jsons {
		c, err := parseJSON(j)
		if err != nil {
			return nil, err
		}
		res = append(res, c)
	}
	return res, nil
}

func parseJSON(j string) (*lib.Condition, error) {
	c := &lib.Condition{}
	err := json.Unmarshal([]byte(j), c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
