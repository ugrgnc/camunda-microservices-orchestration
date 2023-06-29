package camunda

import (
	"context"
	"log"
	"math/rand"

	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
)

// var paymentProcessHandlerChanel = make(chan struct{})

func subscribePaymentProcess() {

	log.Println("For Payment-Process subscribed")
	client.NewJobWorker().JobType("payment-process").Handler(paymentProcessHandler).Open()

	// log.Println("subscribePaymentProcess waiting...")
	// <-paymentProcessHandlerChannel

	// log.Println("subscribePaymentProcess finishing...")
	// jobWorker.Close()
	// jobWorker.AwaitClose()
	// log.Println("subscribePaymentProcess finished")
}

func paymentProcessHandler(client worker.JobClient, job entities.Job) {

	jobKey := job.GetKey()

	variables, err := job.GetVariablesAsMap()
	if err != nil {
		// failed to handle job as we require the variables
		failJob(client, job)
		return
	}

	// do the payment
	variables["payment_id"] = rand.Intn((888 - 222) + 222)
	request, err := client.NewCompleteJobCommand().JobKey(jobKey).VariablesFromMap(variables)
	if err != nil {
		// failed to set the updated variables
		failJob(client, job)
		return
	}

	log.Println("Complete job", jobKey, "of type", job.Type)
	log.Println("Payment-Process -> basket_id :", variables["basket_id"])
	log.Println("Payment-Process -> customer_id", variables["customer_id"])
	log.Println("Payment-Process -> order_id ", variables["order_id"])

	ctx := context.Background()
	_, err = request.Send(ctx)
	if err != nil {
		panic(err)
	}

	log.Println("Payment-Process job successfully completed!")

	// paymentProcessHandlerChanel <- struct{}{}
}
