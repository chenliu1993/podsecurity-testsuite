package main

import (
	"fmt"

	"github.com/chenliu1993/podsecurity-check/pkg/fixtures"
)

func main() {
	labels, err := fixtures.GetNamespaceWithSecurityLabels()
	if err != nil {
		panic(err)
	}
	fmt.Println(labels)
}
