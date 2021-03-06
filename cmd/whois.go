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
	"strings"

	"github.com/spf13/cobra"
)

// whoisCmd represents the whois command
var whoisCmd = &cobra.Command{
	Use:   "whois",
	Short: "Retrieve whois information",
	Long: `Retrieve whois information for domains.

virustotal domain whois (-g) -a {{ api_key }} -d {{ domains }}
`,
	Run: func(cmd *cobra.Command, args []string) {
		responses := retrieveDomainInformation()

		for _, resp := range responses {
			whoisLines := strings.Split(resp.Whois, "\n")
			if !grepable {
				printStringSlice("Whois", whoisLines)
			} else {
				for _, line := range whoisLines {
					fmt.Println(line)
				}
			}
		}
	},
}

func init() {
	domainCmd.AddCommand(whoisCmd)

	whoisCmd.Flags().BoolVarP(&grepable, "grep", "g", false, "Make the output grepable")
}
