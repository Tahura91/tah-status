package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

// routeCmd represents the route command
var routeCmd = &cobra.Command{
	Use:   "route",
	Short: "route command will give the route form ur ip to the destination",
	Long:  `route command will give the route form ur ip to the destination`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Usage: route <destination>")
			os.Exit(1)
		}
		fmt.Printf("Route to %s\n", args[0])

		dest := args[0]
		ipAddr, err := net.ResolveIPAddr("ip4", dest)
		if err != nil {
			fmt.Printf("ResolveIPAddr error: %v\n", err)
			os.Exit(1)
		}

		conn, err := net.Dial("ip4:icmp", ipAddr.String())
		if err != nil {
			fmt.Printf("Dial error: %v\n", err)
			os.Exit(1)
		}
		defer conn.Close()

		fmt.Println()
		fmt.Println("Route completed.")
		fmt.Println()

	},
}

func init() {
	rootCmd.AddCommand(routeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// routeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// routeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
