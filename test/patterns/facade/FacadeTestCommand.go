//
//  FacadeTestCommand.go
//  PureMVC Go Standard
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package facade

import (
	"github.com/puremvc/puremvc-go-standard-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-standard-framework/src/patterns/command"
)

/*
A SimpleCommand subclass used by FacadeTest.
*/
type FacadeTestCommand struct {
	command.SimpleCommand
}

/*
  Fabricate a result by multiplying the input by 2

  - parameter note: the Notification carrying the FacadeTestVO
*/
func (facade *FacadeTestCommand) Execute(notification interfaces.INotification) {
	var vo = notification.Body().(*FacadeTestVO)

	// Fabricate a Result
	vo.Result = 2 * vo.Input
}
