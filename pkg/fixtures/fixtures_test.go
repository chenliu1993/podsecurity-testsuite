package fixtures

import (
	"os"
	"testing"

	"github.com/chenliu1993/podsecurity-check/pkg/cli"
	"github.com/chenliu1993/podsecurity-check/pkg/files"
)

func TestFixtures(t *testing.T) {
	// Get server version
	version, err := getServerVersion()
	if err != nil {
		t.Fatal(err)
	}
	if !older(version, serverVersion) {
		namespaces, err := GetNamespaceWithSecurityLabels()
		if err != nil {
			t.Fatal(err)
		}
		err = GenerateCases(namespaces)
		if err != nil {
			t.Fatal(err)
		}
		tests, err := files.WalkPath(testdataDir)
		if err != nil {
			t.Fatal(err)
		}
		errs := []error{}
		for _, tname := range tests {
			// expectedResult := getExpectedResult(filepath.Base(tname))
			content, err := os.OpenFile(tname, os.O_RDONLY, 0400)
			if err != nil {
				errs = append(errs, err)
			}
			_, err = cli.Kubectl(content, "apply", "-f", "-", "--dry-run=server")
			if err != nil {
				errs = append(errs, err)
			}
			// t.Fatal(res)
		}

		// files.Cleanup(testdataDir)
	}

}
