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
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/tabwriter"

	"github.com/alex-phillips/qbittorrent/lib/log"
	"github.com/alex-phillips/qbittorrent/lib/qbittorrent"
	"github.com/alex-phillips/qbittorrent/lib/utils"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var (
	pause  bool
	resume bool
	delete bool
	purge  bool
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

		// Build search filters based on flag
		searchFilters := make(map[string]string)
		for _, search := range strings.Split(cmd.Flag("search").Value.String(), ",") {
			if search != "" {
				searchPair := strings.Split(search, "=")
				searchFilters[searchPair[0]] = searchPair[1]
			}
		}

		results := gjson.Parse(Api.GetTorrents(filters))

		var filtered []qbittorrent.Torrent
		results.ForEach(func(key, t gjson.Result) bool {
			if len(searchFilters) > 0 {
				for field, regex := range searchFilters {
					if match, _ := regexp.MatchString("(?i)"+regex, gjson.Get(t.String(), field).String()); match == false {
						return true
					}
				}
			}

			var torrent qbittorrent.Torrent
			json.Unmarshal([]byte(t.String()), &torrent)
			filtered = append(filtered, torrent)

			return true
		})

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		// fmt.Fprintln(w, "Done\tDown\tUp\tETA\tRatio\tState\tName")
		for _, torrent := range filtered {
			if cmd.Flag("savepath").Value.String() != "" {
				if _, err := Api.SetSavePath(torrent.Hash, cmd.Flag("savepath").Value.String()); err != nil {
					log.Error.Fatalln(err)
				}

				fmt.Fprintln(w, fmt.Sprintf("MOVED\t%s", torrent.Name))
			}

			if pause == true {
				if _, err := Api.Pause(torrent.Hash); err != nil {
					log.Error.Fatalln(err)
				}
				fmt.Fprintln(w, fmt.Sprintf("PAUSING\t%s", torrent.Name))
			} else if resume == true {
				if _, err := Api.Resume(torrent.Hash); err != nil {
					log.Error.Fatalln(err)
				}
				fmt.Fprintln(w, fmt.Sprintf("RESUMING\t%s", torrent.Name))
			} else if delete == true {
				if _, err := Api.Delete(torrent.Hash); err != nil {
					log.Error.Fatalln(err)
				}
				fmt.Fprintln(w, fmt.Sprintf("DELETED\t%s", torrent.Name))
			} else if purge == true {
				if _, err := Api.DeletePermanently(torrent.Hash); err != nil {
					log.Error.Fatalln(err)
				}
				fmt.Fprintln(w, fmt.Sprintf("PURGED\t%s", torrent.Name))
			} else {
				fmt.Fprintln(w, fmt.Sprintf("%.0f%%\t%s\t%s\t%s\t%.2f\t%s\t%s",
					torrent.Progress*100,
					utils.ByteCountBinary(torrent.DLSpeed),
					utils.ByteCountBinary(torrent.ULSpeed),
					utils.SecondsToHuman(torrent.ETA),
					torrent.Ratio,
					torrent.State,
					torrent.Name,
				))
			}
		}

		w.Flush()
	},
}

func init() {
	torrentsCmd.Flags().StringP("filter", "f", "", "Filter torrents by status")
	torrentsCmd.Flags().StringP("category", "c", "", "Filter torrents by category")
	torrentsCmd.Flags().StringP("search", "S", "", "Filter torrents by category")
	torrentsCmd.Flags().StringP("savepath", "s", "", "Set save path of filtered torrents")

	torrentsCmd.Flags().BoolVar(&pause, "pause", false, "Pause all filtered torrents")
	torrentsCmd.Flags().BoolVar(&resume, "resume", false, "Resume all filtered torrents")
	torrentsCmd.Flags().BoolVar(&delete, "delete", false, "Delete all filtered torrents")
	torrentsCmd.Flags().BoolVar(&purge, "purge", false, "Delete and remove data of all filtered torrents")

	rootCmd.AddCommand(torrentsCmd)
}
