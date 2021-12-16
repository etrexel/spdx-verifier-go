package main

import (
	"fmt"

	"github.com/etrexel/spdx-verifier-go/pkg/utils"
	"github.com/spf13/cobra"
)

type VerifyConfigOptions struct {
	AllowedLicensesFile string
}

var VerifyConfig = &VerifyConfigOptions{}

func verifyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "verify -l allowed_licenses.txt sbom.spdx",
		Aliases: []string{"v"},
		Short:   "Verify licenses from an SPDX file",
		RunE:    verifyRun,
	}

	cmd.Flags().StringVarP(&VerifyConfig.AllowedLicensesFile, "licenses", "l", "allowed_licenses.txt", "File of licenses that are allowed")

	return cmd
}

func verifyRun(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("no input SPDX file provided")
	}

	doc, err := utils.LoadSPDXDoc(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Document: %s\n", doc.CreationInfo.DocumentName)

	packageLicenses := utils.GetPackageLicenses(*doc)
	allowedLicenses, err := utils.LoadAllowedLicenses(VerifyConfig.AllowedLicensesFile)
	if err != nil {
		return err
	}
	disallowedPackages := utils.FilterPackageLicenses(packageLicenses, allowedLicenses)

	if len(disallowedPackages) > 0 {
		fmt.Printf("VERIFICATION FAILED: some packages contain a license that is not allowed\n")
		fmt.Printf("Problem packages:\n")
		for _, packageLicense := range disallowedPackages {
			fmt.Printf("  Package: %30s, License: %30s\n", packageLicense.PackageName, packageLicense.PackageLicense)
		}
	} else {
		fmt.Printf("VERIFICATION SUCCESS: all packages have acceptable licenses\n")
	}

	return nil
}
