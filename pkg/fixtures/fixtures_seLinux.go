package fixtures

import (
	"github.com/chenliu1993/podsecurity-check/pkg/constants"
	corev1 "k8s.io/api/core/v1"
)

func fixturesSELinux() Case {
	return Case{
		Name: "selinux",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.SELinux(constants.POD, &ContainerInitT),
				g.SELinux(constants.CONTAINER, &ContainerT),
				g.SELinux(constants.INITCONTAINER, &ContainerKvmT),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.SELinux(constants.POD, &EmptyUserRoleOpts),
				g.SELinux(constants.CONTAINER, &EmptyUserRoleOpts),
				g.SELinux(constants.INITCONTAINER, &EmptyUserRoleOpts),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.SELinux(constants.POD, &ContainerInitKvmT),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.SELinux(constants.POD, nil),
				g.SELinux(constants.CONTAINER, &ContainerInitKvmT),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.SELinux(constants.POD, nil),
				g.SELinux(constants.INITCONTAINER, &ContainerInitKvmT),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.SELinux(constants.POD, &NonemptyUserRoleOpts),
				g.SELinux(constants.CONTAINER, &NonemptyUserRoleOpts),
				g.SELinux(constants.INITCONTAINER, &NonemptyUserRoleOpts),
			}),
		},
	}
}

func fixturesEphemeralSELinux() Case {
	return Case{
		Name: "ephemeral-selinux",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.SELinux(constants.POD, &ContainerInitT),
				g.SELinux(constants.EPHEMERALCONTAINER, &ContainerT),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.SELinux(constants.POD, &EmptyUserRoleOpts),
				g.SELinux(constants.EPHEMERALCONTAINER, &EmptyUserRoleOpts),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.SELinux(constants.POD, &ContainerInitKvmT),
				g.SELinux(constants.EPHEMERALCONTAINER, &ContainerInitKvmT),
			}),
			g.Overlay([]func() *corev1.PodTemplate{
				g.SELinux(constants.POD, &NonemptyUserRoleOpts),
				g.SELinux(constants.EPHEMERALCONTAINER, &NonemptyUserRoleOpts),
			}),
		},
	}
}

func init() {
	fixturesMap["selinux"] = []func() Case{
		fixturesSELinux,
	}
	fixturesMap["ephemeralSELinux"] = []func() Case{
		fixturesEphemeralSELinux,
	}
}
