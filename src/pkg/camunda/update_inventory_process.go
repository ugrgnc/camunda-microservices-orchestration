package camunda

import (
	"context"
	"log"

	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
)

func subscribeUpdateInventoryProcess() {

	log.Println("For Update-Inventory-Process subscribed")
	client.NewJobWorker().JobType("update-inventory-process").Handler(updateInventoryProcessHandler).Open()
}

func updateInventoryProcessHandler(client worker.JobClient, job entities.Job) {

	jobKey := job.GetKey()

	variables, err := job.GetVariablesAsMap()
	if err != nil {
		// failed to handle job as we require the variables
		failJob(client, job)
		return
	}

	request, err := client.NewCompleteJobCommand().JobKey(jobKey).VariablesFromMap(variables)
	if err != nil {
		// failed to set the updated variables
		failJob(client, job)
		return
	}

	log.Println("Complete job", jobKey, "of type", job.Type)
	log.Println("Update-Inventory -> basket_id :", variables["basket_id"])
	log.Println("Update-Inventory -> customer_id", variables["customer_id"])
	log.Println("Update-Inventory -> order_id ", variables["order_id"])
	log.Println("Update-Inventory -> payment_id ", variables["payment_id"])
	log.Println("Update-Inventory -> delivery_id", variables["delivery_id"])

	ctx := context.Background()
	_, err = request.Send(ctx)
	if err != nil {
		panic(err)
	}

	log.Println("Update-Inventory", jobKey, "successfully completed!")
}
