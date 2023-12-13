//
//  ViewTestNote.go
//  PureMVC Go Standard
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package view

import (
	"github.com/puremvc/puremvc-go-standard-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-standard-framework/src/patterns/observer"
)

const ViewTestNote_NAME string = "ViewTestNote"

type ViewTestNote struct {
}

func ViewTestNoteNew(body interface{}) interfaces.INotification {
	return observer.NewNotification(ViewTestNote_NAME, body, "")
}
