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

	list     = kingpin.Command("list", "Print out the reminders in a list")
	listName = list.Arg("list name", "Name of the list whose contents to show").Required().String()

	purge     = kingpin.Command("purge", "Delete all completed reminders (from a specified list")
	purgeList = purge.Arg("list", "The list whose reminders to clear, defaults to all").Default("all-lists").String()
)

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("0.1").Author("Boris Barath")
	kingpin.CommandLine.Help = "A simple command line interface for Apple Reminders"

	switch kingpin.Parse() {
	case "add":
		remindersinterface.AddNewReminder(*addList, *addText)
	case "lists":
		remindersinterface.GetReminderListNames()
	case "list":
		remindersinterface.GetRemindersFromList(*listName)
	case "purge":
		remindersinterface.Purge(*purgeList)
	default:
		fmt.Println("Error: Please enter a command (add/lists)")
	}
}
