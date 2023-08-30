package fixtures

import (
	"github.com/chenliu1993/podsecurity-check/pkg/generator"
	corev1 "k8s.io/api/core/v1"
)

func fixturesSeccompBaseline() Case {
	return Case{
		Name:   "seccomp-baseline",
		ErrMsg: "seccompProfile",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.SeccompProfile = &corev1.SeccompProfile{
						Type: corev1.SeccompProfileTypeRuntimeDefault,
					}
					pod.Template.Spec.Containers[0].SecurityContext.SeccompProfile = &corev1.SeccompProfile{
						Type: corev1.SeccompProfileTypeRuntimeDefault,
					}
					return pod
				},
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.SeccompProfile = &corev1.SeccompProfile{
						Type: corev1.SeccompProfileTypeUnconfined,
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Containers[0].SecurityContext.SeccompProfile = &corev1.SeccompProfile{
						Type: corev1.SeccompProfileTypeUnconfined,
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.InitContainers[0].SecurityContext.SeccompProfile = &corev1.SeccompProfile{
						Type: corev1.SeccompProfileTypeUnconfined,
					}
					return pod
				},
			}),
		},
	}
}

func fixturesSeccompRestricted() Case {
	return Case{
		Name:   "seccomp-restricted",
		ErrMsg: "seccompProfile",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.SeccompProfile = &corev1.SeccompProfile{
						Type: corev1.SeccompProfileTypeRuntimeDefault,
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					profile := "foo"
					pod.Template.Spec.SecurityContext.SeccompProfile = &corev1.SeccompProfile{
						Type:             corev1.SeccompProfileTypeLocalhost,
						LocalhostProfile: &profile,
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					profile := "foo"
					pod.Template.Spec.SecurityContext.SeccompProfile = nil
					pod.Template.Spec.Containers[0].SecurityContext.SeccompProfile = &corev1.SeccompProfile{
						Type: corev1.SeccompProfileTypeRuntimeDefault,
					}
					pod.Template.Spec.InitContainers[0].SecurityContext.SeccompProfile = &corev1.SeccompProfile{
						Type:             corev1.SeccompProfileTypeLocalhost,
						LocalhostProfile: &profile,
					}
					return pod
				},
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.SeccompProfile = nil
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.SeccompProfile = &corev1.SeccompProfile{
						Type: corev1.SeccompProfileTypeUnconfined,
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.SeccompProfile = nil
					pod.Template.Spec.Containers[0].SecurityContext.SeccompProfile = &corev1.SeccompProfile{
						Type: corev1.SeccompProfileTypeRuntimeDefault,
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.SeccompProfile = nil
					pod.Template.Spec.InitContainers[0].SecurityContext.SeccompProfile = &corev1.SeccompProfile{
						Type: corev1.SeccompProfileTypeRuntimeDefault,
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				generator.EnsureSecurityContext,
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.SecurityContext.SeccompProfile = nil
					pod.Template.Spec.Containers[0].SecurityContext.SeccompProfile = &corev1.SeccompProfile{
						Type: corev1.SeccompProfileTypeRuntimeDefault,
					}
					pod.Template.Spec.InitContainers[0].SecurityContext.SeccompProfile = &corev1.SeccompProfile{
						Type: corev1.SeccompProfileTypeRuntimeDefault,
					}
					return pod
				},
			}),
		},
	}
}

func init() {
	fixturesMap[seccomp] = []func() Case{fixturesSeccompBaseline}
	fixturesMap[seccompRestricted] = []func() Case{fixturesSeccompRestricted}
}
