package controllers

import (
	"encoding/json"
	"golang-learn/examples/52_nsq/common/nsq"
	"golang-learn/examples/52_nsq/interactors"
	"golang-learn/examples/52_nsq/domain"
)

type Controller struct {
	Interactor interactors.Interactor
	Consumer   nsq.INSQConsumer
}

// Starts listening from nsq channel
func (controller *Controller) Start(controlChannel chan int) error {
	controller.Consumer.SetHandler(controller, 10)

	return controller.Consumer.Start(controlChannel)
}

// Starts listening from nsq channel
func (controller *Controller) Stop() error {
	return controller.Consumer.Stop()
}

// INSQHandler implementations
// Handle the NSQ message and parse it to be managed by interactors
func (controller *Controller) HandleMessage(data []byte) error {
	something := &domain.Something{}

	if err := json.Unmarshal(data, something); err != nil {
		return err
	}

	return controller.Interactor.DoSomething()
}
