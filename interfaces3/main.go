package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

type Manager struct {
	Individual    Employee
	Salary        float64
	CommisionRate float64
}

type Payer interface {
	Pay() (string, float64)
}

func (d *Developer) FullName() string {

	return d.Individual.FirstName + " " + d.Individual.LastName
}

func (m *Manager) FullName() string {
	return m.Individual.FirstName + " " + m.Individual.LastName
}

func (d *Developer) Pay() (string, float64) {
	fullName := d.FullName()
	return fullName, d.HourlyRate * d.HoursWorkedInYear

}

func (m *Manager) Pay() (string, float64) {
	fullName := m.FullName()
	return fullName, m.Salary + (m.Salary * m.CommisionRate)

}

func main() {

	d1 := Developer{
		Individual: Employee{
			Id:        33923,
			FirstName: "John",
			LastName:  "Doe",
		},
		HourlyRate:        34.8,
		HoursWorkedInYear: 200,
	}

	m1 := Manager{
		Individual: Employee{
			Id:        232342,
			FirstName: "Kapil",
			LastName:  "Blabla",
		},
		Salary:        3400.45,
		CommisionRate: 2.8,
	}
	PayDetails(&d1)
	PayDetails(&m1)

}

func PayDetails(p Payer) {

	fn, pay := p.Pay()
	fmt.Println(fn, "has got payed", pay, "year")
}
