package fixtures

import (
	"github.com/chenliu1993/podsecurity-check/pkg/generator"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/pointer"
)

func fixturesContainerNonRoot() Case {
	return Case{
		Name: "nonRoot",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.RunAsNonRoot = pointer.Bool(true)
					pod.Template.Spec.Containers[0].SecurityContext.RunAsNonRoot = nil
					pod.Template.Spec.InitContainers[0].SecurityContext.RunAsNonRoot = nil
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.RunAsNonRoot = nil
					pod.Template.Spec.Containers[0].SecurityContext.RunAsNonRoot = pointer.Bool(true)
					pod.Template.Spec.InitContainers[0].SecurityContext.RunAsNonRoot = pointer.Bool(true)
					return pod
				},
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.RunAsNonRoot = nil
					pod.Template.Spec.Containers[0].SecurityContext.RunAsNonRoot = nil
					pod.Template.Spec.InitContainers[0].SecurityContext.RunAsNonRoot = nil
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.RunAsNonRoot = pointer.Bool(false)
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Containers[0].SecurityContext.RunAsNonRoot = pointer.Bool(false)
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.InitContainers[0].SecurityContext.RunAsNonRoot = pointer.Bool(false)
					return pod
				},
			}),
		},
	}
}

func init() {
	fixturesMap[runAsNonRoot] = []func() Case{fixturesContainerNonRoot}
}
