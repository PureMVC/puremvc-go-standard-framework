//
//  ViewTestMediator6.go
//  PureMVC Go Standard
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package view

import (
	"github.com/puremvc/puremvc-go-standard-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-standard-framework/src/patterns/mediator"
)

const ViewTestMediator6_NAME = "ViewTestMediator6" // The Mediator base name

/*
ViewTestMediator6 A Mediator class used by ViewTest.
*/
type ViewTestMediator6 struct {
	mediator.Mediator
}

func (mediator *ViewTestMediator6) ListNotificationInterests() []string {
	return []string{VIEWTEST_NOTE6}
}

func (mediator *ViewTestMediator6) HandleNotification(notification interfaces.INotification) {
	//temp implementation until facade is developed
	mediator.Notifier.Facade.RemoveMediator(mediator.GetMediatorName())
	//var view2 = view.GetInstance("ViewTestKey11", func() interfaces.IView {
	//	return &view.View{Key: "ViewTestKey11"}
	//})
	//view2.RemoveMediator(mediator.GetMediatorName())
}

func (mediator *ViewTestMediator6) OnRemove() {
	mediator.ViewComponent.(*Data).counter++
}
