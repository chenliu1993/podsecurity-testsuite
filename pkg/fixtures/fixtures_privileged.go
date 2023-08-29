package fixtures

import (
	"github.com/chenliu1993/podsecurity-check/pkg/constants"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/pointer"
)

func fixturesPrivileged() Case {
	return Case{
		Name: "privileged",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.EnsureSecurityContext,
				g.Privileged(pointer.Bool(false), constants.CONTAINER),
				g.Privileged(pointer.Bool(false), constants.INITCONTAINER),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.EnsureSecurityContext,
				g.Privileged(pointer.Bool(true), constants.CONTAINER),
				g.AllowPrivilegeEscalation(nil, constants.CONTAINER),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.Privileged(pointer.Bool(true), constants.INITCONTAINER),
				g.AllowPrivilegeEscalation(nil, constants.INITCONTAINER),
			}),
		},
	}
}

func fixturesEphemeralPrivileged() Case {
	return Case{
		Name: "ephemeral-privileged",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.EnsureSecurityContext,
				g.Privileged(pointer.Bool(false), constants.EPHEMERALCONTAINER),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.EnsureSecurityContext,
				g.Privileged(pointer.Bool(true), constants.EPHEMERALCONTAINER),
				g.AllowPrivilegeEscalation(nil, constants.EPHEMERALCONTAINER),
			}),
		},
	}
}

func fixturesAllowPrivilegeEscalation() Case {
	return Case{
		Name: "allowPrivilegeEscalation",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.EnsureSecurityContext,
				g.AllowPrivilegeEscalation(pointer.Bool(true), constants.CONTAINER),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.EnsureSecurityContext,
				g.AllowPrivilegeEscalation(pointer.Bool(true), constants.INITCONTAINER),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.AllowPrivilegeEscalation(nil, constants.CONTAINER),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.EnsureSecurityContext,
				func() *corev1.PodTemplate {
					pod := g.GetPod()
					pod.Template.Spec.Containers[0].SecurityContext = nil
					return pod
				},
			}),
		},
	}
}

func fixturesEphemeralAllowPrivilegeEscalation() Case {
	return Case{
		Name: "ephemeral-allowPrivilegeEscalation",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.AllowPrivilegeEscalation(pointer.Bool(true), constants.EPHEMERALCONTAINER),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.AllowPrivilegeEscalation(pointer.Bool(true), constants.EPHEMERALCONTAINER),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.AllowPrivilegeEscalation(nil, constants.EPHEMERALCONTAINER),
			}),
		},
	}
}

func init() {
	fixturesMap[privileged] = []func() Case{fixturesPrivileged}
	fixturesMap[allowPrivilegeEscalation] = []func() Case{fixturesAllowPrivilegeEscalation}
	fixturesMap[ephemeralPrivileged] = []func() Case{fixturesEphemeralPrivileged}
	fixturesMap[ephemeralAllowPrivilegeEscalation] = []func() Case{fixturesEphemeralAllowPrivilegeEscalation}
}
