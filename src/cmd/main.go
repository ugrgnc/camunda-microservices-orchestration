package main

import (
	"fmt"
	"net/http"

	"github.com/ugrgnc/camunda-microservices-orchestration/bootstrap"
	"github.com/ugrgnc/camunda-microservices-orchestration/pkg/camunda"
)

func main() {

	fmt.Printf("Camunda-Microservices-Orchestration\n")

	app := bootstrap.App()
	camunda.Connect(app.CamundaConfig)
	camunda.CreateJobWorkers()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.HandleFunc("/start-workflow", func(w http.ResponseWriter, r *http.Request) {
		instance := camunda.StartWorkflowInstance()
		fmt.Fprintf(w, "The workflow instance created -> ID : %d\n", instance.ProcessInstanceKey)
	})

	http.ListenAndServe(":8080", nil)
}
