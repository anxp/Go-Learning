package main

import (
	"fmt"
	"strings"
)

func main() {
	const (
		Monday    string = "Monday"
		Tuesday   string = "Tuesday"
		Wednesday string = "Wednesday"
		Thursday  string = "Thursday"
		Friday    string = "Friday"
		Saturday  string = "Saturday"
	)

	var day string

	fmt.Print("Please, enter a day to get a tip for this day: ")
	_, err := fmt.Scan(&day)

	if err != nil {
		fmt.Println(err)
		return
	}

	//Make first letter of user entered word uppercase:
	day = strings.Title(day)

	switch day {
	case Monday:
		fmt.Println("Laugh on Monday, laugh for danger.")
	case Tuesday:
		fmt.Println("Laugh on Tuesday, kiss a stranger.")
	case Wednesday:
		fmt.Println("Laugh on Wednesday, laugh for a letter.")
	case Thursday:
		fmt.Println("Laugh on Thursday, something better.")
	case Friday:
		fmt.Println("Laugh on Friday, laugh for sorrow.")
	case Saturday:
		fmt.Println("Laugh on Saturday, joy tomorrow.")
	default:
		fmt.Println("No information for that day.")
	}
}
