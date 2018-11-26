package main

import (
	"fmt"
)

type salaryCalculator interface {
	CalculateSalary() int
}

type permanent struct {
	empID    int
	basicpay int
	pf       int
}

type contract struct {
	empID    int
	basicpay int
}

// CalculateSalary salary of permanent employee is sum of basic pay and pf
func (p permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

// CalculateSalary salary of contract employee is the basic pay alone
func (c contract) CalculateSalary() int {
	return c.basicpay
}

/*
total expense is calculated by iterating though the SalaryCalculator slice and summing
the salaries of the individual employees
*/
func totalExpense(s []salaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {
	pemp1 := permanent{1, 5000, 20}
	pemp2 := permanent{2, 6000, 30}
	cemp1 := contract{3, 3000}
	employees := []salaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
}
