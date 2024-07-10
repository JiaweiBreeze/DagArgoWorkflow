package workflow

import (
	"context"
	"testing"

	"github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned/fake"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNewArgoClient(t *testing.T) {
	// This function remains unchanged.
	_, err := NewArgoClient("fake-kubeconfig")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestCreateWorkflow(t *testing.T) {
	clientset := fake.NewSimpleClientset()
	argoClient := &ArgoClient{clientset: clientset}

	jobs := []Job{
		{
			Name:  "job1",
			Image: "alpine:latest",
			Cmd:   []string{"echo", "Hello from Job 1"},
		},
		{
			Name:  "job2",
			Image: "alpine:latest",
			Cmd:   []string{"echo", "Hello from Job 2"},
		},
	}

	dagWorkflow := NewDAGWorkflow("example-dag", jobs)

	err := argoClient.CreateWorkflow(dagWorkflow)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	wf, err := clientset.ArgoprojV1alpha1().Workflows("default").Get(context.TODO(), dagWorkflow.Name, metav1.GetOptions{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if wf.Name != dagWorkflow.Name {
		t.Errorf("Expected workflow name %s, got %s", dagWorkflow.Name, wf.Name)
	}
}
