package applescript

import "github.com/andybrewer/mack"

const (
	ReminderLists = `
	get name of lists
	`

	ListRemindersFrom = `
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

	DeleteAllDoneReminders = `
	set myLists to name of every list
	repeat with currList in myLists
		tell list currList
			delete (every reminder whose completed is true)
		end tell
	end repeat
	`

	DeleteDoneRemindersFromList = `
	set myList to list "%s"
	tell myList
		delete (every reminder whose completed is true)
	end tell
	`

	MarkReminderAsDone = `
	set mylist to list "%s"
	set notCompleted to reminders in myList whose completed is false
	set iterator to 0
	set doneIndex to %d
	set doneText to "%s"
	repeat with currentReminder in notCompleted
		if (iterator equals doneIndex) or (doneText equals name of currentReminder as string) then
			set completion date of currentReminder to current date
		end if
		set iterator to iterator + 1
	end repeat
	`

	AddReminderToList = `
	set myList to list "%s" 
	tell myList
		make new reminder at end with properties {name:"%s"}
	end tell
	`
)

func TellReminders(command string) (string, error) {
	return mack.Tell("Reminders", command)
}
