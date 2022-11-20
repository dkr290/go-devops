package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// create custom type for budget category
type budgetCategory string

type transaction struct {
	id       int
	payee    string
	spent    float64
	category budgetCategory
}

const (
	autoFuel   budgetCategory = "fuel"
	food       budgetCategory = "food"
	mortgage   budgetCategory = "mortgage"
	repairs    budgetCategory = "repairs"
	insurance  budgetCategory = "insurance"
	utilities  budgetCategory = "utilities"
	retirement budgetCategory = "retirement"
)

// create a custom error when budget category is not found
var (
	ErrBudgetCategoryNotFound = errors.New("the budget category was not found")
)

func main() {

	bankFile := flag.String("c", "", "location of the bank transaction csv file")
	logFile := flag.String("l", "", "location fo the log file")

	flag.Parse()

	if *bankFile == "" {
		fmt.Println("The banlFile is required")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *logFile == "" {
		fmt.Println("The log file is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	_, err := os.Stat(*bankFile)
	if os.IsNotExist(err) {
		fmt.Println("Bankfile does not exist:", *bankFile)
		os.Exit(1)
	}
	_, err = os.Stat(*logFile)
	if !os.IsNotExist(err) {
		os.Remove(*logFile)

	}
	csvFile, err := os.Open(*bankFile)
	if err != nil {
		fmt.Println("Error opening file: ", *bankFile)
		os.Exit(1)
	}

	trx := parseBankFile(csvFile, *logFile)

	fmt.Println()
	for _, trx := range trx {
		fmt.Printf("%v\n", trx)
	}

}

func writeErrorToLog(msg string, err error, data string, logfile string) {

	if msg != "" && err != nil && data != "" {
		msg += "\n" + err.Error() + "\nData: " + data + "\n\n"
		if err := ioutil.WriteFile("log.log", []byte(msg), 0644); err != nil {
			log.Fatalln("Cannot write to the log file")
		}

	}

}

func parseBankFile(bankTransactions io.Reader, logFile string) []transaction {

	r := csv.NewReader(bankTransactions)
	trsx := []transaction{}

	header := true

	for {
		trx := transaction{}
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		if !header {

			for idx, value := range record {

				switch idx {
				case 0:
					value = strings.TrimSpace(value)
					trx.id, err = strconv.Atoi(value)
					if err != nil {
						fmt.Println("Error converting")
					}

				case 1:
					value = strings.TrimSpace(value)
					trx.payee = value

				case 2:
					value = strings.TrimSpace(value)
					trx.spent, err = strconv.ParseFloat(value, 64)
					if err != nil {
						fmt.Println("Error converting")
					}

				case 3:
					trx.category, err = convertToBudgetCategory(value)
					if err != nil {
						s := strings.Join(record, ", ")
						writeErrorToLog("error converting csv category column -", err, s, logFile)
					}
				}

			}
			trsx = append(trsx, trx)

		}
		header = false
	}
	return trsx
}

func convertToBudgetCategory(value string) (budgetCategory, error) {

	value = strings.TrimSpace(strings.ToLower(value))

	switch value {
	case "fuel", "gas":
		return autoFuel, nil
	case "food":
		return food, nil
	case "mortgage":
		return mortgage, nil
	case "repairs":
		return repairs, nil

	case "car insurance", "life insurance":
		return insurance, nil
	case "utilities":
		return utilities, nil
	default:
		return "", ErrBudgetCategoryNotFound

	}

}
