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
	"bufio"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate the default configuration file in your home directory",
	Long: `Generate the default configuration file

The file that will be created:
- $HOME/.dirlink.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runInit(); err != nil {
			log.Fatal(err)
		}
	},
}

func runInit() error {
	tmpl, err := template.New("config").Parse(string(MustAsset("cmd/data/templates/.dirlink.yaml")))
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/.dirlink.yaml", os.Getenv("HOME"))
	if fileExists(path) {
		log.Printf("File already exists: %s\n", path)
		return nil
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	if err = tmpl.Execute(w, ""); err != nil {
		return err
	}
	if err = w.Flush(); err != nil {
		return err
	}
	log.Printf("File created at %s\n", path)
	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
