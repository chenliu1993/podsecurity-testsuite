package fixtures

import (
	corev1 "k8s.io/api/core/v1"
)

func fixturesHostPID() Case {
	return Case{
		Name:   "hostPID",
		ErrMsg: "host namespaces",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.HostPID = false
					return pod
				},
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.HostPID = true
					return pod
				},
			}),
		},
	}
}

func fixturesHostIPC() Case {
	return Case{
		Name:   "hostIPC",
		ErrMsg: "host namespaces",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.HostIPC = false
					return pod
				},
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.HostIPC = true
					return pod
				},
			}),
		},
	}
}

func fixturesHostNetwork() Case {
	return Case{
		Name:   "hostNetwork",
		ErrMsg: "host namespaces",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.HostNetwork = false
					return pod
				},
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.HostNetwork = true
					return pod
				},
			}),
		},
	}
}

func fixturesConainerHostPorts() Case {
	return Case{
		Name:   "hostports",
		ErrMsg: "hostPort",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Containers[0].Ports = []corev1.ContainerPort{
						{
							ContainerPort: 65535,
						},
					}
					pod.Template.Spec.InitContainers[0].Ports = []corev1.ContainerPort{
						{
							ContainerPort: 65535,
						},
					}
					return pod
				},
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Containers[0].Ports = []corev1.ContainerPort{
						{
							ContainerPort: 65535,
							HostPort:      65535,
						},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.InitContainers[0].Ports = []corev1.ContainerPort{
						{
							ContainerPort: 65535,
							HostPort:      65535,
						},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Containers[0].Ports = []corev1.ContainerPort{
						{
							ContainerPort: 65535,
							HostPort:      65535,
						},
						{
							ContainerPort: 4444,
						},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.InitContainers[0].Ports = []corev1.ContainerPort{
						{
							ContainerPort: 65535,
							HostPort:      65535,
						},
						{
							ContainerPort: 4444,
						},
					}
					return pod
				},
			}),
		},
	}
}

func init() {
	fixturesMap[hostpid] = []func() Case{fixturesHostPID}
	fixturesMap[hostipc] = []func() Case{fixturesHostIPC}
	fixturesMap[hostnetwork] = []func() Case{fixturesHostNetwork}
	fixturesMap[hostports] = []func() Case{fixturesConainerHostPorts}
}
