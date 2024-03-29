//
//  ViewTestMediator4.go
//  PureMVC Go Standard
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package view

import "github.com/puremvc/puremvc-go-standard-framework/src/patterns/mediator"

const ViewTestMediator4_NAME = "ViewTestMediator4"

/*
ViewTestMediator4 A Mediator class used by ViewTest.
*/
type ViewTestMediator4 struct {
	mediator.Mediator
}

func (mediator *ViewTestMediator4) OnRegister() {
	mediator.ViewComponent.(*Data).onRegisterCalled = true
}

func (mediator *ViewTestMediator4) OnRemove() {
	mediator.ViewComponent.(*Data).onRemoveCalled = true
}
