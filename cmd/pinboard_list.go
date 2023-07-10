/*
Copyright © 2023 Mike Pountney <mike.pountney@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/imwally/pinboard"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List pinboard links",
	Long:  `List pinboard links`,
	Run:   runPinboardList,
}

func init() {
	pinboardCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runPinboardList(cmd *cobra.Command, args []string) {
	// Replace with your Pinboard API token
	apiToken := viper.Get("pinboard_api_token")
	if apiToken == nil {
		fmt.Printf("Must set MK_PINBOARD_API_TOKEN\n")
		return
	}

	pinboard.SetToken(apiToken.(string))

	// Retrieve all bookmarks
	bookmarks, err := pinboard.PostsRecent(nil)
	if err != nil {
		fmt.Printf("Error retrieving bookmarks: %v\n", err)
		return
	}

	// Print the retrieved bookmarks
	for _, bookmark := range bookmarks {
		fmt.Printf("URL: %s\n", bookmark.Href)
		fmt.Printf("Description: %s\n", bookmark.Description)
		fmt.Printf("Tags: %s\n", bookmark.Tags)
		fmt.Printf("Time: %s\n\n", bookmark.Time)
	}
}
