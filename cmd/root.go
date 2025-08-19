/*
Copyright © 2024 Dennis Schoepf <dev@dnsc.io>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"freed/internal/database"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Path string `mapstructure:"path"`
}

type Config struct {
	DB DatabaseConfig `mapstructure:"database"`
}

type AppContext struct {
	Config Config
}

var (
	appContext = &AppContext{}
	cfgFile    string
	rootCmd    = &cobra.Command{
		Use:   "freed",
		Short: "A f[r]eed aggregator with an intentionally reduced discovery mechanism.",
		Long: `A f[r]eed aggregator with an intentionally reduced discovery mechanism.

f[r]eed aims to handle all sort of different inputs that emit recurring content. This includes: RSS feeds, youtube channels, and more. If there is a type of input that you wish to be handled (or you find a bug), open an issue at https://github.com/dennisschoepf/freed/issues.

Where f[r]eed is different to other feed aggregators or bookmarking services is in its discovery mechanism. Due to the overwhelming amount of content created every day, the constant stream of input can be overwhelming. That is why only a set amount of new content from your input feeds is shown to you when discovering new items. The amount is going to be configurable in the future.

The application also includes a set of recommended "Small Web" and topic-specific curated feeds. These are also available through the configuration.
	`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			initDB()
		},
	}
)

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: ~/.config/freed/config.toml)")
}

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

	viper.SetDefault("database.path", "freed.db")

	if err := viper.Unmarshal(&appContext.Config); err != nil {
		fmt.Printf("Could not read config file: %s", err)
		os.Exit(1)
	}
}

func initDB() {
	err := database.Open(appContext.Config.DB.Path)

	if err != nil {
		fmt.Printf("Could not connect to the database: %s", err)
	}
}
