package fixtures

import (
	"github.com/chenliu1993/podsecurity-check/pkg/generator"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/pointer"
)

func fixturesRunAsUser() Case {
	return Case{
		Name: "runAsUser",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.RunAsUser = pointer.Int64(1001)
					pod.Template.Spec.Containers[0].SecurityContext.RunAsUser = pointer.Int64(1001)
					pod.Template.Spec.InitContainers[0].SecurityContext.RunAsUser = pointer.Int64(1001)
					return pod
				},
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.RunAsUser = pointer.Int64(0)
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Containers[0].SecurityContext.RunAsUser = pointer.Int64(0)
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.InitContainers[0].SecurityContext.RunAsUser = pointer.Int64(0)
					return pod
				},
			}),
		},
	}
}

func init() {
	fixturesMap[runAsUser] = []func() Case{fixturesRunAsUser}
}
