package main

import (
	"fmt"
	"go_reminders/src/remindersinterface"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	add     = kingpin.Command("add", "Add a new reminder.").Default()
	addList = add.Flag("list", "List into which to insert the reminder").Default("Reminders").Short('l').String()
	addText = add.Arg("text", "Text of the reminder to create").Required().String()

	lists = kingpin.Command("lists", "Print out the available lists of reminders")
)

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("0.1").Author("Boris Barath")
	kingpin.CommandLine.Help = "A simple command line interface for Apple Reminders"

	switch kingpin.Parse() {
	case "add":
		remindersinterface.AddNewReminder(*addList, *addText)
	case "lists":
		remindersinterface.GetReminderListNames()
	default:
		fmt.Println("Error: Please enter a command (add/lists)")
	}
}
