package fixtures

import corev1 "k8s.io/api/core/v1"

func fixturesSupplementalGroups() Case {
	return Case{
		Name: "supplementalGroups",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.SupplementalGroups(SupplementalGroupsBaselinePass),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.SupplementalGroups(SupplementalGroupsBaselineFail),
			}),
		},
	}
}

func init() {
	fixturesMap[supplementalgroups] = []func() Case{fixturesSupplementalGroups}
}
