package fixtures

import (
	"testing"

	"github.com/chenliu1993/podsecurity-check/pkg/files"
)

func TestFixtures(t *testing.T) {
	err := GenerateCases()
	if err != nil {
		t.Fatal(err)
	}
	tests, err := files.WalkPath(testdataDir)
	if err != nil {
		t.Fatal(err)
	}

	for _, tname := range tests {
		t.Log(tname)
	}

	// files.Cleanup(testdataDir)
}
