package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	// panic("Please implement the Welcome function")
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	// panic("Please implement the HappyBirthday function")
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	// panic("Please implement the AssignTable function")
	newTable := fmt.Sprintf("%03d", table)

	output := fmt.Sprintf("Welcome to my party, %s!\n", name)
	output += fmt.Sprintf("You have been assigned to table %s. Your table is %s, exactly %.1f meters from here.\n", newTable, direction, distance)
	output += fmt.Sprintf("You will be sitting next to %s.", neighbor)

	return output
}
