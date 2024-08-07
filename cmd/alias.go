package cmd

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "alias is a command that will return the IP addresses of a website, the local IP address, and the DNS server address.",
	Long:  `alias is a command that will return the IP addresses of a website, the local IP address, and the DNS server address.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Usage: alias <destination>")
			os.Exit(1)
		}
		website := args[0]
		ips, err := net.LookupIP(website)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
			os.Exit(1)
		}

		// Get DNS server using ipconfig /all
		cmdOutput := &bytes.Buffer{}
		command := exec.Command("ipconfig", "/all")
		command.Stdout = cmdOutput
		err = command.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not get the IP: %v\n", err)
			os.Exit(1)
		}

		re := regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)
		output := cmdOutput.String()
		lines := strings.Split(output, "\n")
		var dnsServer string
		for _, line := range lines {
			if match := re.FindString(line); match != "" && len(line) > len("DNS Servers") {
				dnsServer = match
				break

			}
		}

		fmt.Println()

		if dnsServer != "" {
			fmt.Println("DNS Server: ", dnsServer)
		} else {
			fmt.Println("DNS Server not found")
		}

		fmt.Println()

		fmt.Println("Non-authoritative answer:")
		fmt.Println()
		fmt.Println("Name:", website)
		fmt.Println("Addresses:", ips[0])
		fmt.Println("\t  ", ips[1])
	},
}

func init() {
	rootCmd.AddCommand(aliasCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aliasCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aliasCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
