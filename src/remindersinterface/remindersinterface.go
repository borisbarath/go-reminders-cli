package remindersinterface

import (
	"fmt"

	"github.com/andybrewer/mack"
)

func AddNewReminder(listName string, reminderText string) {
	cmd := `
	set myList to list "%s" 
	tell myList
		make new reminder at end with properties {name:"%s"}
	end tell`
	addCommand := fmt.Sprintf(cmd, listName, reminderText)
	mack.Tell("Reminders", addCommand)
}
