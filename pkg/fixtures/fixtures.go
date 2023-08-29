package fixtures

import (
	"path/filepath"
	"strings"

	"github.com/chenliu1993/podsecurity-check/pkg/cli"
	"github.com/chenliu1993/podsecurity-check/pkg/files"
	"github.com/chenliu1993/podsecurity-check/pkg/generator"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
)

func init() {
	// init the pod template generator
	g = generator.NewGenerator()

	fixturesMap = make(map[string][]func() Case)
}

// Caller should make sure the path is absolute to fixtures.
func GenerateCases(namespaces map[string]map[string]string) error {
	clusterSecurityPolicies := convertNamespacesIntoSecurityPolicy(namespaces)
	if len(clusterSecurityPolicies.SecurityPolicies) == 0 {
		return nil
	}

	if err := generateClusterSecurityCases(clusterSecurityPolicies); err != nil {
		return err
	}

	return nil
}

func generateClusterSecurityCases(clusterSecurityPolicies SecurityPolicies) error {
	var cases []Case
	if err := checkExpectedDir(testdataDir); err != nil {
		return err
	}
	for _, securitypolicy := range clusterSecurityPolicies.SecurityPolicies {
		ns := securitypolicy.Namespace
		if securitypolicy.Warn != nil {
			cases = genCasesByMode(cases, securitypolicy.Warn, ns, "warn")

		}
		if securitypolicy.Enforce != nil {
			cases = genCasesByMode(cases, securitypolicy.Enforce, ns, "enforce")

		}
		if securitypolicy.Audit != nil {
			cases = genCasesByMode(cases, securitypolicy.Audit, ns, "audit")
		}
	}
	for _, tc := range cases {
		for _, template := range tc.PassSet {
			if err := generateCaseFiles(template, tc.Namespace, tc.Name, "pass"); err != nil {
				return err
			}
		}
		for _, template := range tc.FailSet {
			if err := generateCaseFiles(template, tc.Namespace, tc.Name, "fail"); err != nil {
				return err
			}
		}
	}
	return nil
}

func generateCaseFiles(podTemplate *corev1.PodTemplate, ns, name, expected string) error {
	if !strings.Contains(name, "enforce") {

		content, err := generator.DeploymentWrapper(podTemplate, ns, name)
		if err != nil {
			return err
		}
		err = files.WriteFile(filepath.Join(testdataDir, expected, expected+"-"+name+"-deploy.yaml"), content)
		if err != nil {
			return err
		}
		/* content, err = generator.ReplicaSetWrapper(podTemplate, ns, name)
		if err != nil {
			return err
		}
		err = files.WriteFile(filepath.Join(testdataDir, expected, expected+"-"+name+"rs.yaml"), content)
		if err != nil {
			return err
		}
		content, err = generator.ReplicationControllerWrapper(podTemplate, ns, name)
		if err != nil {
			return err
		}
		err = files.WriteFile(filepath.Join(testdataDir, expected, expected+"-"+name+"-rc.yaml"), content)
		if err != nil {
			return err
		}
		content, err = generator.JobWrapper(podTemplate, ns, name)
		if err != nil {
			return err
		}
		err = files.WriteFile(filepath.Join(testdataDir, expected, expected+"-"+name+"-job.yaml"), content)
		if err != nil {
			return err
		}
		content, err = generator.CronJobWrapper(podTemplate, ns, name)
		if err != nil {
			return err
		}
		err = files.WriteFile(filepath.Join(testdataDir, expected, expected+"-"+name+"-cj.yaml"), content)
		if err != nil {
			return err
		} */
		content, err = generator.DaemonsetWrapper(podTemplate, ns, name)
		if err != nil {
			return err
		}
		err = files.WriteFile(filepath.Join(testdataDir, expected, expected+"-"+name+"-ds.yaml"), content)
		if err != nil {
			return err
		}
	}
	content, err := generator.PodWrapper(podTemplate, ns, name)
	if err != nil {
		return err
	}
	err = files.WriteFile(filepath.Join(testdataDir, expected, expected+"-"+name+"-pod.yaml"), content)
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

func genCasesByMode(cases []Case, standatd *Standard, ns, level string) []Case {
	for _, securityFocus := range testSets[standatd.Mode] {
		for _, f := range fixturesMap[securityFocus] {
			c := f()
			c.Name = c.Name + "-" + level
			c.Namespace = ns
			cases = append(cases, c)
		}
	}
	return cases
}

func getExpectedResult(name string) string {
	items := strings.Split(name, "-")
	return items[0]
}

// This function should be called for psa
func GetNamespaceWithSecurityLabels() (map[string]map[string]string, error) {
	nsData, err := cli.Kubectl(nil, "get", "namespace", "-o", "yaml")
	if err != nil {
		return nil, err
	}
	var nsList corev1.NamespaceList
	err = yaml.Unmarshal([]byte(nsData), &nsList)
	if err != nil {
		return nil, err
	}
	nsLabels := make(map[string]map[string]string)
	for _, ns := range nsList.Items {
		nsLabels[ns.Name] = getPSALabels(ns.Labels)
	}
	return nsLabels, nil
}

func getPSALabels(labels map[string]string) map[string]string {
	psaLabels := make(map[string]string)
	for k, v := range labels {
		if strings.HasPrefix(k, "pod-security.kubernetes.io/") {
			psaLabels[k] = v
		}
	}
	if len(psaLabels) == 0 {
		return nil
	}
	return psaLabels
}

func getServerVersion() (string, error) {
	data, err := cli.Kubectl(nil, "version")
	if err != nil {
		return "", err
	}
	serverVersionLine := strings.Split(data, "\n")[2]
	return strings.Split(serverVersionLine, " ")[2], nil
}

func convertNamespacesIntoSecurityPolicy(namespaces map[string]map[string]string) SecurityPolicies {
	clusterSecurityPolicies := SecurityPolicies{
		SecurityPolicies: []SecurityPolicy{},
	}
	for ns, labels := range namespaces {
		if labels == nil {
			// This namespace will be skipped because no pod security labels are set.
			continue
		}
		securityPolicy := SecurityPolicy{
			Enforce: &Standard{},
			Audit:   &Standard{},
			Warn:    &Standard{},
		}
		securityPolicy.Namespace = ns
		for k, v := range labels {
			switch strings.TrimPrefix(k, "pod-security.kubernetes.io/") {
			case "enforce":
				securityPolicy.Enforce.Mode = v
			case "audit":
				securityPolicy.Audit.Mode = v
			case "warn":
				securityPolicy.Warn.Mode = v
			case "enforce-version":
				securityPolicy.Enforce.Version = v
			case "audit-version":
				securityPolicy.Audit.Version = v
			case "warn-version":
				securityPolicy.Warn.Version = v
			}
		}
		clusterSecurityPolicies.SecurityPolicies = append(clusterSecurityPolicies.SecurityPolicies, securityPolicy)
	}
	return clusterSecurityPolicies
}
