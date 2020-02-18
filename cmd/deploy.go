package cmd

import (
	"fmt"
	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"os"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
)

var (
)

func DeployCmd(
	ctx *server.Context,
	cdc *codec.Codec,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy files for a Konstellation network",
		RunE: func(_ *cobra.Command, _ []string) error {
			// Use SSH key authentication from the auth package
			// we ignore the host key in this example, please change this if you use this library
			clientConfig, err := auth.PrivateKey("username", "/home/kirillbeldyaga/Documents/.ssh/id_rsa", )
			if err != nil {
				return err
			}

			// For other authentication methods see ssh.ClientConfig and ssh.AuthMethod

			// Create a new SCP client
			client := scp.NewClient("95.216.212.119:22", &clientConfig)

			// Connect to the remote server
			if err := client.Connect(); err != nil {
				fmt.Println("Couldn't establish a connection to the remote server ", err)
				return err
			}

			// Open a file
			f, _ := os.Open("./Makefile")

			// Close client connection after the file has been copied
			defer client.Close()

			// Close the file after it has been copied
			defer f.Close()

			// Finaly, copy the file over
			// Usage: CopyFile(fileReader, remotePath, permission)

			err = client.CopyFile(f, "/root/", "0655")

			if err != nil {
				fmt.Println("Error while copying file ", err)
			}
			return err
		},
	}

	return cmd
}