package main

import (
	"os"

	"cloudfreexiao/ant-graphql/backend-go/cmd"
	_ "cloudfreexiao/ant-graphql/backend-go/dao"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
