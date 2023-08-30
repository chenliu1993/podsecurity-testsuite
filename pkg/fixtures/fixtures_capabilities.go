package fixtures

import (
	"github.com/chenliu1993/podsecurity-check/pkg/generator"
	corev1 "k8s.io/api/core/v1"
)

func ensureCapabilities(pod *corev1.PodTemplate) *corev1.PodTemplate {
	pod = generator.EnsureSecurityContext(pod)
	for i := range pod.Template.Spec.Containers {
		if pod.Template.Spec.Containers[i].SecurityContext.Capabilities == nil {
			pod.Template.Spec.Containers[i].SecurityContext.Capabilities = &corev1.Capabilities{}
		}
	}
	for i := range pod.Template.Spec.InitContainers {
		if pod.Template.Spec.InitContainers[i].SecurityContext.Capabilities == nil {
			pod.Template.Spec.InitContainers[i].SecurityContext.Capabilities = &corev1.Capabilities{}
		}
	}
	return pod
}

func fixturesCapabilitiesBaseline() Case {
	return Case{
		Name:   "capabilities-baseline",
		ErrMsg: "capabilities",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					if pod.Template.Spec.Containers[0].SecurityContext != nil && pod.Template.Spec.Containers[0].SecurityContext.Capabilities != nil {
						for _, capability := range pod.Template.Spec.Containers[0].SecurityContext.Capabilities.Drop {
							if capability == corev1.Capability("ALL") {
								return nil
							}
						}
					}
					pod = ensureCapabilities(pod)
					pod.Template.Spec.Containers[0].SecurityContext.Capabilities.Add = []corev1.Capability{
						"AUDIT_WRITE", "CHOWN", "DAC_OVERRIDE", "FOWNER", "FSETID", "KILL", "MKNOD", "NET_BIND_SERVICE", "SETFCAP", "SETGID", "SETPCAP", "SETUID", "SYS_CHROOT",
					}
					pod.Template.Spec.InitContainers[0].SecurityContext.Capabilities.Add = []corev1.Capability{
						"AUDIT_WRITE", "CHOWN", "DAC_OVERRIDE", "FOWNER", "FSETID", "KILL", "MKNOD", "NET_BIND_SERVICE", "SETFCAP", "SETGID", "SETPCAP", "SETUID", "SYS_CHROOT",
					}
					return pod
				},
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod = ensureCapabilities(pod)
					pod.Template.Spec.Containers[0].SecurityContext.Capabilities.Add = []corev1.Capability{
						"chmod",
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod = ensureCapabilities(pod)
					pod.Template.Spec.Containers[0].SecurityContext.Capabilities.Add = []corev1.Capability{
						"CAP_CHOWN", "NET_RAW",
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod = ensureCapabilities(pod)
					pod.Template.Spec.InitContainers[0].SecurityContext.Capabilities.Add = []corev1.Capability{
						"CAP_CHOWN", "NET_RAW",
					}
					return pod
				},
			}),
		},
	}
}
func fixturesCapabilitiesRestricted() Case {
	return Case{
		Name:   "capabilities-restricted",
		ErrMsg: "unrestricted capabilities",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod = ensureCapabilities(pod)
					pod.Template.Spec.Containers[0].SecurityContext.Capabilities.Drop = []corev1.Capability{"ALL"}
					pod.Template.Spec.InitContainers[0].SecurityContext.Capabilities.Drop = []corev1.Capability{"ALL"}
					pod.Template.Spec.Containers[0].SecurityContext.Capabilities.Add = []corev1.Capability{"NET_BIND_SERVICE"}
					pod.Template.Spec.InitContainers[0].SecurityContext.Capabilities.Add = []corev1.Capability{"NET_BIND_SERVICE"}
					return pod
				},
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod = ensureCapabilities(pod)
					pod.Template.Spec.Containers[0].SecurityContext.Capabilities.Drop = []corev1.Capability{}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod = ensureCapabilities(pod)
					pod.Template.Spec.InitContainers[0].SecurityContext.Capabilities.Drop = []corev1.Capability{}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod = ensureCapabilities(pod)
					pod.Template.Spec.Containers[0].SecurityContext.Capabilities.Drop = []corev1.Capability{
						"SYS_TIME", "SYS_MODULE", "SYS_RAWIO", "SYS_PACCT", "SYS_ADMIN", "SYS_NICE",
						"SYS_RESOURCE", "SYS_TIME", "SYS_TTY_CONFIG", "MKNOD", "AUDIT_WRITE",
						"AUDIT_CONTROL", "MAC_OVERRIDE", "MAC_ADMIN", "NET_ADMIN", "SYSLOG",
						"CHOWN", "NET_RAW", "DAC_OVERRIDE", "FOWNER", "DAC_READ_SEARCH",
						"FSETID", "KILL", "SETGID", "SETUID", "LINUX_IMMUTABLE", "NET_BIND_SERVICE",
						"NET_BROADCAST", "IPC_LOCK", "IPC_OWNER", "SYS_CHROOT", "SYS_PTRACE",
						"SYS_BOOT", "LEASE", "SETFCAP", "WAKE_ALARM", "BLOCK_SUSPEND",
					}
					pod.Template.Spec.InitContainers[0].SecurityContext.Capabilities.Drop = []corev1.Capability{
						"SYS_TIME", "SYS_MODULE", "SYS_RAWIO", "SYS_PACCT", "SYS_ADMIN", "SYS_NICE",
						"SYS_RESOURCE", "SYS_TIME", "SYS_TTY_CONFIG", "MKNOD", "AUDIT_WRITE",
						"AUDIT_CONTROL", "MAC_OVERRIDE", "MAC_ADMIN", "NET_ADMIN", "SYSLOG",
						"CHOWN", "NET_RAW", "DAC_OVERRIDE", "FOWNER", "DAC_READ_SEARCH",
						"FSETID", "KILL", "SETGID", "SETUID", "LINUX_IMMUTABLE", "NET_BIND_SERVICE",
						"NET_BROADCAST", "IPC_LOCK", "IPC_OWNER", "SYS_CHROOT", "SYS_PTRACE",
						"SYS_BOOT", "LEASE", "SETFCAP", "WAKE_ALARM", "BLOCK_SUSPEND",
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod = ensureCapabilities(pod)
					pod.Template.Spec.Containers[0].SecurityContext.Capabilities.Drop = []corev1.Capability{"ALL"}
					pod.Template.Spec.InitContainers[0].SecurityContext.Capabilities.Drop = []corev1.Capability{"ALL"}
					pod.Template.Spec.Containers[0].SecurityContext.Capabilities.Add = []corev1.Capability{
						"AUDIT_WRITE", "CHOWN", "DAC_OVERRIDE", "FOWNER", "FSETID", "KILL", "MKNOD", "NET_BIND_SERVICE", "SETFCAP", "SETGID", "SETPCAP", "SETUID", "SYS_CHROOT",
					}
					pod.Template.Spec.InitContainers[0].SecurityContext.Capabilities.Add = []corev1.Capability{
						"AUDIT_WRITE", "CHOWN", "DAC_OVERRIDE", "FOWNER", "FSETID", "KILL", "MKNOD", "NET_BIND_SERVICE", "SETFCAP", "SETGID", "SETPCAP", "SETUID", "SYS_CHROOT",
					}
					return pod
				},
			}),
		},
	}
}

func init() {
	fixturesMap[capabilitiesBaseline] = []func() Case{
		fixturesCapabilitiesBaseline,
	}
	fixturesMap[capabilitiesRestricted] = []func() Case{
		fixturesCapabilitiesRestricted,
	}
}
