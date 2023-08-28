package fixtures

import (
	"github.com/chenliu1993/podsecurity-check/pkg/generator"
	corev1 "k8s.io/api/core/v1"
)

const (
	// below defines a series of security function names

	apparmor                          = "apparmor"
	capabilitiesBaseline              = "capabilitiesBaseline"
	ephemeralCapabilitiesBaseline     = "ephemeralCapabilitiesBaseline"
	capabilitiesRestricted            = "capabilitiesRestricted"
	ephemeralCapabilitiesRestricted   = "ephemeralCapabilitiesRestricted"
	fsGroup                           = "fsGroup"
	hostpid                           = "hostpid"
	hostipc                           = "hostipc"
	hostnetwork                       = "hostnetwork"
	hostports                         = "hostports"
	ephemeralHostports                = "ephemeralhostports"
	privileged                        = "privileged"
	ephemeralPrivileged               = "ephemeral-privileged"
	allowPrivilegeEscalation          = "allowPrivilegeEscalation"
	ephemeralAllowPrivilegeEscalation = "ephemeral-allowPrivilegeEscalation"
	procMount                         = "procMount"
	ephemeralprocMount                = "ephemeralprocMount"
	readOnly                          = "readOnly"
	ephemeralReadOnly                 = "ephemeralReadOnly"
	runAsNonRoot                      = "runAsNonRoot"
	ephemeralrunAsNonRoot             = "ephemeralrunAsNonRoot"
	runAsUser                         = "runAsUser"
	ephemeralrunAsUser                = "ephemeralrunAsUser"
	seccomp                           = "seccomp"
	ephemeralSeccomp                  = "ephemeralSeccomp"
	seccompRestricted                 = "seccompRestricted"
	ephemeralSeccompRestricted        = "ephemeralSeccompRestricted"
	selinux                           = "selinux"
	ephemeralSELinux                  = "ephemeralSELinux"
	supplementalgroups                = "supplementalgroups"
	sysctls                           = "sysctls"
	volumesBaseline                   = "volumesBaseline"
	volumesRestricted                 = "volumesRestricted"
)

var (
	baselineSet = []string{
		hostpid,
		hostipc,
		hostnetwork,
		privileged,
		capabilitiesBaseline,
		volumesBaseline,
		hostports,
		apparmor,
		selinux,
		procMount,
		seccomp,
		sysctls,
	}
	restrictedSet = []string{
		hostpid,
		hostipc,
		hostnetwork,
		privileged,
		capabilitiesRestricted,
		volumesRestricted,
		hostports,
		apparmor,
		selinux,
		procMount,
		seccompRestricted,
		sysctls,
		allowPrivilegeEscalation,
		runAsNonRoot,
		runAsUser,
	}

	testSets = map[string][]string{
		"baseline":   baselineSet,
		"restricted": restrictedSet,
	}
)

