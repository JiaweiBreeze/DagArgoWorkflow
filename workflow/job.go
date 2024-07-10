package workflow

import (
	"fmt"

	wfv1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
)

type Job struct {
	Name  string
	Image string
	Cmd   []string
}

func NewDAGWorkflow(name string, jobs []Job) *wfv1.Workflow {
	dagTasks := make([]wfv1.DAGTask, len(jobs))
	for i, job := range jobs {
		dagTasks[i] = wfv1.DAGTask{
			Name:     job.Name,
			Template: job.Name,
		}
	}

	templates := make([]wfv1.Template, len(jobs))
	for i, job := range jobs {
		templates[i] = wfv1.Template{
			Name: job.Name,
			Container: &wfv1.Container{
				Image:   job.Image,
				Command: job.Cmd,
			},
		}
	}

	return &wfv1.Workflow{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: fmt.Sprintf("%s-", name),
		},
		Spec: wfv1.WorkflowSpec{
			Entrypoint: "main",
			Templates: []wfv1.Template{
				{
					Name: "main",
					DAG: &wfv1.DAGTemplate{
						Tasks: dagTasks,
					},
				},
			},
		},
	}
}
