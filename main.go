package main

import (
	"dag/config"
	"dag/workflow"
	"log"
)

func main() {
	config := config.LoadConfig()

	client, err := workflow.NewArgoClient(config.Kubeconfig)
	if err != nil {
		log.Fatalf("Failed to create Argo client: %v", err)
	}

	jobs := []workflow.Job{
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

	dagWorkflow := workflow.NewDAGWorkflow("example-dag", jobs)

	err = client.CreateWorkflow(dagWorkflow)
	if err != nil {
		log.Fatalf("Failed to create workflow: %v", err)
	}

	log.Println("Workflow created successfully")
}
