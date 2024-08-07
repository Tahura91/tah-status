/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
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

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "ip command will give the local ip address",
	Long:  `ip command will give the local ip address`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmdOutput := &bytes.Buffer{}
			command := exec.Command("ipconfig")
			command.Stdout = cmdOutput
			err := command.Run()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not get the IP: %v\n", err)
				os.Exit(1)
			}

			re := regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)
			output := cmdOutput.String()
			lines := strings.Split(output, "\n")
			var ipv4 string
			for _, line := range lines {
				if match := re.FindString(line); match != "" && len(line) > len("IPv4 Address") {
					ipv4 = match
					break

				}
			}
			fmt.Println()
			fmt.Println("Local IPv4 Address : ", ipv4)
			fmt.Println()

		} else if len(args) == 1 {
			website := args[0]
			ips, err := net.LookupIP(website)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
				os.Exit(1)
			}

			fmt.Println()
			fmt.Println("IP addresses of ", website)
			fmt.Println("  IPv6 Address: " + ips[0].String())
			fmt.Println("  IPv4 Address: " + ips[1].String())
			fmt.Println()
		} else {
			fmt.Println()
			fmt.Println("Usage:")
			fmt.Println("\t ip")
			fmt.Println("\t ip <destination>")
			fmt.Println()
		}

	},
}

func init() {
	rootCmd.AddCommand(ipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
