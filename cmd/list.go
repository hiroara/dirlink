/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/hiroara/dirlink/config"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list bind entries",
	Long:  "list bind entries defined in your configuration.",
	Run: func(cmd *cobra.Command, args []string) {
		verbose := viper.GetBool("verbose")
		group := viper.GetBool("group")
		if err := runList(group, verbose); err != nil {
			log.Fatal(err)
		}
	},
}

func runList(group, verbose bool) error {
	var cfg config.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}
	var keys []string
	if group {
		keys = cfg.GroupKeys()
	} else {
		keys = cfg.BindKeys()
	}
	for _, key := range keys {
		fmt.Printf("%s\n", key)
		if verbose {
			entry := cfg.Bind[key]
			for _, link := range entry.Links {
				fmt.Printf("  %s -> %s\n", entry.Src, link)
			}
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	viper.BindPFlag("verbose", listCmd.Flags().Lookup("verbose"))

	listCmd.Flags().BoolP("group", "g", false, "Show list of groups")
	viper.BindPFlag("group", listCmd.Flags().Lookup("group"))
}
