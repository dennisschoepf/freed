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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Lists all currently stored feeds",
	Long:    `This command lists all currently stored feeds with some metadata and their IDs. In the future it is possible to manipulate these feeds by referencing their ID.`,
	Example: `freed list
freed ls`,
	Run: func(cmd *cobra.Command, args []string) {
		tableData, err := internal.GetAllFeedsAsTable()

		if err != nil {
			pterm.Error.Printf("Error listing feeds: %v\n", err)
			return
		}

		pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
