package fixtures

import (
	corev1 "k8s.io/api/core/v1"
)

func fixturesVolumesBaseline() Case {
	return Case{
		Name:   "volumes-baseline",
		ErrMsg: "hostPath",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{
							Name: "volume-emptydir",
							VolumeSource: corev1.VolumeSource{
								EmptyDir: &corev1.EmptyDirVolumeSource{},
							},
						},
						{
							Name: "volume-hostpath0",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/fixtures",
								},
							},
						},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{
							Name: "volume-hostpath1",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/fixtures",
								},
							},
						},
					}
					return pod
				},
			}),
		},
	}
}

func fixturesVolumesRestricted() Case {
	return Case{
		Name:   "volumes-restricted",
		ErrMsg: "restricted volume types",
		PassSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume0", VolumeSource: corev1.VolumeSource{}},
						{Name: "volume1", VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}}},
						{Name: "volume2", VolumeSource: corev1.VolumeSource{Secret: &corev1.SecretVolumeSource{SecretName: "test"}}},
						{Name: "volume3", VolumeSource: corev1.VolumeSource{PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{ClaimName: "test"}}},
						{Name: "volume4", VolumeSource: corev1.VolumeSource{DownwardAPI: &corev1.DownwardAPIVolumeSource{Items: []corev1.DownwardAPIVolumeFile{{Path: "labels", FieldRef: &corev1.ObjectFieldSelector{FieldPath: "metadata.labels"}}}}}},
						{Name: "volume5", VolumeSource: corev1.VolumeSource{ConfigMap: &corev1.ConfigMapVolumeSource{LocalObjectReference: corev1.LocalObjectReference{Name: "test"}}}},
						{Name: "volume6", VolumeSource: corev1.VolumeSource{Projected: &corev1.ProjectedVolumeSource{Sources: []corev1.VolumeProjection{}}}},
					}
					return pod
				},
			}),
		},
		FailSet: []*corev1.PodTemplate{
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{GCEPersistentDisk: &corev1.GCEPersistentDiskVolumeSource{PDName: "test"}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{AWSElasticBlockStore: &corev1.AWSElasticBlockStoreVolumeSource{VolumeID: "test"}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{GitRepo: &corev1.GitRepoVolumeSource{Repository: "github.com/kubernetes/kubernetes"}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{NFS: &corev1.NFSVolumeSource{Server: "fixtures", Path: "/fxitures"}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{ISCSI: &corev1.ISCSIVolumeSource{TargetPortal: "fixtures", IQN: "iqn.2023-08.com.example:storage.kube.sys1.xyz", Lun: 0}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{Glusterfs: &corev1.GlusterfsVolumeSource{Path: "fixtures", EndpointsName: "fixtures"}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{RBD: &corev1.RBDVolumeSource{CephMonitors: []string{"fixtures"}, RBDImage: "fixtures"}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{FlexVolume: &corev1.FlexVolumeSource{Driver: "test"}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{Cinder: &corev1.CinderVolumeSource{VolumeID: "test"}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{CephFS: &corev1.CephFSVolumeSource{Monitors: []string{"test"}}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{Flocker: &corev1.FlockerVolumeSource{DatasetName: "test"}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{FC: &corev1.FCVolumeSource{WWIDs: []string{"test"}}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{AzureFile: &corev1.AzureFileVolumeSource{SecretName: "test", ShareName: "test"}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{VsphereVolume: &corev1.VsphereVirtualDiskVolumeSource{VolumePath: "test"}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{Quobyte: &corev1.QuobyteVolumeSource{Registry: "localhost:8080", Volume: "test"}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{AzureDisk: &corev1.AzureDiskVolumeSource{DiskName: "test", DataDiskURI: "https://test.blob.core.windows.net/test/test.vhd"}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{PortworxVolume: &corev1.PortworxVolumeSource{VolumeID: "test", FSType: "ext4"}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{ScaleIO: &corev1.ScaleIOVolumeSource{VolumeName: "test", Gateway: "localhost", System: "test"}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{StorageOS: &corev1.StorageOSVolumeSource{VolumeName: "test"}}},
					}
					return pod
				},
			}),
			g.Overlay([]func(*corev1.PodTemplate) *corev1.PodTemplate{
				func(pod *corev1.PodTemplate) *corev1.PodTemplate {
					pod.Template.Spec.Volumes = []corev1.Volume{
						{Name: "volume1", VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: "/dev/null"}}},
					}
					return pod
				},
			}),
		},
	}
}

func init() {
	fixturesMap[volumesBaseline] = []func() Case{fixturesVolumesBaseline}
	fixturesMap[volumesRestricted] = []func() Case{fixturesVolumesRestricted}
}
