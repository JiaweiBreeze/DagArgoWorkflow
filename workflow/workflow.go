package workflow

import (
	"context"

	wfv1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

type ArgoClient struct {
	clientset versioned.Interface
}

func NewArgoClient(kubeconfig string) (*ArgoClient, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &ArgoClient{clientset: clientset}, nil
}

func (c *ArgoClient) CreateWorkflow(workflow *wfv1.Workflow) error {
	wfClient := c.clientset.ArgoprojV1alpha1().Workflows("default")
	_, err := wfClient.Create(context.TODO(), workflow, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}
