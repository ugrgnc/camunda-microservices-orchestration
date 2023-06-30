package camunda

import (
	"context"
	"log"

	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
)

func subscribeSendNotificationProcess() {

	log.Println("For Send-Notificaion-Process subscribed")
	client.NewJobWorker().JobType("send-notification-process").Handler(sendNotificationProcessHandler).Open()
}

func sendNotificationProcessHandler(client worker.JobClient, job entities.Job) {

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
	log.Println("Send-Notification -> basket_id :", variables["basket_id"])
	log.Println("Send-Notification -> customer_id", variables["customer_id"])
	log.Println("Send-Notification -> order_id ", variables["order_id"])
	log.Println("Send-Notification -> payment_id ", variables["payment_id"])
	log.Println("Send-Notification -> delivery_id", variables["delivery_id"])
	log.Println("Send-Notification -> shipment_id", variables["shipment_id"])

	ctx := context.Background()
	_, err = request.Send(ctx)
	if err != nil {
		panic(err)
	}

	log.Println("Send-Notification", jobKey, "successfully completed!")
}
