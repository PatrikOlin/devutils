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
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PatrikOlin/devutils/utils"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var baseUrl = "https://api.fejk.company/v1/"
// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate fake data",
	Long: `Generate top tier fake identities and company data. 
		NOTE: only to be used for mischief.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a type argument")
		}
		if utils.IsValidType(args[0]) {
			return nil
		}
		return fmt.Errorf("invalid type specified: %v", args[0])
	},

	Run: func(cmd *cobra.Command, args []string) {
		var res string

		a, _ := cmd.Flags().GetInt("multiple")

		switch args[0] {
		case "person":
			res = getPeople(a)
		case "company":
			res = getCompanies(a)
		case "article":
			res = getArticles(a)
		}

		c, _ := cmd.Flags().GetBool("clipboard")


		if (a == 1) {
			res = strings.Trim(res, "[")
			res = strings.Trim(res, "]")
		}

		if (c == true) {
			clipboard.WriteAll(res)
			fmt.Println("Entity saved to clipboard")
		} else {
			fmt.Println(res)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// generateCmd.Flags().StringP("company", "c", "", "Returns a randomly generated company")
	// generateCmd.Flags().StringP("person", "p", "", "Returns a randomly generated person")
	// generateCmd.Flags().StringP("article", "a", "", "Returns a randomly generated article")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	generateCmd.Flags().BoolP("clipboard", "c", false, "Copy result to clipboard")
	generateCmd.Flags().IntP("multiple", "m", 1, "Get multiple generated entitys (up to 10)")
}



func getPeople(n int) string {
	url := fmt.Sprintf("%speople?amount=%d", baseUrl, n)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	return bodyString
}

func getCompanies(n int) string {
	url := fmt.Sprintf("%scompanies?amount=%d", baseUrl, n)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)

	return bodyString
}

func getArticles(n int) string {
	url := fmt.Sprintf("%sarticles?amount=%d", baseUrl, n)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	return bodyString
}
