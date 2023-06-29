package camunda

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/camunda/zeebe/clients/go/v8/pkg/pb"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
	"github.com/ugrgnc/camunda-microservices-orchestration/bootstrap"
)

var (
	WorkflowChannel = make(chan struct{})
)

var (
	client zbc.Client
)

func Connect(cfg *bootstrap.CamundaConfig) {
	client, _ = createCamundaConnection(cfg)
}

func CreateJobWorkers() {

	go subscribeCreateOrderProcess()
	go subscribePaymentProcess()
	go subscribeDeliveryOrderProcess()
	go subscribeUpdateInventoryProcess()
	go subscribeSendNotificationProcess()
}

func StartWorkflowInstance() *pb.CreateProcessInstanceResponse {

	// create workflow instance variables
	variables := make(map[string]interface{})
	variables["basket_id"] = rand.Intn((9999 - 3333) + 3333)
	variables["customer_id"] = rand.Intn((29999 - 9333) + 3333)

	// create a workflow instance create request
	request, err := client.NewCreateInstanceCommand().BPMNProcessId("microservice-orchestration-process").LatestVersion().VariablesFromMap(variables)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	// create a workflow process instance
	msg, err := request.Send(ctx)
	if err != nil {
		panic(err)
	}

	// show the result of the process instance command
	fmt.Println(msg.String())

	return msg
}
