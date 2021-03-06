// Copyright © 2016 Kevin Kirsche <kev.kirsche@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// subdomainsCmd represents the subdomains command
var subdomainsCmd = &cobra.Command{
	Use:   "subdomains",
	Short: "Print subdomains and domain siblings",
	Long: `Prints the domain siblings or subdomain of all requested domains.

virustotal domain subdomain (-g) -a {{ api_key }} -d {{ domains }}`,
	Run: func(cmd *cobra.Command, args []string) {
		responses := retrieveDomainInformation()

		for _, resp := range responses {
			if !grepable {
				printStringSlice("Domain Siblings", resp.DomainSiblings)
				printStringSlice("Subdomains", resp.Subdomains)
			} else {
				for _, domain := range resp.DomainSiblings {
					if domain == "" {
						continue
					}
					fmt.Println(domain)
				}

				for _, domain := range resp.Subdomains {
					if domain == "" {
						continue
					}
					fmt.Println(domain)
				}
			}
		}
	},
}

func init() {
	domainCmd.AddCommand(subdomainsCmd)

	subdomainsCmd.PersistentFlags().BoolVarP(&grepable, "grep", "g", false, "Make the output grepable")
}
