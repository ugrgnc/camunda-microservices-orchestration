package camunda

import (
	"context"
	"log"
	"math/rand"

	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
)

func subscribeCreateOrderProcess() {

	log.Println("for Create-Order-Process subscribed")
	client.NewJobWorker().JobType("create-order-process").Handler(createOrderProcessHandler).Open()

	// jobworker := client.NewJobWorker()
	// job := jobworker.JobType("create-order-process")
	// handler := job.Handler(createOrderProcessHandler)
	// handler.Concurrency(5) // concurrency
	// //handler.Timeout(time.Second * 3)
	// handler.RequestTimeout(time.Second * 5)
	// handler.Name("create-order-process-consumer")
	// handler.Open()
	//worker.Close()

}

func createOrderProcessHandler(client worker.JobClient, job entities.Job) {

	//time.Sleep(time.Second * 5)
	jobKey := job.GetKey()

	variables, err := job.GetVariablesAsMap()
	if err != nil {
		// failed to handle job as we require the variables
		failJob(client, job)
		return
	}

	// create the order
	variables["order_id"] = rand.Intn((333 - 111) + 111)
	request, err := client.NewCompleteJobCommand().JobKey(jobKey).VariablesFromMap(variables)
	if err != nil {
		// failed to set the updated variables
		failJob(client, job)
		return
	}

	log.Println("Complete job", jobKey, "of type", job.Type)
	log.Println("Create-Order -> basket_id ", variables["basket_id"])
	log.Println("Create-Order -> customer_id ", variables["customer_id"])

	ctx := context.Background()
	_, err = request.Send(ctx)
	if err != nil {
		panic(err)
	}

	log.Println("Create-Order job", jobKey, "successfully completed!")
}
