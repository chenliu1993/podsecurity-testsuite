package fixtures

import (
	corev1 "k8s.io/api/core/v1"
)

func fixturesApparmor() Case {
	return Case{
		Name: "apparmor",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.AppArmorProfile("runtime/default"),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.AppArmorProfile("local/bar"),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.AppArmorProfile("runtime/other"),
			}),
		},
	}
}

func init() {
	fixturesMap["apparmor"] = []func() Case{fixturesApparmor}
}
