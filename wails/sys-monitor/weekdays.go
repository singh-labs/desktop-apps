package main

type Weekday string

const (
	Sunday    Weekday = "Sunday"
	Monday    Weekday = "Monday"
	Tuesday   Weekday = "Tuesday"
	Wednesday Weekday = "Wednesday"
	Thursday  Weekday = "Thursday"
	Friday    Weekday = "Friday"
	Saturday  Weekday = "Saturday"
)

var AllWeekdays = []struct {
	Value  Weekday
	TSName string
}{
	{Sunday, "SUNDAY"},
	{Monday, "MONDAY"},
	{Tuesday, "TUESDAY"},
	{Wednesday, "WEDNESDAY"},
	{Thursday, "THURSDAY"},
	{Friday, "FRIDAY"},
	{Saturday, "SATURDAY"},
}
