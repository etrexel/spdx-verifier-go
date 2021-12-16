package utils

import (
	"os"
	"strings"

	"github.com/etrexel/spdx-verifier-go/pkg/types"
	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/tvloader"
)

func LoadSPDXDoc(filename string) (*spdx.Document2_2, error) {
	file, err := os.Open(filename)
	if err != nil {
		return &spdx.Document2_2{}, err
	}

	return tvloader.Load2_2(file)
}

func GetPackageLicenses(doc spdx.Document2_2) []types.PackageLicense {
	var output []types.PackageLicense
	for _, p := range doc.Packages {
		output = append(output, types.PackageLicense{
			PackageName:    p.PackageName,
			PackageLicense: p.PackageLicenseConcluded,
		})
	}
	return output
}

func LoadAllowedLicenses(filename string) (map[string]struct{}, error) {
	output := map[string]struct{}{}
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	values := strings.Split(string(data), "\n")
	for _, val := range values {
		output[val] = struct{}{}
	}
	return output, nil
}

func FilterPackageLicenses(packageLicenses []types.PackageLicense, allowedLicenses map[string]struct{}) []types.PackageLicense {
	var output []types.PackageLicense
	for _, packageLicense := range packageLicenses {
		_, ok := allowedLicenses[packageLicense.PackageLicense]
		if !ok {
			output = append(output, packageLicense)
		}
	}
	return output
}
