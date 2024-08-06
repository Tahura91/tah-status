/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type check struct {
	ping   int
	Ip     string
	status string
}

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the status of a website",
	Long:  `Check the status of a website. It can be used to check the status of a website and get the response time of the website.`,
	Run: func(cmd *cobra.Command, args []string) {
		https := "https://"

		var httpStatus []check

		url := args[0]

		if !strings.HasPrefix(url, https) {
			url = https + url
		}
		table := tablewriter.NewWriter(os.Stdout)
		headers := []string{"Ping (ms)", "Address", "Status"}
		table.SetHeader(headers)

		if len(args) > 1 {
			pings, _ := strconv.Atoi(args[1])
			fmt.Println("Pinging to the Following URL", url)

			fmt.Println()
			fmt.Println("Check the following table for the status of the website")
			fmt.Println()

			if pings != 0 {

				for i := 0; i < pings; i++ {
					start := time.Now()
					resp, err := http.Get(url)
					if err != nil {
						fmt.Println("Error:", err)
					}
					ping := time.Since(start).Milliseconds()

					ips, _ := net.LookupIP(resp.Request.URL.Hostname())
					ip := ips[0].String()
					status := "❌"
					if resp.StatusCode == http.StatusOK {
						status = "✅"
					}

					value := check{int(ping), ip, status}

					httpStatus = append(httpStatus, value)

				}

				for i := 0; i < pings; i++ {
					table.Append([]string{strconv.Itoa(httpStatus[i].ping), httpStatus[i].Ip, httpStatus[i].status})
				}

			}
		} else {
			pings := 4
			fmt.Println("Pinging to the Following URL", url)

			fmt.Println()
			fmt.Println("Check the following table for the status of the website")
			fmt.Println()

			if pings != 0 {
				for i := 0; i < pings; i++ {
					start := time.Now()
					resp, err := http.Get(url)
					if err != nil {
						fmt.Println("Error:", err)
					}
					ping := time.Since(start).Milliseconds()

					ips, _ := net.LookupIP(resp.Request.URL.Hostname())
					ip := ips[0].String()
					status := "❌"
					if resp.StatusCode == http.StatusOK {
						status = "✅"
					}

					value := check{int(ping), ip, status}

					httpStatus = append(httpStatus, value)

				}

				for i := 0; i < pings; i++ {
					table.Append([]string{strconv.Itoa(httpStatus[i].ping), httpStatus[i].Ip, httpStatus[i].status})
				}
			}
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
