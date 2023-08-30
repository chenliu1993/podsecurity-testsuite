package fixtures

import (
	"github.com/chenliu1993/podsecurity-check/pkg/generator"
	corev1 "k8s.io/api/core/v1"
)

func fixturesSELinux() Case {
	return Case{
		Name:   "selinux",
		ErrMsg: "seLinuxOptions",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.SELinuxOptions = nil
					pod.Template.Spec.Containers[0].SecurityContext.SELinuxOptions = nil
					pod.Template.Spec.InitContainers[0].SecurityContext.SELinuxOptions = &corev1.SELinuxOptions{}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				generator.EnsureSELinuxOptions,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.SELinuxOptions.Type = "container_t"
					pod.Template.Spec.Containers[0].SecurityContext.SELinuxOptions.Type = "container_kvm_t"
					pod.Template.Spec.InitContainers[0].SecurityContext.SELinuxOptions.Type = "container_init_t"
					pod.Template.Spec.SecurityContext.SELinuxOptions.Level = ""
					pod.Template.Spec.Containers[0].SecurityContext.SELinuxOptions.Level = "empty"
					return pod
				},
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				generator.EnsureSELinuxOptions,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.SELinuxOptions.Type = "somevalue"
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				generator.EnsureSELinuxOptions,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Containers[0].SecurityContext.SELinuxOptions.Type = "somevalue"
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				generator.EnsureSELinuxOptions,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.InitContainers[0].SecurityContext.SELinuxOptions.Type = "somevalue"
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				generator.EnsureSELinuxOptions,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.SELinuxOptions.User = "somevalue"
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				generator.EnsureSELinuxOptions,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.SELinuxOptions.Role = "somevalue"
					return pod
				},
			}),
		},
	}
}

func init() {
	fixturesMap[selinux] = []func() Case{
		fixturesSELinux,
	}
}
