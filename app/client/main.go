package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewLazyClient(client.Options{
		HostPort: "localhost:7233",
	})
	if err != nil {
		log.Fatalln("error create workflow client", err)
	}

	wo := client.StartWorkflowOptions{
		ID: "hello_world_muchlis_test",
		/*
			TaskQueue must be same with taskQueue name in worker
			example code in worker => w := worker.New(c, "hello-world", worker.Options{})
		*/
		TaskQueue: "hello-world",
	}

	// we, err := c.ExecuteWorkflow(context.Background(), wo, mworkflow.MyWorkflow, "Muchlis")
	/*
		MyWorkflow got from named workflow registered in worker
		example code
		w.RegisterWorkflowWithOptions(
		myworkflow.MyWorkflow,
		workflow.RegisterOptions{
			Name:                          "MyWorkflow",
			DisableAlreadyRegisteredCheck: true,
		})
	*/
	we, err := c.ExecuteWorkflow(context.Background(), wo, "MyWorkflow", "Muchlis")
	if err != nil {
		log.Fatalln("error execute workflow", err)
	}

	// do syncronously wait for the workflow completion
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatal("unable get workflow result", err)
	}

	log.Println("workflow result: ", result)
}
