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
	Use:     "mount",
	Aliases: []string{"m"},
	Short:   "Mount specific bind entries",
	Long: `Mount specific bind entries.
Please specify "bind" entries defined in your configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		group := viper.GetBool("mount.group")
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			log.Fatal(err)
		}
		if err := runMount(args, group); err != nil {
			log.Fatal(err)
		}
	},
}

func runMount(names []string, group bool) error {
	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return err
	}
	var bs []*config.BindEntry
	if group {
		bs, err = cfg.GroupedBindEntries(names)
	} else {
		bs, err = cfg.BindEntries(names)
	}
	if err != nil {
		return err
	}
	ctl, err := controller.FromEntries(bs, true)
	if err != nil {
		return err
	}
	return ctl.Mount()
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
	mountCmd.Flags().BoolP("group", "g", false, "Mount groups")
	viper.BindPFlag("mount.group", mountCmd.Flags().Lookup("group"))
}
