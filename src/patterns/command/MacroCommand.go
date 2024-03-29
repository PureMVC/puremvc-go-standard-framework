//
//  MacroCommand.go
//  PureMVC Go Standard
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package command

import (
	"github.com/puremvc/puremvc-go-standard-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-standard-framework/src/patterns/facade"
)

/*
MacroCommand A base ICommand implementation that executes other ICommands.

A MacroCommand maintains a list of
ICommand Class references called SubCommands.

When execute is called, the MacroCommand
retrieves ICommands by executing funcs and then calls
execute on each of its SubCommands turn.
Each SubCommand will be passed a reference to the original
INotification that was passed to the MacroCommand's
execute method.

Unlike SimpleCommand, your subclass
should not override execute, but instead, should
override the initializeMacroCommand method,
calling addSubCommand once for each SubCommand
to be executed.
*/
type MacroCommand struct {
	facade.Notifier
	SubCommands []func() interfaces.ICommand
}

/*
InitializeMacroCommand Initialize the MacroCommand.

In your subclass, override this method to
initialize the MacroCommand's *SubCommand*
list with func references like
this:

	// Initialize MyMacroCommand
	func (self *MacroCommandTestCommand) InitializeMacroCommand() {
	  self.AddSubCommand(func() interfaces.ICommand { return &FirstCommand{} })
	  self.AddSubCommand(func() interfaces.ICommand { return &SecondCommand{} })
	  self.AddSubCommand(func() interfaces.ICommand { return &ThirdCommand{} })
	}

Note that SubCommands may be any func returning ICommand
implementor, MacroCommands or SimpleCommands are both acceptable.
*/
func (self *MacroCommand) InitializeMacroCommand() {

}

/*
AddSubCommand Add a SubCommand.

The SubCommands will be called in First In/First Out (FIFO)
order.

- parameter factory: reference that returns ICommand.
*/
func (self *MacroCommand) AddSubCommand(factory func() interfaces.ICommand) {
	self.SubCommands = append(self.SubCommands, factory)
}

/*
Execute this MacroCommand's SubCommands.

The SubCommands will be called in First In/First Out (FIFO)
order.

- parameter notification: the INotification object to be passsed to each SubCommand.
*/
func (self *MacroCommand) Execute(notification interfaces.INotification) {
	self.InitializeMacroCommand()
	for len(self.SubCommands) > 0 {
		factory := self.SubCommands[0]
		self.SubCommands = self.SubCommands[1:]

		commandInstance := factory()
		commandInstance.InitializeNotifier()
		commandInstance.Execute(notification)
	}
}
