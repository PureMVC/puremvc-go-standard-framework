//
//  Facade.go
//  PureMVC Go Standard
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package facade

import (
	"github.com/puremvc/puremvc-go-standard-framework/src/core/controller"
	"github.com/puremvc/puremvc-go-standard-framework/src/core/model"
	"github.com/puremvc/puremvc-go-standard-framework/src/core/view"
	"github.com/puremvc/puremvc-go-standard-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-standard-framework/src/patterns/observer"
	"sync"
)

/*
A base Singleton IFacade implementation.
*/
type Facade struct {
	controller interfaces.IController // Reference to the Controller
	model      interfaces.IModel      // Reference to the Model
	view       interfaces.IView       // Reference to the View
}

var instance interfaces.IFacade    // The Singleton Facade instance.
var instanceMutex = sync.RWMutex{} // instanceMutex for the instance

/*
  Facade Singleton Factory method

  - parameter facadeFunc: reference that returns IFacade

  - returns: the Singleton instance of the IFacade
*/
func GetInstance(facadeFunc func() interfaces.IFacade) interfaces.IFacade {
	instanceMutex.Lock()
	defer instanceMutex.Unlock()

	if instance == nil {
		instance = facadeFunc()
		instance.InitializeFacade()
	}
	return instance
}

/*
  Initialize the Singleton Facade instance.

  Called automatically by the GetInstance. Override in your
  subclass to do any subclass specific initializations. Be
  sure to call self.Facade.initializeFacade(), though.
*/
func (self *Facade) InitializeFacade() {
	self.InitializeModel()
	self.InitializeController()
	self.InitializeView()
}

/*
  Initialize the Controller.

  Called by the initializeFacade method.
  Override this method in your subclass of Facade
  if one or both of the following are true:

  * You wish to initialize a different IController.

  * You have Commands to register with the Controller at startup.

  If you don't want to initialize a different IController,
  call self.Facade.initializeController() at the beginning of your
  method, then register Commands.
*/
func (self *Facade) InitializeController() {
	self.controller = controller.GetInstance(func() interfaces.IController { return &controller.Controller{} })
}

/*
  Initialize the Model.

  Called by the initializeFacade method.
  Override this method in your subclass of Facade
  if one or both of the following are true:

  * You wish to initialize a different IModel.

  * You have Proxys to register with the Model that do not retrieve a reference to the Facade at construction time.

  If you don't want to initialize a different IModel,
  call self.Facade.initializeModel() at the beginning of your
  method, then register Proxys.

  Note: This method is rarely overridden; in practice you are more
  likely to use a Command to create and register Proxys
  with the Model, since Proxys with mutable data will likely
  need to send INotifications and thus will likely want to fetch a reference to
  the Facade during their construction.
*/
func (self *Facade) InitializeModel() {
	self.model = model.GetInstance(func() interfaces.IModel { return &model.Model{} })
}

/*
  Initialize the View.

  Called by the initializeFacade method.
  Override this method in your subclass of Facade
  if one or both of the following are true:

  * You wish to initialize a different IView.

  * You have Observers to register with the View

  If you don't want to initialize a different IView,
  call self.Facade.initializeView() at the beginning of your
  method, then register IMediator instances.

  Note: This method is rarely overridden; in practice you are more
  likely to use a Command to create and register Mediators
  with the View, since IMediator instances will need to send
  INotifications and thus will likely want to fetch a reference
  to the Facade during their construction.
*/
func (self *Facade) InitializeView() {
	self.view = view.GetInstance(func() interfaces.IView { return &view.View{} })
}

/*
  Register an ICommand with the Controller by Notification name.

  - parameter notificationName: the name of the INotification to associate the ICommand with

  - parameter commandFunc: reference that returns ICommand
*/
func (self *Facade) RegisterCommand(notificationName string, commandFunc func() interfaces.ICommand) {
	self.controller.RegisterCommand(notificationName, commandFunc)
}

/*
  Remove a previously registered ICommand to INotification mapping from the Controller.

  - parameter notificationName: the name of the INotification to remove the ICommand mapping for
*/
func (self *Facade) RemoveCommand(notificationName string) {
	self.controller.RemoveCommand(notificationName)
}

/*
  Check if a Command is registered for a given Notification

  - parameter notificationName:

  - returns: whether a Command is currently registered for the given notificationName.
*/
func (self *Facade) HasCommand(notificationName string) bool {
	return self.controller.HasCommand(notificationName)
}

/*
  Register an IProxy with the Model by name.

  - parameter proxy: the IProxy instance to be registered with the Model.
*/
func (self *Facade) RegisterProxy(proxy interfaces.IProxy) {
	self.model.RegisterProxy(proxy)
}

/*
  Retrieve an IProxy from the Model by name.

  - parameter proxyName: the name of the proxy to be retrieved.

  - returns: the IProxy instance previously registered with the given proxyName.
*/
func (self *Facade) RetrieveProxy(proxyName string) interfaces.IProxy {
	return self.model.RetrieveProxy(proxyName)
}

/*
  Remove an IProxy from the Model by name.

  - parameter proxyName: the IProxy to remove from the Model.

  - returns: the IProxy that was removed from the Model
*/
func (self *Facade) RemoveProxy(proxyName string) interfaces.IProxy {
	return self.model.RemoveProxy(proxyName)
}

/*
  Check if a Proxy is registered

  - parameter proxyName:

  - returns: whether a Proxy is currently registered with the given proxyName.
*/
func (self *Facade) HasProxy(proxyName string) bool {
	return self.model.HasProxy(proxyName)
}

/*
  Register a IMediator with the View.

  - parameter mediator: a reference to the IMediator
*/
func (self *Facade) RegisterMediator(mediator interfaces.IMediator) {
	self.view.RegisterMediator(mediator)
}

/*
  Retrieve an IMediator from the View.

  - parameter mediatorName:

  - returns: the IMediator previously registered with the given mediatorName.
*/
func (self *Facade) RetrieveMediator(mediatorName string) interfaces.IMediator {
	return self.view.RetrieveMediator(mediatorName)
}

/*
  Remove an IMediator from the View.

  - parameter mediatorName: name of the IMediator to be removed.

  - returns: the IMediator that was removed from the View
*/
func (self *Facade) RemoveMediator(mediatorName string) interfaces.IMediator {
	return self.view.RemoveMediator(mediatorName)
}

/*
  Check if a Mediator is registered or not

  - parameter mediatorName:

  - returns: whether a Mediator is registered with the given mediatorName.
*/
func (self *Facade) HasMediator(mediatorName string) bool {
	return self.view.HasMediator(mediatorName)
}

/*
  Create and send an INotification.

  Keeps us from having to construct new notification
  instances in our implementation code.

  - parameter notificationName: the name of the notiification to send

  - parameter body: the body of the notification (optional)

  - parameter _type: the type of the notification
*/
func (self *Facade) SendNotification(notificationName string, body interface{}, _type string) {
	self.NotifyObservers(observer.NewNotification(notificationName, body, _type))
}

/*
  Notify Observers.

  This method is left mostly for backward
  compatibility, and to allow you to send custom
  notification classes using the facade.

  Usually you should just call sendNotification
  and pass the parameters, never having to
  construct the notification yourself.

  - parameter notification: the INotification to have the View notify Observers of.
*/
func (self *Facade) NotifyObservers(notification interfaces.INotification) {
	self.view.NotifyObservers(notification)
}

/*
  Set the Singleton key for this facade instance.

  Not called directly, but instead from the
  GetInstance when it is is invoked.
*/
func (self *Facade) InitializeNotifier() {

}