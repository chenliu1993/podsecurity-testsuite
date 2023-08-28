package validate

import (
	"fmt"
	"regexp"
)

var (
	WarningMsgRe         = regexp.MustCompile("Warning: would violate PodSecurity.*")
	EnforceNewMsgRe      = regexp.MustCompile("error when creating.*: pods.* is forbidden: violates PodSecurity")
	EnforceExistingMsgRe = regexp.MustCompile("Warning: existing pods in namespace.*violate the new PodSecurity enforce level")
)

func CheckWarningMsg(msg string) error {
	var err error
	if WarningMsgRe.MatchString(msg) {
		err = fmt.Errorf("break warning mode: %s", msg)
	}
	return err
}

func CheckEnforceNewMsg(msg string) error {
	var err error
	if EnforceNewMsgRe.MatchString(msg) {
		err = fmt.Errorf("break enforce mode: %s", msg)
	}
	return err
}

func CheckEnforceExistingMsg(msg string) error {
	var err error
	if EnforceExistingMsgRe.MatchString(msg) {
		err = fmt.Errorf("break enforce mode on an existing namespace: %s", msg)
	}
	return err
}
