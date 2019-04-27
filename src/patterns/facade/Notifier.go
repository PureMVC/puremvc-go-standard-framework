//
//  Notifier.go
//  PureMVC Go Standard
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package facade

import "github.com/puremvc/puremvc-go-standard-framework/src/interfaces"

/*
A Base INotifier implementation.

MacroCommand, Command, Mediator and Proxy
all have a need to send Notifications.

The INotifier interface provides a common method called
sendNotification that relieves implementation code of
the necessity to actually construct Notifications.

The Notifier class, which all of the above mentioned classes
extend, provides an initialized reference to the Facade
Singleton, which is required for the convienience method
for sending Notifications, but also eases implementation as these
classes have frequent Facade interactions and usually require
access to the facade anyway.
*/
type Notifier struct {
	Facade interfaces.IFacade
}

/*
  Create and send an INotification.

  Keeps us from having to construct new INotification
  instances in our implementation code.

  - parameter notificationName: the name of the notification to send

  - parameter body: the body of the notification (optional)

  - parameter type: the _type of the notification
*/
func (self *Notifier) SendNotification(notificationName string, body interface{}, _type string) {
	self.Facade.SendNotification(notificationName, body, _type)
}

/*
  Initialize this INotifier instance.

  This is how a Notifier get to Calls to
  sendNotification or to access the
  facade will fail until after this method
  has been called.

  Mediators, Commands or Proxies may override
  this method in order to send notifications
  or access the Facade instance as
  soon as possible. They CANNOT access the facade
  in their constructors, since this method will not
  yet have been called.
*/
func (self *Notifier) InitializeNotifier() {
	self.Facade = GetInstance(func() interfaces.IFacade { return &Facade{} })
}
