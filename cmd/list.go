package main

import (
	"fmt"

	"github.com/etrexel/spdx-verifier-go/pkg/utils"
	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list sbom.spdx",
		Aliases: []string{"l"},
		Short:   "List licenses from an SPDX file",
		RunE:    listRun,
	}

	return cmd
}

func listRun(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("no input SPDX file provided")
	}

	doc, err := utils.LoadSPDXDoc(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Document: %s\n", doc.CreationInfo.DocumentName)

	packageLicenses := utils.GetPackageLicenses(*doc)
	for _, packageLicense := range packageLicenses {
		fmt.Printf("  Package: %30s, License: %30s\n", packageLicense.PackageName, packageLicense.PackageLicense)
	}

	return nil
}
