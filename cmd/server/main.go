package main

import (
	"math/rand"
	"time"

	"github.com/Zzocker/multicast/algo"
	"github.com/Zzocker/multicast/server"
	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:   "multicast",
	Short: "start multicast simulator server",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		protocol, _ := cmd.Flags().GetString("algo")
		count, _ := cmd.Flags().GetInt("count")

		server.Run(port, algo.ALGO_NAME(protocol), count)
	},
}

func init() {
	rootCMD.Flags().IntP("port", "p", 9090, "multicast server running port")
	rootCMD.Flags().String("algo", "anti_entropy", "protocol to simulate")
	rootCMD.Flags().IntP("count", "c", 8, "peer count")
}
func main() {

	rand.Seed(time.Now().Unix())
	if err := rootCMD.Execute(); err != nil {
		panic(err)
	}
}
