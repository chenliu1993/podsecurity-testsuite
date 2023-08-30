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

const (
	annotationKeyPod             = "seccomp.security.alpha.kubernetes.io/pod"
	annotationKeyContainerPrefix = "container.seccomp.security.alpha.kubernetes.io/"
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
	ErrMsg    string
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
