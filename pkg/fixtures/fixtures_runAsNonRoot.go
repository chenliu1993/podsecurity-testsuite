package fixtures

import (
	"github.com/chenliu1993/podsecurity-check/pkg/constants"
	corev1 "k8s.io/api/core/v1"
)

func fixturesContainerNonRoot() Case {
	return Case{
		Name: "nonRoot",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.NonRoot(constants.CONTAINER, false),
				g.NonRoot(constants.INITCONTAINER, false),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.NonRoot(constants.CONTAINER, true),
				g.NonRoot(constants.INITCONTAINER, true),
			}),
		},
	}
}

func fixturesEphemeralNonRoot() Case {
	return Case{
		Name: "ephemeral-nonRoot",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.NonRoot(constants.EPHEMERALCONTAINER, false),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.NonRoot(constants.EPHEMERALCONTAINER, true),
			}),
		},
	}
}

func init() {
	fixturesMap[runAsNonRoot] = []func() Case{fixturesContainerNonRoot}
	fixturesMap[ephemeralrunAsNonRoot] = []func() Case{fixturesEphemeralNonRoot}
}
