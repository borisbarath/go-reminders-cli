package main

import (
	"fmt"
	"go_reminders/src/remindersinterface"

	flag "github.com/spf13/pflag"
)

func main() {

	flag.Usage = func() {
		fmt.Println("usage: reminders add \"Reminder text\" [flags]")
		fmt.Println("       reminders lists")
		flag.PrintDefaults()
	}

	reminderList := flag.StringP("list", "l", "Reminders", "Specify a list on which to perform the operation")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Please enter a command [add]")
	}

	command := flag.Arg(0)
	switch command {
	case "add":
		addReminder(*reminderList)
	case "lists":
		listReminderLists()
	default:
		fmt.Println("Error: Please enter a command (add/list)")
	}
}

func listReminderLists() {
	remindersinterface.GetReminderListNames()
}

func addReminder(reminderList string) {
	text := flag.Arg(1)
	if len(text) == 0 {
		fmt.Println("Please enter the reminder text")
	} else {
		remindersinterface.AddNewReminder(reminderList, text)
	}
}
