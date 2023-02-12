package main

import (
	"context"
	"fmt"
	"hello-world-temporal/app"

	"log"

	"go.temporal.io/sdk/client"
)

func main() {

	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// The WorkflowOptions specify the options for starting the workflow execution.
	// Here we set the WorkflowID which must be unique in a given domain.
	workflowOptions := client.StartWorkflowOptions{
		ID:        "hello_world_workflow_id",
		TaskQueue: app.GreetingTaskQueue,
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, app.GreetingWorkflow, "Temporal")

	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}

	printResult(greeting, we.GetID(), we.GetRunID())
}

func printResult(greeting string, workflowID string, runID string) {
	fmt.Println("WorkflowID:", workflowID)
	fmt.Println("RunID:", runID)
	fmt.Println("Result:", greeting)
}
