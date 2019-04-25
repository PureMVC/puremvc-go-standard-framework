//
//  ObserverTest.go
//  PureMVC Go Standard
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package view

import "github.com/puremvc/puremvc-go-standard-framework/src/interfaces"

/**
  A test variable that proves the viewTestMethod was
  invoked by the View.
*/
var ViewTestVar int

type ObserverTest struct {
	NotifyMethod func(note interfaces.INotification)
}

func NotifyTestMethod(note interfaces.INotification) {
	ViewTestVar = note.Body().(int)
}
