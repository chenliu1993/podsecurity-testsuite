package fixtures

import (
	"github.com/chenliu1993/podsecurity-check/pkg/constants"
	"github.com/chenliu1993/podsecurity-check/pkg/generator"
	corev1 "k8s.io/api/core/v1"
)

func fixturesRunAsUser() Case {
	g.SetPod(generator.GetBasePod().DeepCopy())
	return Case{
		Name: "runAsUser",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.RunAsUser(constants.CONTAINER, int64(1000)),
				g.RunAsUser(constants.INITCONTAINER, int64(1000)),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.RunAsUser(constants.CONTAINER, int64(0)),
				g.RunAsUser(constants.INITCONTAINER, int64(0)),
			}),
		},
	}
}

func fixturesEphemeralRunAsUser() Case {
	g.SetPod(generator.GetBasePod().DeepCopy())
	return Case{
		Name: "ephemeral-runAsUser",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.RunAsUser(constants.EPHEMERALCONTAINER, int64(1000)),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.RunAsUser(constants.EPHEMERALCONTAINER, int64(0)),
			}),
		},
	}
}

func init() {
	fixturesMap[runAsUser] = []func() Case{fixturesRunAsUser}
	fixturesMap[ephemeralrunAsUser] = []func() Case{fixturesEphemeralRunAsUser}
}
