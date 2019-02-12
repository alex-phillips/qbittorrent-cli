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
	"regexp"

	"github.com/alex-phillips/qbittorrent/lib/log"
	"github.com/spf13/cobra"
)

// downloadSubCmd represents the freeleech command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Download torrents from files or magnet links",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var links []string
		var files []string
		params := make(map[string]string)

		if category := cmd.Flag("category").Value.String(); category != "" {
			params["category"] = category
		}
		if savepath := cmd.Flag("savepath").Value.String(); savepath != "" {
			params["savepath"] = savepath
		}

		for _, item := range args {
			if match, _ := regexp.MatchString("^magnet:", item); match == true {
				links = append(links, item)
			} else {
				files = append(files, item)
			}
		}

		for _, link := range links {
			api.UploadLink(link, params)
			log.Info.Println("Added link...")
		}
	},
}

func init() {
	addCmd.Flags().StringP("category", "c", "", "Filter torrents by category")
	addCmd.Flags().StringP("savepath", "s", "", "Set save path of filtered torrents")
	rootCmd.AddCommand(addCmd)
}
