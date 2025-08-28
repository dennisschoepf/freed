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
	"freed/internal"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.ExactArgs(1),
	Short: "Adds a new feed to the application.",
	Long: `Validates and stores a feed in the application's database. Depending on the feed type, articles, videos, or updates are fetched right away.

Supported types currently are:
- RSS feeds
- Atom feeds
`,
	Run: func(cmd *cobra.Command, args []string) {
		feedName, articlesCount, err := internal.AddFeed(args[0])

		if err != nil {
			pterm.Error.Printf("Error adding feed: %v\n", err)
			return
		}

		pterm.Success.Printf("Added new feed: \"%s\" and imported %d items\n", feedName, articlesCount)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
