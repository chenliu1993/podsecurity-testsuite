package fixtures

import (
	"github.com/chenliu1993/podsecurity-check/pkg/constants"
	corev1 "k8s.io/api/core/v1"
)

func fixturesHostPID() Case {
	return Case{
		Name: "hostPID",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.HostPID(false),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.HostPID(true),
			}),
		},
	}
}

func fixturesHostIPC() Case {
	return Case{
		Name: "hostIPC",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.HostIPC(false),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.HostIPC(true),
			}),
		},
	}
}

func fixturesHostNetwork() Case {
	return Case{
		Name: "hostNetwork",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.HostNetwork(false),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.HostNetwork(true),
			}),
		},
	}
}

func fixturesConainerHostPorts() Case {
	return Case{
		Name: "hostports",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.HostPort(constants.CONTAINER, HostPortBaselinePass),
				g.HostPort(constants.INITCONTAINER, HostPortBaselinePass),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.HostPort(constants.CONTAINER, HostPortBaselinePass),
				g.HostPort(constants.INITCONTAINER, HostPortBaselinePass),
			}),
		},
	}
}

func fixturesEphemeralHostPorts() Case {
	return Case{
		Name: "ephemeral-hostports",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.HostPort(constants.EPHEMERALCONTAINER, HostPortBaselinePass),
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func() *corev1.PodTemplate{
				g.HostPort(constants.EPHEMERALCONTAINER, HostPortBaselinePass),
			}),
		},
	}
}

func init() {
	fixturesMap["hostpid"] = []func() Case{fixturesHostPID}
	fixturesMap["hostipc"] = []func() Case{fixturesHostIPC}
	fixturesMap["hostnetwork"] = []func() Case{fixturesHostNetwork}
	fixturesMap["hostports"] = []func() Case{fixturesConainerHostPorts}
	fixturesMap["ephemeralhostports"] = []func() Case{fixturesEphemeralHostPorts}
}
