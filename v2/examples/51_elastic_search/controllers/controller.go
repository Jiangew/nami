package controllers

import (
	"golang-learn/examples/51_elastic_search/interactors"
	"golang-learn/examples/51_elastic_search/domain"
)

type Controller struct {
	Interactor interactors.Interactor
}

func (controller *Controller) CreateIndex(index string) error {
	return controller.Interactor.CreateIndex(index)
}

func (controller *Controller) Insert(data []domain.Something) error {
	return controller.Interactor.Insert(data)
}
