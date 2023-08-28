package fixtures

import (
	"github.com/chenliu1993/podsecurity-check/pkg/constants"
	corev1 "k8s.io/api/core/v1"
)

func fixturesProcMount() Case {
	return Case{
		Name: "procMount",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.ProcMount(constants.CONTAINER, &DefaultProcMount),
				g.ProcMount(constants.INITCONTAINER, &DefaultProcMount),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.ProcMount(constants.CONTAINER, &UmaskedProcMount),
				g.ProcMount(constants.INITCONTAINER, &UmaskedProcMount),
			}),
		},
	}
}

func fixturesEphemeralProcMount() Case {
	return Case{
		Name: "ephemeral-procMount",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.ProcMount(constants.EPHEMERALCONTAINER, &DefaultProcMount),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.ProcMount(constants.EPHEMERALCONTAINER, &UmaskedProcMount),
			}),
		},
	}
}

func init() {
	fixturesMap[procMount] = []func() Case{fixturesProcMount}
	fixturesMap[ephemeralprocMount] = []func() Case{fixturesEphemeralProcMount}
}
