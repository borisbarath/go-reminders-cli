package remindersinterface

import (
	"fmt"
	"go_reminders/src/applescript"
	"strings"

	"github.com/andybrewer/mack"
)

func GetReminderListNames() {
	listNames, _ := mack.Tell("Reminders", applescript.ReminderLists)
	fmt.Println(listNames)
}

func GetRemindersFromList(listName string) {
	cmd := fmt.Sprintf(applescript.ListRemindersFrom, listName)
	reminderNames, _ := applescript.TellReminders(cmd)
	for index, reminderText := range strings.Split(reminderNames, "$next") {
		if reminderText != "" {
			fmt.Println(index, reminderText)
		}
	}
}

func Purge(listName string) {
	cmd := ""
	if listName == "all-lists" {
		cmd = applescript.DeleteAllDoneReminders
	} else {
		cmd = fmt.Sprintf(applescript.DeleteDoneRemindersFromList, listName)
	}
	applescript.TellReminders(cmd)
}

func Done(doneList string, doneReminder string, doneIndex int) {
	// TODO(borisbarath) match text to reminder text to mark as done
	cmd := fmt.Sprintf(applescript.MarkReminderAsDone, doneList, doneIndex, doneReminder)
	applescript.TellReminders(cmd)
}

func AddNewReminder(listName string, reminderText string) {

	cmd := fmt.Sprintf(applescript.AddReminderToList, listName, reminderText)
	applescript.TellReminders(cmd)
}
