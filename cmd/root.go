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

	"github.com/alex-phillips/qbittorrent/lib/log"
	"github.com/alex-phillips/qbittorrent/lib/qbittorrent"
	"github.com/alex-phillips/qbittorrent/lib/utils"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	Api     *qbittorrent.Api
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "qbt",
	Short: "",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	home, _ := homedir.Dir()

	// Set global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", home+"/.qbittorrent", "config file")
	rootCmd.PersistentFlags().StringP("log-level", "l", "info", "Set log level")
	rootCmd.PersistentFlags().BoolVar(&utils.DryRun, "dry-run", false, "Don't download any files")

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		// Set the log level
		log.Init(cmd.Flag("log-level").Value.String())

		// Set up the config file location from flag and set type
		viper.SetConfigType("yaml")
		viper.SetConfigFile(cfgFile)

		// If the config file doesn't exist,
		if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
			if err := viper.SafeWriteConfigAs(cfgFile); err != nil {
				log.Error.Fatalln(err)
			}
			log.Warn.Println("Creating config file at " + cfgFile + ". You can override these variables with environment variables.")
		} else if err := viper.ReadInConfig(); err != nil {
			log.Error.Fatalln(err)
		}

		viper.AutomaticEnv() // read in environment variables that match

		Api = qbittorrent.GetApi(viper.GetString("host"), viper.GetString("username"), viper.GetString("password"))

		return nil
	}
}
