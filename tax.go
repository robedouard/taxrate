package main

import "fmt"

// It is not a good idea to use floats to calculate currencies - will correct later
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

var grossPay float64
var maritalStatus string

// ask for gross pay
func askGrossPay() {
	fmt.Println("Enter your gross annual pay")
	fmt.Scan(&grossPay)
	fmt.Println("You have entered", grossPay, "as your anual salary")
}

func askMaritalStatus() {
	fmt.Println("Enter your marital status (S or M) = ")
	fmt.Scan(&maritalStatus)
	fmt.Println(maritalStatus)
}

var tax float64

func computeTax() {
	if maritalStatus == "S" || maritalStatus == "s" {
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
		} else if grossPay >= Bracket6 {
			tax = FedRate10*Bracket1 + FedRate12*(grossPay-Bracket1) + FedRate22*(grossPay-Bracket2) + FedRate24*(grossPay-Bracket3) + FedRate32*(grossPay-Bracket4) + FedRate35*(grossPay-Bracket5) + FedRate37*(grossPay-Bracket6)

			fmt.Printf("%v", tax)
		}
	} else if maritalStatus == "M" || maritalStatus == "m" {
		if grossPay <= MBracket1 {
			tax = FedRate10 * grossPay
			fmt.Printf("%v", tax)
		} else if grossPay >= MBracket1 && grossPay <= MBracket2 {
			tax = FedRate10*MBracket1 + FedRate12*(grossPay-MBracket1)
			fmt.Printf("%v", tax)
		} else if grossPay <= MBracket2 && grossPay <= MBracket3 {
			tax = FedRate10*MBracket1 + FedRate12*(grossPay-MBracket1) + FedRate22*(grossPay-MBracket2)
			fmt.Printf("%v", tax)
		} else if grossPay <= MBracket3 && grossPay <= MBracket4 {
			tax = FedRate10*MBracket1 + FedRate12*(grossPay-MBracket1) + FedRate22*(grossPay-MBracket2) + FedRate24*(grossPay-MBracket3)
			fmt.Printf("%v", tax)
		} else if grossPay <= MBracket4 && grossPay <= MBracket5 {
			tax = FedRate10*MBracket1 + FedRate12*(grossPay-MBracket1) + FedRate22*(grossPay-MBracket2) + FedRate24*(grossPay-MBracket3) + FedRate32*(grossPay-MBracket4)
			fmt.Printf("%v", tax)
		} else if grossPay <= MBracket5 && grossPay <= MBracket6 {
			tax = FedRate10*MBracket1 + FedRate12*(grossPay-MBracket1) + FedRate22*(grossPay-MBracket2) + FedRate24*(grossPay-MBracket3) + FedRate32*(grossPay-MBracket4) + FedRate35*(grossPay-MBracket5)
			fmt.Printf("%v", tax)
		} else if grossPay >= MBracket6 {
			tax = FedRate10*MBracket1 + FedRate12*(grossPay-MBracket1) + FedRate22*(grossPay-MBracket2) + FedRate24*(grossPay-MBracket3) + FedRate32*(grossPay-MBracket4) + FedRate35*(grossPay-MBracket5) + FedRate37*(grossPay-MBracket6)
			fmt.Printf("%v", tax)
		}
	}
}

func main() {
	askGrossPay()
	askMaritalStatus()
	computeTax()
}
