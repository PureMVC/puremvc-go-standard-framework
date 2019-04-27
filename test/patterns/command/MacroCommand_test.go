//
//  MacroCommand_test.go
//  PureMVC Go Standard
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package command

import (
	"github.com/puremvc/puremvc-go-standard-framework/src/core/controller"
	"github.com/puremvc/puremvc-go-standard-framework/src/core/view"
	"github.com/puremvc/puremvc-go-standard-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-standard-framework/src/patterns/command"
	"github.com/puremvc/puremvc-go-standard-framework/src/patterns/observer"
	"testing"
)

/*
Tests operation of a MacroCommand.

This test creates a new Notification, adding a
MacroCommandTestVO as the body.
It then creates a MacroCommandTestCommand and invokes
its execute method, passing in the
Notification.

The MacroCommandTestCommand has defined an
initializeMacroCommand method, which is
called automatically by its constructor. In this method
the MacroCommandTestCommand adds 2 SubCommands
to itself, MacroCommandTestSub1Command and
MacroCommandTestSub2Command.

The MacroCommandTestVO has 2 result properties,
one is set by MacroCommandTestSub1Command by
multiplying the input property by 2, and the other is set
by MacroCommandTestSub2Command by multiplying
the input property by itself.

Success is determined by evaluating the 2 result properties
on the MacroCommandTestVO that was passed to
the MacroCommandTestCommand on the Notification
body.
*/
func TestMacroCommandExecute(t *testing.T) {
	// Create the VO
	var vo = MacroCommandTestVO{Input: 5}

	// Create the Notification (note)
	var note = observer.NewNotification("MacroCommandTest", &vo, "")

	// Create the SimpleCommand
	var command = MacroCommandTestCommand{MacroCommand: command.MacroCommand{}}
	command.Notifier.InitializeNotifier()

	// Execute the SimpleCommand
	command.Execute(note)

	// test assertions
	if vo.Result1 != 10 {
		t.Error("Expecting vo.Result1 == 10")
	}
	if vo.Result2 != 25 {
		t.Error("Expecting vo.Result2 == 25")
	}
}

/*
  Testing MacroCommand via Controller and notify via View
*/
func TestMacroCommandExecuteViaControllerView(t *testing.T) {
	var controller = controller.GetInstance(func() interfaces.IController { return &controller.Controller{} })
	var view = view.GetInstance(func() interfaces.IView { return &view.View{}})

	controller.RegisterCommand("MacroCommandTestViaControllerView", func() interfaces.ICommand { return &MacroCommandTestCommand{} })

	var vo = MacroCommandTestVO{Input: 5}
	var note = observer.NewNotification("MacroCommandTestViaControllerView", &vo, "")
	view.NotifyObservers(note)

	if vo.Result1 != 10 {
		t.Error("Expecting vo.Result1 == 10")
	}
	if vo.Result2 != 25 {
		t.Error("Expecting vo.Result2 == 25")
	}
}
