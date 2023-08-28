package generator

import (
	"github.com/chenliu1993/podsecurity-check/pkg/constants"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/pointer"
)

// NonRoot modifies the runAsNonRoot field of the SecurityContext
// restricted: true
func (g *Generator) NonRoot(contType string, allow bool) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		pod := ensureSecurityContext(g.pod, contType)
		switch contType {
		case constants.POD:
			pod.Template.Spec.SecurityContext.RunAsNonRoot = pointer.Bool(allow)
		case constants.CONTAINER:
			pod.Template.Spec.Containers[0].SecurityContext.RunAsNonRoot = pointer.Bool(allow)
		case constants.INITCONTAINER:
			pod.Template.Spec.InitContainers[0].SecurityContext.RunAsNonRoot = pointer.Bool(allow)
		case constants.EPHEMERALCONTAINER:
			pod.Template.Spec.EphemeralContainers[0].EphemeralContainerCommon.SecurityContext.RunAsNonRoot = pointer.Bool(allow)
		}
		return pod
	}
}

// RunAsUser modifies the runAsUser field of the SecurityContext
// restricted: non-zero or nil/undefined values
func (g *Generator) RunAsUser(contType string, userID int64) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		pod := ensureSecurityContext(g.pod, contType)
		switch contType {
		case constants.POD:
			pod.Template.Spec.SecurityContext.RunAsUser = pointer.Int64(userID)
		case constants.CONTAINER:
			pod.Template.Spec.Containers[0].SecurityContext.RunAsUser = pointer.Int64(userID)
		case constants.INITCONTAINER:
			pod.Template.Spec.InitContainers[0].SecurityContext.RunAsUser = pointer.Int64(userID)
		case constants.EPHEMERALCONTAINER:
			pod.Template.Spec.EphemeralContainers[0].EphemeralContainerCommon.SecurityContext.RunAsUser = pointer.Int64(userID)
		}
		return pod
	}
}

// AppArmorProfile sets the AppArmor profile of the container
// The baseline allowed values are:
//
//	runtime/default: The default AppArmor profile for the container runtime.
//	unconfined: No AppArmor profile will be applied.
//	localhost/*: self-defined files
func (g *Generator) AppArmorProfile(profile string) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		if g.pod.Template.Annotations == nil {
			g.pod.Template.Annotations = make(map[string]string)
		}
		g.pod.Template.Annotations["container.apparmor.security.beta.kubernetes.io/"+g.pod.Template.Spec.Containers[0].Name] = profile
		return g.pod
	}
}

// CapabilitiesXXX sets the capabilities of the container based on different security level

func (g *Generator) CapabilitiesAdd(contType string, addSet []corev1.Capability) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		pod := ensureSecurityContext(g.pod, contType)
		switch contType {
		case constants.CONTAINER:
			pod.Template.Spec.Containers[0].SecurityContext.Capabilities = &corev1.Capabilities{
				Add: addSet,
			}
		case constants.INITCONTAINER:
			pod.Template.Spec.InitContainers[0].SecurityContext.Capabilities = &corev1.Capabilities{
				Add: addSet,
			}
		case constants.EPHEMERALCONTAINER:
			pod.Template.Spec.EphemeralContainers[0].EphemeralContainerCommon.SecurityContext.Capabilities = &corev1.Capabilities{
				Add: addSet,
			}
		}
		return pod
	}
}

func (g *Generator) CapabilitiesDrop(contType string, dropSet []corev1.Capability) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		pod := ensureSecurityContext(g.pod, contType)
		switch contType {
		case constants.CONTAINER:
			pod.Template.Spec.Containers[0].SecurityContext.Capabilities = &corev1.Capabilities{
				Drop: dropSet,
			}
		case constants.INITCONTAINER:
			pod.Template.Spec.InitContainers[0].SecurityContext.Capabilities = &corev1.Capabilities{
				Drop: dropSet,
			}
		case constants.EPHEMERALCONTAINER:
			pod.Template.Spec.EphemeralContainers[0].EphemeralContainerCommon.SecurityContext.Capabilities = &corev1.Capabilities{
				Drop: dropSet,
			}
		}
		return pod
	}
}

// HostNamepsace-related security
// baseline/restricted: nil/false
func (g *Generator) HostNetwork(allow bool) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		g.pod.Template.Spec.HostNetwork = allow
		return g.pod
	}
}

