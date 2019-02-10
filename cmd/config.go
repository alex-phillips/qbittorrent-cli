// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/tidwall/gjson"

	"github.com/spf13/cobra"
)

// downloadSubCmd represents the freeleech command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Get and set qBittorrent preferences",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			Api.SetPreference(args[0], args[1])
			fmt.Println("Successfully set " + args[0] + " to " + args[1])
		} else {
			result := gjson.Parse(Api.GetPreferences())

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			result.ForEach(func(key, value gjson.Result) bool {
				if len(args) == 0 || args[0] == key.String() {
					fmt.Fprintln(w, key.String()+"\t"+value.String())

					if len(args) == 1 {
						return false
					}
				}

				return true
			})

			w.Flush()
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
