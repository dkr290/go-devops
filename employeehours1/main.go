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

var workWeek = make(map[Weekday]int)

func (d *Developer) LogHours(day Weekday, hours int) {

	workWeek[day] = hours

}

func (d *Developer) HoursWorked() int {
	total := 0

	for _, v := range workWeek {
		total += v
	}
	return total
}

func (d *Developer) PayDay() (int, bool) {
	if d.HoursWorked() > 40 {
		hoursOver := d.HoursWorked() - 40
		overtime := hoursOver * 2 * d.HourlyRate
		regularPay := d.HoursWorked() * d.HourlyRate
		return regularPay + overtime, true
	}
	return d.HoursWorked() * d.HourlyRate, false

}

func (d *Developer) PayDetails() {

	fmt.Println("Hours work this week", d.HoursWorked())
	weeklyPay, overtime := d.PayDay()

	fmt.Println("Pay for this week", weeklyPay)
	fmt.Println("Is overtime pay", overtime)
}

func main() {

	d := Developer{
		Individual: Employee{
			Id:        3234,
			FirstName: "John",
			LastName:  "Michr",
		},
		HourlyRate: 10,
	}

	d.LogHours(Monday, 8)
	d.LogHours(Tuesday, 10)
	d.LogHours(Wednesday, 10)
	d.LogHours(Thursday, 10)
	d.LogHours(Friday, 6)
	d.LogHours(Saturday, 8)

	fmt.Println("Hours worked on Monday: ", workWeek[Monday])
	fmt.Println("Hours Worked on Tuesday", workWeek[Tuesday])
	fmt.Println("Hours worked this week", d.HoursWorked())

	x := nonLoggedHours()
	fmt.Println("Tracking hours worked thus for today", x(2))
	fmt.Println("Tracking hours worked thus for today", x(3))
	fmt.Println("Tracking hours worked thus for today", x(5))

	d.PayDetails()
}

func nonLoggedHours() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}
