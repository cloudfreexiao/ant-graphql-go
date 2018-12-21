package cmd

import (
	"github.com/spf13/cobra"

	"cloudfreexiao/ant-graphql/backend-go/server"
)

var (
	debug       bool
	disableAuth bool
	port        int
)

var RootCmd = &cobra.Command{
	Use:   "graphql-server",
	Short: "GraphQL API server in golang to get linux system info",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := server.NewServer(&server.Config{
			Debug:       debug,
			DisableAuth: disableAuth,
		})
		if err != nil {
			return err
		}
		s.Run(port)
		return nil
	},
}

func init() {
	RootCmd.Flags().BoolVarP(&debug, "debug", "d", true, "debug mode")
	RootCmd.Flags().BoolVarP(&disableAuth, "disable-auth", "", false, "disable auth middleware")
	RootCmd.Flags().IntVarP(&port, "port", "p", 9527, "port number")
}