func (g *Generator) HostPID(allow bool) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		g.pod.Template.Spec.HostPID = allow
		return g.pod
	}
}

func (g *Generator) HostIPC(allow bool) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		g.pod.Template.Spec.HostIPC = allow
		return g.pod
	}
}

// HostPort sets the host port of the container
// baseline: undefined/nil or 0
func (g *Generator) HostPort(contType string, hostport []corev1.ContainerPort) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		switch contType {
		case constants.CONTAINER:
			g.pod.Template.Spec.Containers[0].Ports = hostport
		case constants.INITCONTAINER:
			g.pod.Template.Spec.InitContainers[0].Ports = hostport
		case constants.EPHEMERALCONTAINER:
			g.pod.Template.Spec.EphemeralContainers[0].EphemeralContainerCommon.Ports = hostport
		}
		return g.pod
	}
}

// Proviledged sets if containers can be privileged
// baseline/restricted: nil/false
func (g *Generator) Privileged(allow bool, contType string) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		pod := ensureSecurityContext(g.pod, contType)
		switch contType {
		case constants.CONTAINER:
			pod.Template.Spec.Containers[0].SecurityContext.Privileged = pointer.Bool(allow)
		case constants.INITCONTAINER:
			pod.Template.Spec.InitContainers[0].SecurityContext.Privileged = pointer.Bool(allow)
		case constants.EPHEMERALCONTAINER:
			pod.Template.Spec.EphemeralContainers[0].EphemeralContainerCommon.SecurityContext.Privileged = pointer.Bool(allow)
		}
		return pod
	}
}

func (g *Generator) AllowPrivilegeEscalation(allow bool, contType string) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		pod := ensureSecurityContext(g.pod, contType)
		switch contType {
		case constants.CONTAINER:
			pod.Template.Spec.Containers[0].SecurityContext.AllowPrivilegeEscalation = pointer.Bool(allow)
		case constants.INITCONTAINER:
			pod.Template.Spec.InitContainers[0].SecurityContext.AllowPrivilegeEscalation = pointer.Bool(allow)
		case constants.EPHEMERALCONTAINER:
			pod.Template.Spec.EphemeralContainers[0].EphemeralContainerCommon.SecurityContext.AllowPrivilegeEscalation = pointer.Bool(allow)
		}
		return pod
	}
}

// procmount sets the proc mount type of the container
// baseline: undefined/nil or DefaultProcMount
func (g *Generator) ProcMount(contType string, proc *corev1.ProcMountType) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		pod := ensureSecurityContext(g.pod, contType)
		switch contType {
		case constants.CONTAINER:
			pod.Template.Spec.Containers[0].SecurityContext.ProcMount = proc
		case constants.INITCONTAINER:
			pod.Template.Spec.InitContainers[0].SecurityContext.ProcMount = proc
		case constants.EPHEMERALCONTAINER:
			pod.Template.Spec.EphemeralContainers[0].EphemeralContainerCommon.SecurityContext.ProcMount = proc
		}
		return pod
	}
}

// Seccomp sets the seccomp fields of the container
// baseline: undefined/nil or RuntimeDefault or Localhost
// restricted: RuntimeDefault or Localhost
func (g *Generator) Seccomp(contType string, seccompProfile *corev1.SeccompProfile) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		pod := ensureSecurityContext(g.pod, contType)
		switch contType {
		case constants.POD:
			pod.Template.Spec.SecurityContext.SeccompProfile = seccompProfile
		case constants.CONTAINER:
			pod.Template.Spec.Containers[0].SecurityContext.SeccompProfile = seccompProfile
		case constants.INITCONTAINER:
			pod.Template.Spec.InitContainers[0].SecurityContext.SeccompProfile = seccompProfile
		case constants.EPHEMERALCONTAINER:
			pod.Template.Spec.EphemeralContainers[0].EphemeralContainerCommon.SecurityContext.SeccompProfile = seccompProfile
		}
		return pod
	}
}

