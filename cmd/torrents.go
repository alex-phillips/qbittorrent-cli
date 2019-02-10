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

	"github.com/alex-phillips/qbittorrent/lib/utils"
	"github.com/spf13/cobra"
)

// downloadSubCmd represents the freeleech command
var torrentsCmd = &cobra.Command{
	Use:   "torrents",
	Short: "Download all media from a subreddit",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		filters := map[string]string{
			"filter":   cmd.Flag("filter").Value.String(),
			"category": cmd.Flag("category").Value.String(),
		}
		torrents := Api.GetTorrents(filters)

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "Done\tDown\tUp\tETA\tRatio\tState\tName")
		for _, torrent := range torrents {
			row := fmt.Sprintf("%.0f%%\t%s\t%s\t%s\t%.2f\t%s\t%s",
				torrent.Progress*100,
				utils.ByteCountBinary(torrent.DLSpeed),
				utils.ByteCountBinary(torrent.ULSpeed),
				utils.SecondsToHuman(torrent.ETA),
				torrent.Ratio,
				torrent.State,
				torrent.Name,
			)
			fmt.Fprintln(w, row)
		}

		w.Flush()
	},
}

func init() {
	torrentsCmd.Flags().StringP("filter", "f", "", "Filter torrents by status")
	torrentsCmd.Flags().StringP("category", "c", "", "Filter torrents by category")
	rootCmd.AddCommand(torrentsCmd)
}
