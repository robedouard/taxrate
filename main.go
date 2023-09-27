package main

import "fmt"

// It is not a good idea to use floats to calculate currencies - will correct later

var standardDeduction float64 = 14600.00

// Federal tax rates
var FedRate10 float64 = .10
var FedRate12 float64 = .12
var FedRate22 float64 = .22
var FedRate24 float64 = .24
var FedRate32 float64 = .32
var FedRate35 float64 = .35
var FedRate37 float64 = .37

// Single tax brackets
var Bracket1 float64 = 9525.00
var Bracket2 float64 = 38700.00
var Bracket3 float64 = 82500.00
var Bracket4 float64 = 157500.00
var Bracket5 float64 = 200000.00
var Bracket6 float64 = 500000.00

// Married tax brackets
var MBracket1 float64 = 19050.00
var MBracket2 float64 = 77400.00
var MBracket3 float64 = 165000.00
var MBracket4 float64 = 315000.00
var MBracket5 float64 = 400000.00
var MBracket6 float64 = 600000.00

var salary float64

// ask for gross pay
func askSalary() {
	fmt.Println("Enter your gross annual pay")
	fmt.Scan(&salary)
	fmt.Println("You have entered", salary, "as your anual salary")
	fmt.Println("The standard deduction is currently 14,600.00. Your taxable income is ", salary-standardDeduction)
}

var grossPay float64 = salary - standardDeduction

var tax float64

func computeTax() {

	if grossPay <= Bracket1 {
		tax = FedRate10 * grossPay
		fmt.Printf("%v", tax)
	} else if grossPay >= Bracket1 && grossPay <= Bracket2 {
		tax = FedRate10*Bracket1 + FedRate12*(grossPay-Bracket1)
		fmt.Printf("%v", tax)
	} else if grossPay <= Bracket2 && grossPay <= Bracket3 {
		tax = FedRate10*Bracket1 + FedRate12*(grossPay-Bracket1) + FedRate22*(grossPay-Bracket2)
		fmt.Printf("%v", tax)
	} else if grossPay <= Bracket3 && grossPay <= Bracket4 {
		tax = FedRate10*Bracket1 + FedRate12*(grossPay-Bracket1) + FedRate22*(grossPay-Bracket2) + FedRate24*(grossPay-Bracket3)
		fmt.Printf("%v", tax)
	} else if grossPay <= Bracket4 && grossPay <= Bracket5 {
		tax = FedRate10*Bracket1 + FedRate12*(grossPay-Bracket1) + FedRate22*(grossPay-Bracket2) + FedRate24*(grossPay-Bracket3) + FedRate32*(grossPay-Bracket4)
		fmt.Printf("%v", tax)
	} else if grossPay <= Bracket5 && grossPay <= Bracket6 {
		tax = FedRate10*Bracket1 + FedRate12*(grossPay-Bracket1) + FedRate22*(grossPay-Bracket2) + FedRate24*(grossPay-Bracket3) + FedRate32*(grossPay-Bracket4) + FedRate35*(grossPay-Bracket5)
		fmt.Printf("%v", tax)
	} else if grossPay >= Bracket6 {
		tax = FedRate10*Bracket1 + FedRate12*(grossPay-Bracket1) + FedRate22*(grossPay-Bracket2) + FedRate24*(grossPay-Bracket3) + FedRate32*(grossPay-Bracket4) + FedRate35*(grossPay-Bracket5) + FedRate37*(grossPay-Bracket6)
		fmt.Printf("%v", tax)
	}
}

func main() {
	askSalary()
	computeTax()
}
