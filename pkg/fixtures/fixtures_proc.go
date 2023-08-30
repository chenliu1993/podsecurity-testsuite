package fixtures

import (
	"github.com/chenliu1993/podsecurity-check/pkg/generator"
	corev1 "k8s.io/api/core/v1"
)

func fixturesProcMount() Case {
	return Case{
		Name:   "procMount",
		ErrMsg: "procMount",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					defaultProcMountType := corev1.DefaultProcMount
					pod.Template.Spec.Containers[0].SecurityContext.ProcMount = &defaultProcMountType
					pod.Template.Spec.InitContainers[0].SecurityContext.ProcMount = &defaultProcMountType
					return pod
				},
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					unmaskedProcMountType := corev1.UnmaskedProcMount
					pod.Template.Spec.Containers[0].SecurityContext.ProcMount = &unmaskedProcMountType
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					unmaskedProcMountType := corev1.UnmaskedProcMount
					pod.Template.Spec.InitContainers[0].SecurityContext.ProcMount = &unmaskedProcMountType
					return pod
				},
			}),
		},
	}
}

func init() {
	fixturesMap[procMount] = []func() Case{fixturesProcMount}
}
