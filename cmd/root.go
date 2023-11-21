/*
Copyright © 2023 (c) Microsoft Corporation

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
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	hyperkv "github.com/Azure/hyperkv"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hypverkv",
	Short: "kv parser for files maintained by hyper-v integration service",
	Long: `A parser for files maintained by hyper-v integration service.
	
	The parser will parse data from Hyper-V data exchange service.
	For more information, please refer to https://learn.microsoft.com/en-us/previous-versions/windows/it-pro/windows-server-2012-R2-and-2012/dn798287(v=ws.11)#linux-guests`,
	Run: Run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&filePath, "file", "f", "/var/lib/hyperv/.kvp_pool_3", "Path to the file to be parsed")
	rootCmd.Flags().StringVarP(&outputFormat, "output", "o", "plain", "Output format, supported json, plain")
}

var (
	filePath     string
	outputFormat string
)

func Run(cmd *cobra.Command, args []string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
	}
	list := hyperkv.Parse(content)
	if outputFormat == "json" {
		content, err := json.Marshal(list)
		if err != nil {
			fmt.Printf("Error marshalling json: %v", err)
		}
		fmt.Printf("%s\n", content)
		return
	} else {
		fmt.Printf("Parsing file %s, len %d\n", filePath, len(content))
		for _, item := range list {
			fmt.Printf("%s=%s\n", item.Key, item.Value)
		}
	}
}
