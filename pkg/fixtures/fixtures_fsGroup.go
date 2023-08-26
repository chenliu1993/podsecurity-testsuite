package fixtures

import corev1 "k8s.io/api/core/v1"

func fixturesFsGroup() Case {
	return Case{
		Name: "fsGroup",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.FsGroup(FsGroupBaselinePass),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.FsGroup(FsGroupBaselineFail),
			}),
		},
	}
}

func init() {
	fixturesMap["fsgroup"] = []func() Case{fixturesFsGroup}
}
