package camunda

import (
	"context"
	"log"
	"math/rand"

	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
)

func subscribeDeliveryOrderProcess() {

	log.Println("for Delivery-Order-Process subscribed")
	client.NewJobWorker().JobType("delivery-order-process").Handler(deliveryOrderProcessHandler).Open()
}

func deliveryOrderProcessHandler(client worker.JobClient, job entities.Job) {

	jobKey := job.GetKey()

	// log.Println("Complete job", jobKey, "of type", job.Type)

	// errorRequest := client.NewThrowErrorCommand().JobKey(jobKey).ErrorCode("material-invalid")
	// ctx := context.Background()
	// _, err := errorRequest.Send(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// return

	variables, err := job.GetVariablesAsMap()
	if err != nil {
		// failed to handle job as we require the variables
		failJob(client, job)
		return
	}

	// do the delivery
	variables["delivery_id"] = rand.Intn((555 - 111) + 111)
	request, err := client.NewCompleteJobCommand().JobKey(jobKey).VariablesFromMap(variables)
	if err != nil {
		// failed to set the updated variables
		failJob(client, job)
		return
	}

	log.Println("Complete job", jobKey, "of type", job.Type)
	log.Println("Deliver-Order -> basket_id :", variables["basket_id"])
	log.Println("Deliver-Order -> customer_id", variables["customer_id"])
	log.Println("Deliver-Order -> order_id ", variables["order_id"])
	log.Println("Deliver-Order -> payment_id ", variables["payment_id"])

	ctx := context.Background()
	_, err = request.Send(ctx)
	if err != nil {
		panic(err)
	}

	log.Println("Deliver-Order", jobKey, "successfully completed!")
}
