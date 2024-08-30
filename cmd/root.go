/*
Copyright © 2024 Dennis Schoepf dev@dnsc.io

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Path string `mapstructure:"path"`
}

type Config struct {
	Db DatabaseConfig `mapstructure:"database"`
}

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "freed",
	Short: "A f[r]eed aggregator with an intentionally reduced discovery mechanism.",
	Long: `A f[r]eed aggregator with an intentionally reduced discovery mechanism.

f[r]eed aims to handle all sort of different inputs that emit recurring content. This includes: RSS feeds, youtube channels, and more. If there is a type of input that you wish to be handled (or you find a bug), open an issue at https://github.com/dennisschoepf/freed/issues.

Where f[r]eed is different to other feed aggregators or bookmarking services is in its discovery mechanism. Due to the overwhelming amount of content created every day, the constant stream of input can be overwhelming. That is why only a set amount of new content from your input feeds is shown to you when discovering new items. The amount is going to be configurable in the future.

The application also includes a set of recommended "Small Web" and topic-specific curated feeds. These are also available through the configuration.
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// rootCmd.SetHelpFunc(internal.HelpFunc)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: ~/.config/freed/config.toml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		homeDir, err := os.UserHomeDir()
		cobra.CheckErr(err)

		userConfigDir, err := os.UserConfigDir()
		cobra.CheckErr(err)

		freedXdgConfigDir := filepath.Join(userConfigDir, "freed")
		freedConfigDir := filepath.Join(homeDir, ".config", "freed")

		viper.AddConfigPath(freedXdgConfigDir)
		viper.AddConfigPath(freedConfigDir)
		viper.SetConfigType("toml")
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Could not load config file: %s", err)
		os.Exit(1)
	}

	var c Config

	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("Could not read config file: %s", err)
		os.Exit(1)
	}
}