var (
	// below defines all the cases different level cares about
	AddCapBaselinPassSet = []corev1.Capability{
		"AUDIT_WRITE", "CHOWN", "DAC_OVERRIDE", "FOWNER", "FSETID", "KILL", "MKNOD", "NET_BIND_SERVICE", "SETFCAP", "SETGID", "SETPCAP", "SETUID", "SYS_CHROOT",
	}
	AddCapBaselinFailSet = []corev1.Capability{
		"NET_ADMIN",
	}
	AddCapRestrictedPassSet = []corev1.Capability{
		"NET_BIND_SERVICE",
	}
	AddCapRestrictedFailSet = []corev1.Capability{
		"NET_ADMIN",
	}
	DropCapRestrictedPassSet = []corev1.Capability{
		"ALL",
	}
	DropCapRestrictedFailSet = []corev1.Capability{
		"NET_ADMIN",
	}

	HostPortBaselinePass = []corev1.ContainerPort{
		{
			HostPort: 0,
		},
		{
			ContainerPort: 8080,
		},
	}

	HostPortBaselineFail = []corev1.ContainerPort{
		{
			HostPort: 12345,
		},
		{
			ContainerPort: 8080,
			HostPort:      8080,
		},
	}

	SysctlsBaselinePass = []corev1.Sysctl{
		{Name: "kernel.shm_rmid_forced", Value: "0"},
		{Name: "net.ipv4.ip_local_port_range", Value: "1024 65535"},
		{Name: "net.ipv4.tcp_syncookies", Value: "0"},
		{Name: "net.ipv4.ping_group_range", Value: "1 0"},
		{Name: "net.ipv4.ip_unprivileged_port_start", Value: "1024"},
	}
	SysctlsBaselineFail = []corev1.Sysctl{
		{Name: "other", Value: "other"},
	}

	VolumesBaselinePass = []corev1.Volume{
		{Name: "hostpathvolume0", VolumeSource: corev1.VolumeSource{}},
	}
	VolumesBaselineFail = []corev1.Volume{
		{Name: "hostpathvolume1", VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "test",
			},
		}},
	}
	VolumesRestrictedPass = []corev1.Volume{
		{Name: "volume0", VolumeSource: corev1.VolumeSource{}}, // implicit empty dir
		{Name: "volume1", VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}}},
		{Name: "volume2", VolumeSource: corev1.VolumeSource{Secret: &corev1.SecretVolumeSource{SecretName: "test"}}},
		{Name: "volume3", VolumeSource: corev1.VolumeSource{PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{ClaimName: "test"}}},
		{Name: "volume4", VolumeSource: corev1.VolumeSource{DownwardAPI: &corev1.DownwardAPIVolumeSource{Items: []corev1.DownwardAPIVolumeFile{{Path: "labels", FieldRef: &corev1.ObjectFieldSelector{FieldPath: "metadata.labels"}}}}}},
		{Name: "volume5", VolumeSource: corev1.VolumeSource{ConfigMap: &corev1.ConfigMapVolumeSource{LocalObjectReference: corev1.LocalObjectReference{Name: "test"}}}},
		{Name: "volume6", VolumeSource: corev1.VolumeSource{Projected: &corev1.ProjectedVolumeSource{Sources: []corev1.VolumeProjection{}}}},
	}

	// Jusr random choose some type
	VolumesRestrictedFail = []corev1.Volume{
		{Name: "volume1", VolumeSource: corev1.VolumeSource{NFS: &corev1.NFSVolumeSource{Server: "test", Path: "/test"}}},
		{Name: "volume2", VolumeSource: corev1.VolumeSource{GitRepo: &corev1.GitRepoVolumeSource{Repository: "github.com/kubernetes/kubernetes"}}},
	}

	ContainerT = corev1.SELinuxOptions{
		Type: "container_t",
	}
	ContainerKvmT = corev1.SELinuxOptions{
		Type: "container_kvm_t",
	}
	ContainerInitT = corev1.SELinuxOptions{
		Type: "container_init_t",
	}
	ContainerInitKvmT = corev1.SELinuxOptions{
		Type: "container_init_kvm_t",
	}
	EmptyUserRoleOpts = corev1.SELinuxOptions{
		User: "",
		Role: "",
	}
	NonemptyUserRoleOpts = corev1.SELinuxOptions{
		User: "system_u",
		Role: "system_r",
	}
	DefaultProcMount    = corev1.DefaultProcMount
	UmaskedProcMount    = corev1.UnmaskedProcMount
	FsGroupBaselinePass = int64(1000)
	FsGroupBaselineFail = int64(0)

	SupplementalGroupsBaselineFail = []int64{0}
	SupplementalGroupsBaselinePass = []int64{1000}

	DefaultSeccompProfile    = corev1.SeccompProfile{Type: corev1.SeccompProfileTypeRuntimeDefault}
	LocalhostProfile         = corev1.SeccompProfile{Type: corev1.SeccompProfileTypeLocalhost}
	UnconfinedSeccompProfile = corev1.SeccompProfile{Type: corev1.SeccompProfileTypeUnconfined}
)

var (
	g *generator.Generator

	serverVersion = "1.21" // This is the last version that pod security policy is supported

	fixturesMap map[string][]func() Case
	testdataDir = "testdata"
	clustersDir = "clusters"
)

type Case struct {
	Name      string
	Namespace string
	PassSet   []*corev1.PodTemplate
	FailSet   []*corev1.PodTemplate
}

type Standard struct {
	Mode    string `yaml:"mode"`
	Version string `yaml:"version"`
}

type SecurityPolicy struct {
	Namespace string    `yaml:"namespace"`
	Enforce   *Standard `yaml:"enforce"`
	Audit     *Standard `yaml:"audity"`
	Warn      *Standard `yaml:"warn"`
}

type SecurityPolicies struct {
	SecurityPolicies []SecurityPolicy `yaml:"securityPolicies"`
}
