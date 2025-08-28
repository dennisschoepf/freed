/*
Copyright © 2025 Dennis Schoepf <dev@dnsc.io>

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
	"freed/internal"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:     "today",
	Aliases: []string{"day", "t"},
	Short:   "Shows an item for the day",
	Long:    `Shows an item from your stored feeds. Only allows for a set amount of items to be read in a day to discourage doom-scrolling.`,
	Example: `freed today`,
	Run: func(cmd *cobra.Command, args []string) {
		url, err := internal.GetArticleUrlForToday()

		if err != nil {
			pterm.Error.Printf("Could not get article for today: %v\n", err)
			return
		}

		title := pterm.LightCyan("Today's article")
		pterm.DefaultBox.WithPadding(2).WithTitle(title).Println(url)
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
