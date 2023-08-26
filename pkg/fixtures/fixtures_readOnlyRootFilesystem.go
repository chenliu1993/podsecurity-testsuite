package fixtures

import (
	"github.com/chenliu1993/podsecurity-check/pkg/constants"
	corev1 "k8s.io/api/core/v1"
)

func fixturesReadOnly() Case {
	return Case{
		Name: "readOnly",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.ReadOnly(false, constants.CONTAINER),
				g.ReadOnly(false, constants.INITCONTAINER),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.ReadOnly(true, constants.CONTAINER),
				g.ReadOnly(true, constants.INITCONTAINER),
			}),
		},
	}
}

func fixturesEphemeralReadOnly() Case {
	return Case{
		Name: "ephemeral-ReadOnly",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.ReadOnly(false, constants.EPHEMERALCONTAINER),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.ReadOnly(true, constants.EPHEMERALCONTAINER),
			}),
		},
	}
}

func init() {
	fixturesMap["readOnly"] = []func() Case{fixturesReadOnly}
	fixturesMap["ephemeralreadOnly"] = []func() Case{fixturesEphemeralReadOnly}
}
