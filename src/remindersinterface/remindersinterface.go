package remindersinterface

import (
	"fmt"
	"strings"

	"github.com/andybrewer/mack"
)

func GetReminderListNames() {
	listCommand := `
	get name of lists
	`
	listNames, _ := mack.Tell("Reminders", listCommand)
	fmt.Println(listNames)
}

func GetRemindersFromList(listName string) {
	cmdTemplate := `
	set myList to list "%s"
	tell myList
		set names to name of reminders whose completed is false 
		set listTexts to ""
		set iterator to 0
		repeat with reminderName in names 
			set listTexts to listTexts & reminderName & "$next"
			set iterator to iterator + 1
		end repeat
		listTexts
	end tell
	`
	cmd := fmt.Sprintf(cmdTemplate, listName)
	reminderNames, _ := mack.Tell("Reminders", cmd)
	for index, reminderText := range strings.Split(reminderNames, "$next") {
		if reminderText != "" {
			fmt.Println(index, reminderText)
		}
	}
}

func Purge(listName string) {
	cmd := ""
	if listName == "all-lists" {
		cmd = `
		set myLists to name of every list
		repeat with currList in myLists
			tell list currList
				delete (every reminder whose completed is true)
			end tell
		end repeat
		`
	} else {
		cmdTemplate := `
		set myList to list "%s"
		tell myList
			delete (every reminder whose completed is true)
		end tell
		`
		cmd = fmt.Sprintf(cmdTemplate, listName)
	}
	mack.Tell("Reminders", cmd)
}

func AddNewReminder(listName string, reminderText string) {
	cmdTemplate := `
	set myList to list "%s" 
	tell myList
		make new reminder at end with properties {name:"%s"}
	end tell`
	cmd := fmt.Sprintf(cmdTemplate, listName, reminderText)
	mack.Tell("Reminders", cmd)
}
