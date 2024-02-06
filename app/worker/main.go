package main

import (
	"log"

	"github.com/muchlist/workflow_tempo/business/myactivity"
	"github.com/muchlist/workflow_tempo/business/myworkflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func main() {
	c, err := client.NewLazyClient(client.Options{
		HostPort: "localhost:7233",
	})
	if err != nil {
		log.Fatalln("error dial workflow", err)
	}
	defer c.Close()

	w := worker.New(c, "hello-world", worker.Options{})

	// w.RegisterWorkflow(workflow.MyWorkflow) -- simple version, need to shared code
	w.RegisterWorkflowWithOptions(
		myworkflow.MyWorkflow,
		workflow.RegisterOptions{
			Name:                          "MyWorkflow",
			DisableAlreadyRegisteredCheck: true,
		})
	w.RegisterActivity(myactivity.Activity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start worker", err)
	}
}
