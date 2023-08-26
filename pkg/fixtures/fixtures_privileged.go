package fixtures

import (
	"github.com/chenliu1993/podsecurity-check/pkg/constants"
	corev1 "k8s.io/api/core/v1"
)

func fixturesPrivileged() Case {
	return Case{
		Name: "privileged",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.Privileged(false, constants.CONTAINER),
				g.Privileged(false, constants.INITCONTAINER),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.Privileged(true, constants.CONTAINER),
				g.Privileged(true, constants.INITCONTAINER),
			}),
		},
	}
}

func fixturesEphemeralPrivileged() Case {
	return Case{
		Name: "ephemeral-privileged",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.Privileged(false, constants.EPHEMERALCONTAINER),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.Privileged(true, constants.EPHEMERALCONTAINER),
			}),
		},
	}
}

func fixturesAllowPrivilegeEscalation() Case {
	return Case{
		Name: "allowPrivilegeEscalation",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.AllowPrivilegeEscalation(false, constants.CONTAINER),
				g.AllowPrivilegeEscalation(false, constants.INITCONTAINER),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.AllowPrivilegeEscalation(true, constants.CONTAINER),
				g.AllowPrivilegeEscalation(true, constants.INITCONTAINER),
			}),
		},
	}
}

func fixturesEphemeralAllowPrivilegeEscalation() Case {
	return Case{
		Name: "ephemeral-allowPrivilegeEscalation",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.AllowPrivilegeEscalation(false, constants.EPHEMERALCONTAINER),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.AllowPrivilegeEscalation(true, constants.EPHEMERALCONTAINER),
			}),
		},
	}
}

func init() {
	fixturesMap["privileged"] = []func() Case{fixturesPrivileged}
	fixturesMap["allowPrivilegeEscalation"] = []func() Case{fixturesAllowPrivilegeEscalation}
	fixturesMap["ephemeralPrivileged"] = []func() Case{fixturesEphemeralPrivileged}
	fixturesMap["ephemeralAllowPrivilegeEscalation"] = []func() Case{fixturesEphemeralAllowPrivilegeEscalation}
}
