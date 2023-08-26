package fixtures

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/chenliu1993/podsecurity-check/pkg/files"
	"github.com/chenliu1993/podsecurity-check/pkg/generator"
	"gopkg.in/yaml.v2"
	corev1 "k8s.io/api/core/v1"
)

var (
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

	testdataDir = "testdata"
	clustersDir = "clusters"
)

var (
	g *generator.Generator

	serverVersion = "1.21" // This is the last version that pod security policy is supported

	fixturesMap map[string][]func() Case
)

type Case struct {
	Name      string
	Namespace string
	PassSet   []*corev1.PodTemplate
	FailSet   []*corev1.PodTemplate
}

type SecurityPolicy struct {
	Namespace string   `yaml:"namespace"`
	Policies  []string `yaml:"policies"`
}

type SecurityPolicies struct {
	SecurityPolicies []SecurityPolicy `yaml:"securityPolicies"`
}

func init() {
	// init the pod template generator
	g = generator.NewGenerator()

	fixturesMap = make(map[string][]func() Case)
}

// Caller should make sure the path is absolute to fixtures.
func GenerateCases() error {
	var clusterSecurityPolicies SecurityPolicies
	clustersList, err := files.WalkPath(filepath.Join(clustersDir))
	if err != nil {
		return err
	}
	for _, clusterItem := range clustersList {
		yamlData, err := os.ReadFile(clusterItem)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(yamlData, &clusterSecurityPolicies)
		if err != nil {
			return err
		}
		if err := generateClusterSecurityCases(strings.TrimSuffix(filepath.Base(clusterItem), filepath.Ext(clusterItem)), clusterSecurityPolicies); err != nil {
			return err
		}
	}
	return nil
}

func generateClusterSecurityCases(clusterName string, clusterSecurityPolicies SecurityPolicies) error {
	var cases []Case
	if err := checkExpectedDir(clusterName); err != nil {
		return err
	}
	for _, securitypolicy := range clusterSecurityPolicies.SecurityPolicies {
		ns := securitypolicy.Namespace
		for _, policy := range securitypolicy.Policies {
			for _, f := range fixturesMap[policy] {
				c := f()
				c.Namespace = ns
				cases = append(cases, c)
			}
		}

	}
	for _, tc := range cases {
		for _, template := range tc.PassSet {
			if err := generateCaseFiles(template, tc.Namespace, tc.Name, "pass", clusterName); err != nil {
				return err
			}
		}
		for _, template := range tc.FailSet {
			if err := generateCaseFiles(template, tc.Namespace, tc.Name, "fail", clusterName); err != nil {
				return err
			}
		}
	}
	return nil
}

func generateCaseFiles(podTemplate *corev1.PodTemplate, ns, name, expected string, clusterName string) error {
	content, err := generator.PodWrapper(podTemplate, ns, name)
	if err != nil {
		return err
	}

	err = files.WriteFile(filepath.Join(testdataDir, clusterName, expected, name+"-pod.yaml"), content)
	if err != nil {
		return err
	}
	content, err = generator.DeploymentWrapper(podTemplate, ns, name)
	if err != nil {
		return err
	}
	err = files.WriteFile(filepath.Join(testdataDir, clusterName, expected, name+"-deploy.yaml"), content)
	if err != nil {
		return err
	}
	content, err = generator.ReplicaSetWrapper(podTemplate, ns, name)
	if err != nil {
		return err
	}
	err = files.WriteFile(filepath.Join(testdataDir, clusterName, expected, name+"-rs.yaml"), content)
	if err != nil {
		return err
	}
	content, err = generator.ReplicationControllerWrapper(podTemplate, ns, name)
	if err != nil {
		return err
	}
	err = files.WriteFile(filepath.Join(testdataDir, clusterName, expected, name+"-rc.yaml"), content)
	if err != nil {
		return err
	}
	content, err = generator.JobWrapper(podTemplate, ns, name)
	if err != nil {
		return err
	}
	err = files.WriteFile(filepath.Join(testdataDir, clusterName, expected, name+"-job.yaml"), content)
	if err != nil {
		return err
	}
	content, err = generator.CronJobWrapper(podTemplate, ns, name)
	if err != nil {
		return err
	}
	err = files.WriteFile(filepath.Join(testdataDir, clusterName, expected, name+"-cj.yaml"), content)
	if err != nil {
		return err
	}
	content, err = generator.DaemonsetWrapper(podTemplate, ns, name)
	if err != nil {
		return err
	}
	err = files.WriteFile(filepath.Join(testdataDir, clusterName, expected, name+"-ds.yaml"), content)
	if err != nil {
		return err
	}
	return nil
}

// older returns true if this version ver1 is older than ver2.
func older(ver1, ver2 string) bool {
	if ver1 == "latest" { // Latest is always consider newer, even than future versions.
		return false
	}
	if ver2 == "latest" {
		return true
	}

	// don't care about the major version
	ver1Minor := strings.Split(ver1, ".")[1]
	ver2Minor := strings.Split(ver2, ".")[1]
	return ver1Minor <= ver2Minor
}

func checkExpectedDir(subdir string) error {
	if err := files.CheckDir(filepath.Join(testdataDir, subdir, "pass")); err != nil {
		return err
	}
	if err := files.CheckDir(filepath.Join(testdataDir, subdir, "fail")); err != nil {
		return err
	}
	return nil
}
