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
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/hiroara/dirlink/config"
	"github.com/hiroara/dirlink/controller"
)

// mountCmd represents the mount command
var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "Mount specific bind entries",
	Long: `Mount specific bind entries.
Please specify "bind" entries defined in your configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			log.Fatal(err)
		}
		if err := runMount(args); err != nil {
			log.Fatal(err)
		}
	},
}

func runMount(names []string) error {
	var cfg config.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}
	bs, err := cfg.BindEntries(names)
	if err != nil {
		return err
	}
	controller.FromEntries(bs, true).Mount()
	return nil
}

func init() {
	rootCmd.AddCommand(mountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
