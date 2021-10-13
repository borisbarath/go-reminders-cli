
# [WIP] go-reminders-cli

go-reminders-cli is a simple Apple Reminders interface written in Go. This is stil very much a WIP

## Usage

Add a new reminder: `add [<flags>] <text>`
```
Flags:
  -l, --list="Reminders"  List into which to insert the reminder
Args:
  <text>  Text of the reminder to create
  ````
Print out the available lists of reminders: `lists`

Print all reminders in a list not marked as done: `list <list name>`

Delete all completed reminders from a list (defaults to all lists): `purge <list name>`

Mark a reminder as done: `done [<flags>] <index of reminder>`
```
Flags:
  -l, --list="Reminders"  List in which to mark reminder as done
  -t, --text=""           The text of the reminder to be marked as done
Args:
  [<index>]  List in which to mark reminder as done
```
