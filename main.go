package main

import (
	"fmt"

	"github.com/chenliu1993/podsecurity-check/pkg/files"
	"github.com/chenliu1993/podsecurity-check/pkg/fixtures"
)

func main() {
	err := fixtures.GenerateCases()
	if err != nil {
		panic(err)
	}
	file, err := files.WalkPath("/home/vagrant/go/src/github.com/chenliu1993/podsecurity-check/pkg/fixtures/testdata")
	if err != nil {
		panic(err)
	}
	fmt.Println(file)
}
