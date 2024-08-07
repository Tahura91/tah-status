package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "http-status-checker [url] [pings]",
	Short: "hsc is a CLI tool to check the status of a website",
	Long:  `hsc is a CLI tool to check the status of a website. It can be used to check the status of a website and get the response time of the website.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	// Override the default help function
	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println("hsc is a CLI tool to check the status of a website.")
		fmt.Println("It can be used to check the status of a website and get the response time of the website.")
		fmt.Println("\nUsage:")
		fmt.Println("  http-status-checker [url] [pings]     Check the status and response time of the specified URL with specified pings.")
		fmt.Println("  http-status-checker [url]             Check the status of the specified URL.")
		fmt.Println("  http-status-checker alias [url]       Resolve and display the alias or canonical name of the specified URL.")
		fmt.Println("  http-status-checker ip                Retrieve and display the IP address of the Local Machine.")
		fmt.Println("  http-status-checker ip [url]          Retrieve and display the IP address of the specified URL")
		fmt.Println("  http-status-checker route [url]       Trace and display the route to the specified URL.")
		fmt.Println("\nFlags:")
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			fmt.Printf("  -%s, --%s: %s\n", flag.Shorthand, flag.Name, flag.Usage)
		})
	})

}
