package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TimeComponent struct {
	hr, min, sec int
}

type TimeParseError struct {
	Type    string
	Message string
	FixTip  string
	Input   string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%s %s %s", t.Type, t.Message, t.Input)
}

func ParseInput(s string) (*TimeComponent, error) {
	if s[2] != ':' || s[5] != ':' {
		return &TimeComponent{}, &TimeParseError{Type: "Separator error", Message: "Please use : separator between numbers", Input: s}
	}
	if len(s) != 8 {
		return &TimeComponent{}, &TimeParseError{Type: "Invalid length", Message: "Input length is incorrect", Input: s}
	}
	parts := strings.Split(s, ":")
	if len(parts) != 3 {
		return &TimeComponent{}, &TimeParseError{Type: "Invalid part length", Message: "Incorrect length of input", Input: s}
	}
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return &TimeComponent{}, &TimeParseError{Type: "Invalid input", Message: "Error converting hours", Input: parts[0]}
	}
	mins, err := strconv.Atoi(parts[1])
	if err != nil {
		return &TimeComponent{}, &TimeParseError{Type: "Invalid input", Message: "Error converting minutes", Input: parts[1]}
	}
	secs, err := strconv.Atoi(parts[2])
	if err != nil {
		return &TimeComponent{}, &TimeParseError{Type: "Invalid input", Message: "Error converting seconds", Input: parts[2]}
	}
	return &TimeComponent{hr: hours, min: mins, sec: secs}, nil
}

func main() {
	testcase := "14:01:33"
	parts, err := ParseInput(testcase)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(parts)
	}

	//--Summary:
	//  Create a function that can parse time strings into component values.
	//
	//--Requirements:
	//* The function must parse a string into a struct containing:
	//  - Hour, minute, and second integer components
	//* If parsing fails, then a descriptive error must be returned
	//* Write some unit tests to check your work
	//  - Run tests with `go test ./exercise/errors`
	//
	//--Notes:

	//* Example time string: 14:07:33
	//* Use the `strings` package from stdlib to get time components
	//* Use the `strconv` package from stdlib to convert strings to ints
	//* Use the `errors` package to generate errors
	//* Bonus: create test cases and use testing package

	// Empty string
	// Invalid separators (e.g., "14-07-33" instead of "14:07:33")
	// Too many/few components (e.g., "14:07" or "14:07:33:45")
	// Non-numeric components
	// Out-of-range values (e.g., "25:70:80")
	// Leading/trailing whitespace
	// Negative numbers

}
