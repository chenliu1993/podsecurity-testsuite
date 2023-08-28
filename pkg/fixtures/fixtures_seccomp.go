package fixtures

import (
	"github.com/chenliu1993/podsecurity-check/pkg/constants"
	corev1 "k8s.io/api/core/v1"
)

func fixturesSeccompBaseline() Case {
	return Case{
		Name: "seccomp",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.Seccomp(constants.POD, &DefaultSeccompProfile),
				g.Seccomp(constants.CONTAINER, &DefaultSeccompProfile),
				g.Seccomp(constants.INITCONTAINER, &DefaultSeccompProfile),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.Seccomp(constants.POD, &LocalhostProfile),
				g.Seccomp(constants.CONTAINER, &LocalhostProfile),
				g.Seccomp(constants.INITCONTAINER, &LocalhostProfile),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.Seccomp(constants.POD, &UnconfinedSeccompProfile),
				g.Seccomp(constants.CONTAINER, &UnconfinedSeccompProfile),
				g.Seccomp(constants.INITCONTAINER, &UnconfinedSeccompProfile),
			}),
		},
	}
}

func fixturesEphemeralSeccompBaseline() Case {
	return Case{
		Name: "ephemeral-seccomp",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.Seccomp(constants.EPHEMERALCONTAINER, &DefaultSeccompProfile),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.Seccomp(constants.EPHEMERALCONTAINER, &UnconfinedSeccompProfile),
			}),
		},
	}
}

func fixturesSeccompRestricted() Case {
	return Case{
		Name: "Seccomp",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.Seccomp(constants.POD, &DefaultSeccompProfile),
				g.Seccomp(constants.CONTAINER, nil),
				g.Seccomp(constants.INITCONTAINER, nil),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.Seccomp(constants.POD, nil),
				g.Seccomp(constants.CONTAINER, &DefaultSeccompProfile),
				g.Seccomp(constants.INITCONTAINER, &DefaultSeccompProfile),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.Seccomp(constants.POD, nil),
				g.Seccomp(constants.CONTAINER, &LocalhostProfile),
				g.Seccomp(constants.INITCONTAINER, &LocalhostProfile),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.Seccomp(constants.POD, &UnconfinedSeccompProfile),
				g.Seccomp(constants.CONTAINER, &UnconfinedSeccompProfile),
				g.Seccomp(constants.INITCONTAINER, &UnconfinedSeccompProfile),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.Seccomp(constants.POD, &DefaultSeccompProfile),
				g.Seccomp(constants.CONTAINER, &DefaultSeccompProfile),
				g.Seccomp(constants.INITCONTAINER, &DefaultSeccompProfile),
			}),
		},
	}
}

func fixturesEphemeralSeccompRestricted() Case {
	return Case{
		Name: "EphemeralSeccomp",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.Seccomp(constants.POD, &DefaultSeccompProfile),
				g.Seccomp(constants.EPHEMERALCONTAINER, nil),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.Seccomp(constants.POD, nil),
				g.Seccomp(constants.EPHEMERALCONTAINER, &DefaultSeccompProfile),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.Seccomp(constants.POD, nil),
				g.Seccomp(constants.EPHEMERALCONTAINER, &LocalhostProfile),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.Seccomp(constants.POD, &UnconfinedSeccompProfile),
				g.Seccomp(constants.EPHEMERALCONTAINER, &UnconfinedSeccompProfile),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.Seccomp(constants.POD, &DefaultSeccompProfile),
				g.Seccomp(constants.EPHEMERALCONTAINER, &DefaultSeccompProfile),
			}),
		},
	}
}

func init() {
	fixturesMap[seccomp] = []func() Case{fixturesSeccompBaseline}
	fixturesMap[ephemeralSeccomp] = []func() Case{fixturesEphemeralSeccompBaseline}
	fixturesMap[seccompRestricted] = []func() Case{fixturesSeccompRestricted}
	fixturesMap[ephemeralSeccompRestricted] = []func() Case{fixturesEphemeralSeccompRestricted}
}
