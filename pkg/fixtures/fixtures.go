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
	if !strings.Contains(name, "enforce") {

		content, err := generator.DeploymentWrapper(podTemplate, ns, name)
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
	}
	content, err := generator.PodWrapper(podTemplate, ns, name)
	if err != nil {
		return err
	}
	err = files.WriteFile(filepath.Join(testdataDir, clusterName, expected, name+"-pod.yaml"), content)
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
