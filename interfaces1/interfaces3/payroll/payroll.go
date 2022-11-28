package payroll

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
