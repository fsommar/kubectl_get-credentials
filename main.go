package main

import (
	"github.com/fsommar/kubectl-get_credentials/cmd"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"os"
)

const version = "v0.1.0"

func main() {
	getCredentialsCmd := cmd.NewGetCredentialsCommand(genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	})
	getCredentialsCmd.Version = version
	if err := getCredentialsCmd.Execute(); err != nil {
		os.Exit(1)
	}
}