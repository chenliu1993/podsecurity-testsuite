package fixtures

import (
	"github.com/chenliu1993/podsecurity-check/pkg/constants"
	corev1 "k8s.io/api/core/v1"
)

func fixturesCapabilitiesBaseline() Case {
	return Case{
		Name: "capabilities-baseline",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.CapabilitiesAdd(constants.CONTAINER, AddCapBaselinPassSet),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.CapabilitiesAdd(constants.INITCONTAINER, AddCapBaselinPassSet),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.CapabilitiesAdd(constants.CONTAINER, AddCapBaselinFailSet),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.CapabilitiesAdd(constants.INITCONTAINER, AddCapBaselinFailSet),
			}),
		},
	}
}
func fixturesCapabilitiesRestricted() Case {
	return Case{
		Name: "capabilities-restricted",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.CapabilitiesAdd(constants.CONTAINER, AddCapRestrictedPassSet),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.CapabilitiesAdd(constants.INITCONTAINER, DropCapRestrictedPassSet),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.CapabilitiesAdd(constants.CONTAINER, AddCapRestrictedFailSet),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.CapabilitiesAdd(constants.INITCONTAINER, DropCapRestrictedFailSet),
			}),
		},
	}
}

func fixturesEphemeralCapabilitesBaseline() Case {
	return Case{
		Name: "ephemeral-capabilities-baseline",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.CapabilitiesAdd(constants.EPHEMERALCONTAINER, AddCapBaselinPassSet),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.CapabilitiesAdd(constants.EPHEMERALCONTAINER, AddCapBaselinFailSet),
			}),
		},
	}
}

func fixturesEphemeralCapabilitesRestricted() Case {
	return Case{
		Name: "ephemeral-capabilities-restricted",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.CapabilitiesAdd(constants.EPHEMERALCONTAINER, AddCapRestrictedPassSet),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.CapabilitiesAdd(constants.EPHEMERALCONTAINER, DropCapRestrictedPassSet),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.CapabilitiesAdd(constants.EPHEMERALCONTAINER, AddCapRestrictedFailSet),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.CapabilitiesAdd(constants.EPHEMERALCONTAINER, DropCapRestrictedFailSet),
			}),
		},
	}
}

func init() {
	fixturesMap["capabilities"] = []func() Case{
		fixturesCapabilitiesBaseline,
		fixturesCapabilitiesRestricted,
	}
	fixturesMap["ephemeral-capabilities"] = []func() Case{
		fixturesEphemeralCapabilitesBaseline,
		fixturesEphemeralCapabilitesRestricted,
	}
}
