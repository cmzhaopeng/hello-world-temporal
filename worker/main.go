package main

import (
	"hello-world-temporal/app"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// The worker is a lightweight object that can be used to poll for and process workflows and activities.
	// It should be created once per process.
	w := worker.New(c, app.GreetingTaskQueue, worker.Options{})

	// Workflows are stateful. So you need a type to create instances.
	w.RegisterWorkflow(app.GreetingWorkflow)

	// Activities are stateless and thread safe. So a single instance is used.
	w.RegisterActivity(app.ComposeGreeting)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
