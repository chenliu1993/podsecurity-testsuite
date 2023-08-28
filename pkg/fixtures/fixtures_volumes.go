package fixtures

import (
	corev1 "k8s.io/api/core/v1"
)

func fixturesVolumesBaseline() Case {
	return Case{
		Name: "volumes-baseline",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.Volumes(VolumesBaselinePass),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.Volumes(VolumesBaselineFail),
			}),
		},
	}
}

func fixturesVolumesRestricted() Case {
	return Case{
		Name: "volumes-restricted",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.Volumes(VolumesRestrictedPass),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.Volumes(VolumesRestrictedFail),
			}),
		},
	}
}

func init() {
	fixturesMap[volumesBaseline] = []func() Case{fixturesVolumesBaseline}
	fixturesMap[volumesRestricted] = []func() Case{fixturesVolumesRestricted}
}
