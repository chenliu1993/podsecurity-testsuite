package fixtures

import (
	"github.com/chenliu1993/podsecurity-check/pkg/generator"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/pointer"
)

func fixturesPrivileged() Case {
	return Case{
		Name: "privileged",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Containers[0].SecurityContext.Privileged = pointer.Bool(false)
					return pod
				},
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.InitContainers[0].SecurityContext.Privileged = pointer.Bool(false)
					return pod
				},
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Containers[0].SecurityContext.Privileged = pointer.Bool(true)
					pod.Template.Spec.Containers[0].SecurityContext.AllowPrivilegeEscalation = nil
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.InitContainers[0].SecurityContext.Privileged = pointer.Bool(true)
					pod.Template.Spec.InitContainers[0].SecurityContext.AllowPrivilegeEscalation = nil
					return pod
				},
			}),
		},
	}
}

func fixturesAllowPrivilegeEscalation() Case {
	return Case{
		Name: "allowPrivilegeEscalation",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Containers[0].SecurityContext.AllowPrivilegeEscalation = pointer.Bool(true)
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.InitContainers[0].SecurityContext.AllowPrivilegeEscalation = pointer.Bool(true)
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.InitContainers[0].SecurityContext.AllowPrivilegeEscalation = nil
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Containers[0].SecurityContext = nil
					return pod
				},
			}),
		},
	}
}

func init() {
	fixturesMap[privileged] = []func() Case{fixturesPrivileged}
	fixturesMap[allowPrivilegeEscalation] = []func() Case{fixturesAllowPrivilegeEscalation}
}
