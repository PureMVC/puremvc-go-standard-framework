//
//  Facade_test.go
//  PureMVC Go Standard
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package facade

import (
	"github.com/puremvc/puremvc-go-standard-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-standard-framework/src/patterns/facade"
	"github.com/puremvc/puremvc-go-standard-framework/src/patterns/mediator"
	"github.com/puremvc/puremvc-go-standard-framework/src/patterns/proxy"
	"testing"
)

/*
Test the PureMVC Facade class.
*/

func TestGetInstance(t *testing.T) {
	// Test Factory Method
	var f = facade.GetInstance(func() interfaces.IFacade { return &facade.Facade{} })

	// test assertions
	if f == nil {
		t.Error("Expecting instance not nil")
	}
}

/*
Tests Command registration and execution via the Facade.

This test gets a Singleton Facade instance
and registers the FacadeTestCommand class
to handle 'FacadeTest' Notifications.

It then sends a notification using the Facade.
Success is determined by evaluating
a property on an object placed in the body of
the Notification, which will be modified by the Command.
*/
func TestRegisterCommandAndSendNotification(t *testing.T) {
	// Create the Facade, register the FacadeTestCommand to
	// handle 'FacadeTest' notifications
	var f = facade.GetInstance(func() interfaces.IFacade { return &facade.Facade{} })
	f.RegisterCommand("FacadeTestNote", func() interfaces.ICommand { return &FacadeTestCommand{} })

	// Send notification. The Command associated with the event
	// (FacadeTestCommand) will be invoked, and will multiply
	// the vo.input value by 2 and set the result on vo.result
	var vo = FacadeTestVO{Input: 32}
	f.SendNotification("FacadeTestNote", &vo, "")

	// test assertions
	if vo.Result != 64 {
		t.Error("Expecting vo.result == 64")
	}
}

/*
Tests Command removal via the Facade.

This test gets a Singleton Facade instance
and registers the FacadeTestCommand class
to handle 'FacadeTest' Notifications. Then it removes the command.

It then sends a Notification using the Facade.
Success is determined by evaluating
a property on an object placed in the body of
the Notification, which will NOT be modified by the Command.
*/
func TestRegisterAndRemoveCommandAndSendNotification(t *testing.T) {
	// Create the Facade, register the FacadeTestCommand to
	// handle 'FacadeTest' events
	var f = facade.GetInstance(func() interfaces.IFacade { return &facade.Facade{} })
	f.RegisterCommand("FacadeTestNote", func() interfaces.ICommand { return &FacadeTestCommand{} })
	f.RemoveCommand("FacadeTestNote")

	// Send notification. The Command associated with the event
	// (FacadeTestCommand) will NOT be invoked, and will NOT multiply
	// the vo.input value by 2
	var vo = FacadeTestVO{Input: 32}
	f.SendNotification("FacadeTestNote", &vo, "")

	// test assertions
	if vo.Result == 64 {
		t.Error("Expecting vo.result != 64")
	}
}

/*
Tests the registering and retrieving Model proxies via the Facade.

Tests registerProxy and retrieveProxy in the same test.
These methods cannot currently be tested separately
in any meaningful way other than to show that the
methods do not throw exception when called.
*/
func TestRegisterAndRetrieveProxy(t *testing.T) {
	// register a proxy and retrieve it.
	var f = facade.GetInstance(func() interfaces.IFacade { return &facade.Facade{} })
	f.RegisterProxy(&proxy.Proxy{Name: "colors", Data: []string{"red", "green", "blue"}})
	var p = f.RetrieveProxy("colors").(*proxy.Proxy)

	// retrieve data from proxy
	var data = p.Data.([]string)

	// test assertions
	if data == nil {
		t.Error("Expecting data not nil")
	}
	if len(data) != 3 {
		t.Error("Expecting len(data) == 3")
	}
	if data[0] != "red" {
		t.Error("Expecting data[0] == 'red'")
	}
	if data[1] != "green" {
		t.Error("Expecting data[1] == 'green'")
	}
	if data[2] != "blue" {
		t.Error("Expecting data[2] == 'blue'")
	}
}