// SELinux sets the selinux options of the container
// baseline: undefined/nil or container_t or container_kvm_t or container_init_t
func (g *Generator) SELinux(contType string, selinuxOpts *corev1.SELinuxOptions) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		pod := ensureSecurityContext(g.pod, contType)
		switch contType {
		case constants.POD:
			pod.Template.Spec.SecurityContext.SELinuxOptions = selinuxOpts
		case constants.CONTAINER:
			for i := range pod.Template.Spec.Containers {
				if pod.Template.Spec.Containers[i].SecurityContext.SELinuxOptions == nil {
					pod.Template.Spec.Containers[i].SecurityContext.SELinuxOptions = &corev1.SELinuxOptions{}
				}
			}
			pod.Template.Spec.Containers[0].SecurityContext.SELinuxOptions = selinuxOpts
		case constants.INITCONTAINER:
			for i := range pod.Template.Spec.InitContainers {
				if pod.Template.Spec.InitContainers[i].SecurityContext.SELinuxOptions == nil {
					pod.Template.Spec.InitContainers[i].SecurityContext.SELinuxOptions = &corev1.SELinuxOptions{}
				}
			}
			pod.Template.Spec.InitContainers[0].SecurityContext.SELinuxOptions = selinuxOpts
		case constants.EPHEMERALCONTAINER:
			pod.Template.Spec.EphemeralContainers[0].EphemeralContainerCommon.SecurityContext.SELinuxOptions = selinuxOpts
		}
		return pod
	}
}

// Sysctls sets the sysctls for the container.
func (g *Generator) Sysctls(sysctls []corev1.Sysctl) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		pod := ensureSecurityContext(g.pod, constants.POD)
		pod.Template.Spec.SecurityContext.Sysctls = sysctls
		return pod
	}
}

// Volumes sets the lists of volumes to the pod
func (g *Generator) Volumes(volumes []corev1.Volume) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		if volumes == nil {
			return g.pod
		}
		g.pod.Template.Spec.Volumes = volumes
		return g.pod
	}
}

// ReadOnly sets the read-only flag of the container
// will be deprecated in the future when psp is dropped
func (g *Generator) ReadOnly(allow bool, contType string) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		pod := ensureSecurityContext(g.pod, contType)
		switch contType {
		case constants.CONTAINER:
			pod.Template.Spec.Containers[0].SecurityContext.ReadOnlyRootFilesystem = pointer.Bool(allow)
		case constants.INITCONTAINER:
			pod.Template.Spec.InitContainers[0].SecurityContext.ReadOnlyRootFilesystem = pointer.Bool(allow)
		case constants.EPHEMERALCONTAINER:
			pod.Template.Spec.EphemeralContainers[0].EphemeralContainerCommon.SecurityContext.ReadOnlyRootFilesystem = pointer.Bool(allow)
		}
		return pod
	}
}

// FsGroup sets the fsGroup of the container
func (g *Generator) FsGroup(fsGroup int64) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		pod := ensureSecurityContext(g.pod, constants.POD)
		pod.Template.Spec.SecurityContext.FSGroup = pointer.Int64(fsGroup)
		return pod
	}
}

func (g *Generator) SupplementalGroups(group []int64) func() *corev1.PodTemplate {
	return func() *corev1.PodTemplate {
		pod := ensureSecurityContext(g.pod, constants.POD)
		pod.Template.Spec.SecurityContext.SupplementalGroups = group
		return pod
	}
}

func ensureSecurityContext(pod *corev1.PodTemplate, contType string) *corev1.PodTemplate {
	pod = pod.DeepCopy()
	switch contType {
	case constants.POD:
		if pod.Template.Spec.SecurityContext == nil {
			pod.Template.Spec.SecurityContext = &corev1.PodSecurityContext{}
		}
	case constants.CONTAINER:
		for i := range pod.Template.Spec.Containers {
			if pod.Template.Spec.Containers[i].SecurityContext == nil {
				pod.Template.Spec.Containers[i].SecurityContext = &corev1.SecurityContext{}
			}
		}
	case constants.INITCONTAINER:
		for i := range pod.Template.Spec.InitContainers {
			if pod.Template.Spec.InitContainers[i].SecurityContext == nil {
				pod.Template.Spec.InitContainers[i].SecurityContext = &corev1.SecurityContext{}
			}
		}
	case constants.EPHEMERALCONTAINER:
		pod.Template.Spec.EphemeralContainers[0].EphemeralContainerCommon.SecurityContext = &corev1.SecurityContext{}

	}
	return pod
}
