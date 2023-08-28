package fixtures

import (
	corev1 "k8s.io/api/core/v1"
)

func fixturesSysctls() Case {
	return Case{
		Name: "sysctls",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.Sysctls(SysctlsBaselinePass),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.Sysctls(SysctlsBaselineFail),
			}),
		},
	}
}

func init() {
	fixturesMap[sysctls] = []func() Case{fixturesSysctls}
}
