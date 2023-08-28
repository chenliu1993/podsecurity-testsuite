package main

import (
	"fmt"
	"regexp"
)

func main() {
	// WarningMsgRe := regexp.MustCompile("Warning: would violate PodSecurity.*")
	EnforceNewMsgRe := regexp.MustCompile("error when creating.*: pods.* is forbidden: violates PodSecurity")
	// EnforceExistingMsgRe := regexp.MustCompile("Warning: existing pods in namespace.*violate the new PodSecurity enforce level")
	// fmt.Println(WarningMsgRe.MatchString("Warning: would violate PodSecurity \"baseline:v1.27\": hostPort (container \"container1\" uses hostPort 12345)"))
	fmt.Println(EnforceNewMsgRe.MatchString("Error from server (Forbidden): error when creating \"hostports0.yaml\": pods \"hostports0\" is forbidden: violates PodSecurity \"baseline:latest\": hostPort (container \"container1\" uses hostPort 12345)"))
}
