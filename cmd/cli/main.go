package main

import (
	"context"
	"fmt"

	pb "github.com/Zzocker/multicast/protos"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var rootCMD = &cobra.Command{
	Use:   "multicast-cli",
	Short: "multicast cli tool for interacting with multicast simulator server",
}

var endpoint string

func main() {
	if err := rootCMD.Execute(); err != nil {
		panic(err)
	}
}

func init() {

	rootCMD.PersistentFlags().StringVar(&endpoint, "endpoint", "localhost:9090", "multicast server endpoint")

	setCMD.Flags().Int64P("value", "v", 0, "value to set")

	rootCMD.AddCommand(getCMD)
	rootCMD.AddCommand(setCMD)

}

var getCMD = &cobra.Command{
	Use:   "get",
	Short: "make get call to multicast server",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
		if err != nil {
			return err
		}
		client := pb.NewMulticastClient(conn)
		data, err := client.Get(context.Background(), &pb.Empty{})
		if err != nil {
			return err
		}
		fmt.Printf("Value from multicast simulator = %d\n", data.Value)
		return nil
	},
}

var setCMD = &cobra.Command{
	Use:   "set",
	Short: "make set call to multicast server",
	RunE: func(cmd *cobra.Command, args []string) error {
		value, _ := cmd.Flags().GetInt64("value")
		conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
		if err != nil {
			return err
		}
		client := pb.NewMulticastClient(conn)
		_, err = client.Set(context.Background(), &pb.Data{Value: value})
		if err != nil {
			return err
		}
		fmt.Printf("Value = %d has been set to multicast server\n", value)
		return nil
	},
}
