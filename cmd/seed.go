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
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/carneles/goseeder/service"
	"github.com/spf13/cobra"
)

var source, database string

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		seeder := service.NewSeeder(database)

		var files []string
		err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				if filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml" {
					files = append(files, path)
				}
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			fileData, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}

			var seedData service.SeedData
			err = yaml.Unmarshal(fileData, &seedData)
			if err != nil {
				log.Fatal(err)
			}

			err = seeder.Seed(seedData)
			if err != nil {
				log.Panic("Cannot initialize data seeder.")
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)

	seedCmd.Flags().StringVarP(&source, "source", "s", "", "folder where the data is stored")
	seedCmd.MarkFlagRequired("source")
	seedCmd.Flags().StringVarP(&database, "database", "d", "", "database connection url")
	seedCmd.MarkFlagRequired("database")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// seedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// seedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
