package generator

import (
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	// base pod is the base that precedent changes will made changes on
	basepod = &corev1.PodTemplate{
		Template: corev1.PodTemplateSpec{
			Spec: corev1.PodSpec{
				InitContainers: []corev1.Container{{Name: "init-cont", Image: "registry.k8s.io/pause"}},
				Containers:     []corev1.Container{{Name: "cont", Image: "registry.k8s.io/pause"}},
			},
		},
	}
)

func GetBasePod() *corev1.PodTemplate {
	return basepod
}

type Generator struct {
	pod *corev1.PodTemplate
}

func NewGenerator() *Generator {
	return &Generator{
		pod: basepod.DeepCopy(),
	}
}

func (g *Generator) SetPod(pod *corev1.PodTemplate) {
	g.pod = pod
}

func (g *Generator) GetPod() *corev1.PodTemplate {
	return g.pod
}

// Overlay takes a base pod and a list of functions that modify the pod
// if there is conflict then the later wins
func (g *Generator) Overlay(modifiers []func() *corev1.PodTemplate) *corev1.PodTemplate {
	g.SetPod(basepod)
	modified := g.pod.DeepCopy()
	for _, fn := range modifiers {
		modified = fn()
	}
	g.pod = modified
	return modified
}

// Below are a series of functions that wrap the pods using the specified controllers in the target namespace

func ReplicationControllerWrapper(obj *corev1.PodTemplate, ns, name string) ([]byte, error) {
	rc := &corev1.ReplicationController{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name + "-rc",
			Namespace: ns,
		},
		Spec: corev1.ReplicationControllerSpec{
			Replicas: pointer.Int32(1),
			Template: &corev1.PodTemplateSpec{
				ObjectMeta: obj.Template.ObjectMeta,
				Spec:       obj.Template.Spec,
			},
		},
	}
	content, err := runtime.Encode(scheme.Codecs.LegacyCodec(corev1.SchemeGroupVersion), rc)
	if err != nil {
		return nil, err
	}
	yamlData, err := yaml.JSONToYAML(content)
	if err != nil {
		return nil, err
	}
	return yamlData, nil
}

func ReplicaSetWrapper(obj *corev1.PodTemplate, ns, name string) ([]byte, error) {
	rs := &appsv1.ReplicaSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:        name + "rs",
			Namespace:   ns,
			Annotations: obj.Annotations,
		},
		Spec: appsv1.ReplicaSetSpec{
			Replicas: pointer.Int32(1),
			Template: obj.Template,
		},
	}
	content, err := runtime.Encode(scheme.Codecs.LegacyCodec(appsv1.SchemeGroupVersion), rs)
	if err != nil {
		return nil, err
	}
	yamlData, err := yaml.JSONToYAML(content)
	if err != nil {
		return nil, err
	}
	return yamlData, nil
}

func DeploymentWrapper(obj *corev1.PodTemplate, ns, name string) ([]byte, error) {
	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name + "-deploy",
			Namespace: ns,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32(1),
			Template: obj.Template,
		},
	}
	content, err := runtime.Encode(scheme.Codecs.LegacyCodec(appsv1.SchemeGroupVersion), deploy)
	if err != nil {
		return nil, err
	}
	yamlData, err := yaml.JSONToYAML(content)
	if err != nil {
		return nil, err
	}
	return yamlData, nil
}

func StatefulSetWrapper(obj *corev1.PodTemplate, ns, name string) ([]byte, error) {
	ss := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:        name + "-ss",
			Namespace:   ns,
			Annotations: obj.Annotations,
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: pointer.Int32(1),
			Template: obj.Template,
		},
	}
	content, err := runtime.Encode(scheme.Codecs.LegacyCodec(corev1.SchemeGroupVersion), ss)
	if err != nil {
		return nil, err
	}
	yamlData, err := yaml.JSONToYAML(content)
	if err != nil {
		return nil, err
	}
	return yamlData, nil
}

func DaemonsetWrapper(obj *corev1.PodTemplate, ns, name string) ([]byte, error) {
	ds := &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name + "-ds",
			Namespace: ns,
		},
		Spec: appsv1.DaemonSetSpec{
			Template: obj.Template,
		},
	}
	content, err := runtime.Encode(scheme.Codecs.LegacyCodec(appsv1.SchemeGroupVersion), ds)
	if err != nil {
		return nil, err
	}
	yamlData, err := yaml.JSONToYAML(content)
	if err != nil {
		return nil, err
	}
	return yamlData, nil
}

func CronJobWrapper(obj *corev1.PodTemplate, ns, name string) ([]byte, error) {
	cj := &batchv1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name + "-cj",
			Namespace: ns,
		},
		Spec: batchv1.CronJobSpec{
			JobTemplate: batchv1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: obj.Template,
				},
			},
		},
	}
	content, err := runtime.Encode(scheme.Codecs.LegacyCodec(batchv1.SchemeGroupVersion), cj)
	if err != nil {
		return nil, err
	}
	yamlData, err := yaml.JSONToYAML(content)
	if err != nil {
		return nil, err
	}
	return yamlData, nil
}

func JobWrapper(obj *corev1.PodTemplate, ns, name string) ([]byte, error) {
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name + "-job",
			Namespace: ns,
		},
		Spec: batchv1.JobSpec{
			Template: obj.Template,
		},
	}
	content, err := runtime.Encode(scheme.Codecs.LegacyCodec(batchv1.SchemeGroupVersion), job)
	if err != nil {
		return nil, err
	}
	yamlData, err := yaml.JSONToYAML(content)
	if err != nil {
		return nil, err
	}
	return yamlData, nil
}

func PodWrapper(obj *corev1.PodTemplate, ns, name string) ([]byte, error) {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:        name + "-pod",
			Namespace:   ns,
			Annotations: obj.Annotations,
		},
		Spec: obj.Template.Spec,
	}
	content, err := runtime.Encode(scheme.Codecs.LegacyCodec(corev1.SchemeGroupVersion), pod)
	if err != nil {
		return nil, err
	}
	yamlData, err := yaml.JSONToYAML(content)
	if err != nil {
		return nil, err
	}
	return yamlData, nil
}
