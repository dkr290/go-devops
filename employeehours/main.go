package main

import "fmt"

type Weekday int

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d *Developer) LogHours(day Weekday, hours int) {
	d.WorkWeek[day] = hours
}

func (d *Developer) HoursWorked() int {
	total := 0

	for _, v := range d.WorkWeek {
		total += v
	}
	return total
}

func main() {

	d := Developer{
		Individual: Employee{
			Id:        3234,
			FirstName: "John",
			LastName:  "Michr",
		},
		HourlyRate: 280,
	}

	d.LogHours(Monday, 8)
	d.LogHours(Tuesday, 10)

	fmt.Println("Hours worked on Monday: ", d.WorkWeek[Monday])
	fmt.Println("Hours Worked on Tuesday", d.WorkWeek[Tuesday])
	fmt.Println("Hours worked this week", d.HoursWorked())
}
