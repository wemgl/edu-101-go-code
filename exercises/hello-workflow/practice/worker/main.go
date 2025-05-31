package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"hello"
	"log"
)

const TaskQueueName = "greeting-tasks"

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, TaskQueueName, worker.Options{})

	w.RegisterWorkflow(hello.GreetSomeone)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
