package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spdx_verifier",
	Short: "Work with licenses in an SPDX file",
}

func init() {
	rootCmd.AddCommand(listCmd())
	rootCmd.AddCommand(verifyCmd())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
