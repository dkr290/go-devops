package main

import (
	"errors"
	"fmt"

	"github.com/dkr290/go-devops/interfaces3/payroll"
)

var employeeReview = make(map[string]interface{})

func init() {

	fmt.Println("Welcome to the employee pay and performance review")
	fmt.Println("**************************************************")
}

func init() {
	fmt.Println("Initizlizing the variables")
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependibility"] = "Unsatisfatory"
}

func main() {

	d1 := payroll.Developer{
		Individual: payroll.Employee{
			Id:        33923,
			FirstName: "John",
			LastName:  "Doe",
		},
		HourlyRate:        34.8,
		HoursWorkedInYear: 200,
		Review:            employeeReview,
	}

	m1 := payroll.Manager{
		Individual: payroll.Employee{
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

func PayDetails(p payroll.Payer) {

	fn, pay := p.Pay()
	fmt.Println(fn, "has got payed", pay, "year")
}

func convertReviewToInt(str string) (int, error) {

	switch str {
	case "Exellent":
		return 5, nil
	case "Good":
		return 4, nil

	case "Fair":
		return 3, nil

	case "Poor":
		return 2, nil

	case "Unsatisfactory":
		return 1, nil

	default:

		return 0, errors.New("There is no review from the list please use appropariate one " + str)
	}
}
