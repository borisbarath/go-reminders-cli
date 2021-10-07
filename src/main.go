package main

import (
	"fmt"
	"go_reminders/src/remindersinterface"

	flag "github.com/spf13/pflag"
)

func main() {

	flag.Usage = func() {
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
	default:
		fmt.Println("Error: Please enter a command (add)")
	}

}

func addReminder(reminderList string) {
	text := flag.Arg(1) // The only argument (that is not a flag option) is the file location (CSV file)
	if len(text) == 0 {
		fmt.Println("Please enter the reminder text")
	} else {
		remindersinterface.AddNewReminder(reminderList, text)
	}
}