/*
Tests the removing Proxies via the Facade.
*/
func TestRegisterAndRemoveProxy(t *testing.T) {
	// register a proxy, remove it, then try to retrieve it
	var f = facade.GetInstance(func() interfaces.IFacade {
		return &facade.Facade{}
	})
	var p interfaces.IProxy = &proxy.Proxy{Name: "sizes", Data: []string{"7", "13", "21"}}
	f.RegisterProxy(p)

	// remove the proxy
	var removedProxy = f.RemoveProxy("sizes")

	// assert that we removed the appropriate proxy
	if removedProxy.GetProxyName() != "sizes" {
		t.Error("Expecting removedProxy.GetProxyName() == 'sizes'")
	}

	// make sure we can no longer retrieve the proxy from the model
	var proxy2 = f.RetrieveProxy("sizes")
	// test assertions
	if proxy2 != nil {
		t.Error("Expecting proxy is nil")
	}
}

/*
Tests registering, retrieving and removing Mediators via the Facade.
*/
func TestRegisterRetrieveAndRemoveMediator(t *testing.T) {
	// register a mediator, remove it, then try to retrieve it
	var f = facade.GetInstance(func() interfaces.IFacade { return &facade.Facade{} })
	f.RegisterMediator(&mediator.Mediator{Name: mediator.NAME, ViewComponent: []string{}})

	// retrieve the mediator
	if f.RetrieveMediator(mediator.NAME) == nil {
		t.Error("Expecting mediator is not nil")
	}

	// remove the mediator
	removedMediator := f.RemoveMediator(mediator.NAME)

	// assert that we have removed the appropriate mediator
	if removedMediator.GetMediatorName() != mediator.NAME {
		t.Error("Expecting removedMediator.GetMediatorName() == Mediator.NAME")
	}

	// assert that the mediator is no longer retrievable
	if f.RetrieveMediator(mediator.NAME) != nil {
		t.Error("Expecting facade.RetrieveMediator( Mediator.NAME ) == nil")
	}
}

/*
Tests the hasProxy Method
*/
func TestHasProxy(t *testing.T) {
	// register a Proxy
	var f = facade.GetInstance(func() interfaces.IFacade { return &facade.Facade{} })
	f.RegisterProxy(&proxy.Proxy{Name: "hasProxyTest", Data: []int{1, 2, 3}})

	// assert that the model.hasProxy method returns true
	// for that proxy name
	if f.HasProxy("hasProxyTest") != true {
		t.Error("Expecting facade.HasProxy('hasProxyTest') == true")
	}
}

/*
Tests the hasMediator Method
*/
func TestHasMediator(t *testing.T) {
	// register a Mediator
	var f = facade.GetInstance(func() interfaces.IFacade { return &facade.Facade{} })
	f.RegisterMediator(&mediator.Mediator{Name: "facadeHasMediatorTest", ViewComponent: []int{}})

	// assert that the facade.hasMediator method returns true
	// for that mediator name
	if f.HasMediator("facadeHasMediatorTest") != true {
		t.Error("Expecting facade.HasMediator('facadeHasMediatorTest') == true")
	}

	f.RemoveMediator("facadeHasMediatorTest")

	// assert that the facade.hasMediator method returns false
	// for that mediator name
	if f.HasMediator("facadeHasMediatorTest") != false {
		t.Error("Expecting facade.HasMediator('facadeHasMediatorTest') == false")
	}
}

/*
Test hasCommand method.
*/
func TestHasCommand(t *testing.T) {
	// register the ControllerTestCommand to handle 'hasCommandTest' notes
	var f = facade.GetInstance(func() interfaces.IFacade {
		return &facade.Facade{}
	})
	f.RegisterCommand("facadeHasCommandTest", func() interfaces.ICommand { return &FacadeTestCommand{} })

	// test that hasCommand returns true for hasCommandTest notifications
	if f.HasCommand("facadeHasCommandTest") != true {
		t.Error("Expecting facade.HasCommand('facadeHasCommandTest') == true")
	}

	// Remove the Command from the Controller
	f.RemoveCommand("facadeHasCommandTest")

	// test that hasCommand returns false for hasCommandTest notifications
	if f.HasCommand("facadeHasCommentTest") != false {
		t.Error("Expecting facade.HasCommand('facadeHasCommandTest') == false")
	}
}
