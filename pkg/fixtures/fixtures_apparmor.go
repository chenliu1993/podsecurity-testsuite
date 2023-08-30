package fixtures

import (
	corev1 "k8s.io/api/core/v1"
)

func fixturesApparmor() Case {
	return Case{
		Name:   "apparmor",
		ErrMsg: "forbidden AppArmor profile",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					if pod.Annotations == nil {
						pod.Annotations = make(map[string]string)
					}
					pod.Annotations[corev1.AppArmorBetaContainerAnnotationKeyPrefix+pod.Template.Spec.Containers[0].Name] = "runtime/default"
					pod.Annotations[corev1.AppArmorBetaContainerAnnotationKeyPrefix+pod.Template.Spec.InitContainers[0].Name] = "localhost/foo"
					return pod
				},
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					if pod.Annotations == nil {
						pod.Annotations = make(map[string]string)
					}
					pod.Annotations[corev1.AppArmorBetaContainerAnnotationKeyPrefix+pod.Template.Spec.Containers[0].Name] = "unconfined"
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					if pod.Annotations == nil {
						pod.Annotations = make(map[string]string)
					}
					pod.Annotations[corev1.AppArmorBetaContainerAnnotationKeyPrefix+pod.Template.Spec.InitContainers[0].Name] = "unconfined"
					return pod
				},
			}),
		},
	}
}

func init() {
	fixturesMap[apparmor] = []func() Case{fixturesApparmor}
}
