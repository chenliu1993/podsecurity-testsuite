package fixtures

import (
	corev1 "k8s.io/api/core/v1"
)

func fixturesSysctls() Case {
	return Case{
		Name:   "sysctls",
		ErrMsg: "forbidden sysctl",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					if pod.Template.Spec.SecurityContext == nil {
						pod.Template.Spec.SecurityContext = &corev1.PodSecurityContext{}
					}
					pod.Template.Spec.SecurityContext.Sysctls = nil
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					if pod.Template.Spec.SecurityContext == nil {
						pod.Template.Spec.SecurityContext = &corev1.PodSecurityContext{}
					}
					pod.Template.Spec.SecurityContext.Sysctls = []corev1.Sysctl{
						{Name: "kernel.shm_rmid_forced", Value: "0"},
						{Name: "net.ipv4.ip_local_port_range", Value: "1024 65535"},
						{Name: "net.ipv4.tcp_syncookies", Value: "0"},
						{Name: "net.ipv4.ping_group_range", Value: "1 0"},
						{Name: "net.ipv4.ip_unprivileged_port_start", Value: "1024"},
					}
					return pod
				},
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					if pod.Template.Spec.SecurityContext == nil {
						pod.Template.Spec.SecurityContext = &corev1.PodSecurityContext{}
					}
					pod.Template.Spec.SecurityContext.Sysctls = []corev1.Sysctl{{Name: "othersysctl", Value: "other"}}
					return pod
				},
			}),
		},
	}
}

func init() {
	fixturesMap[sysctls] = []func() Case{fixturesSysctls}
}
