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

func computeTaxSingle() {
	if maritalStatus == "S" || maritalStatus == "s" {
		if grossPay <= Bracket1 {
			tax = FedRate10 * grossPay
			fmt.Printf("%v", tax)
		} else if grossPay >= Bracket1 && grossPay <= Bracket2 {
			tax = FedRate10*Bracket1 + FedRate12*(grossPay-Bracket1)
			fmt.Printf("%v", tax)
		} else if grossPay <= Bracket3 && grossPay <= Bracket4 {
			tax = FedRate10*Bracket1 + FedRate12*(grossPay-Bracket1) + FedRate22*(grossPay-Bracket2)
			fmt.Printf("%v", tax)
		} else if grossPay <= Bracket4 && grossPay <= Bracket5 {
			tax = FedRate10*Bracket1 + FedRate12*(grossPay-Bracket1) + FedRate22*(grossPay-Bracket2) + FedRate24*(grossPay-Bracket3)
			fmt.Printf("%v", tax)
		} else if grossPay <= Bracket5 && grossPay <= Bracket6 {
			tax = FedRate10*Bracket1 + FedRate12*(grossPay-Bracket1) + FedRate22*(grossPay-Bracket2) + FedRate24*(grossPay-Bracket3)
			fmt.Printf("%v", tax)
		} else if grossPay >= Bracket6 {
			tax = FedRate10*Bracket1 + FedRate12 + (grossPay - Bracket1) + FedRate22*(grossPay-Bracket3) + FedRate24*(grossPay-Bracket4)
			fmt.Printf("%v", tax)
		}
	} else {
		fmt.Println("You are married")
	}
}

func computeTaxMarried() {
	if maritalStatus == "M" || maritalStatus == "m" {
		if grossPay <= MBracket1 {
			tax = FedRate10 * grossPay
			fmt.Printf("%v", tax)
		} else if grossPay >= MBracket1 && grossPay <= MBracket2 {
			tax = FedRate10*MBracket1 + FedRate12*(grossPay-MBracket1)
			fmt.Printf("%v", tax)
		} else if grossPay <= MBracket3 && grossPay <= MBracket4 {
			tax = FedRate10*MBracket1 + FedRate12*(grossPay-MBracket1) + FedRate22*(grossPay-MBracket2)
			fmt.Printf("%v", tax)
		} else if grossPay <= MBracket4 && grossPay <= MBracket5 {
			tax = FedRate10*MBracket1 + FedRate12*(grossPay-MBracket1) + FedRate22*(grossPay-MBracket2) + FedRate24*(grossPay-MBracket3)
			fmt.Printf("%v", tax)
		} else if grossPay <= MBracket5 && grossPay <= MBracket6 {
			tax = FedRate10*MBracket1 + FedRate12*(grossPay-MBracket1) + FedRate22*(grossPay-MBracket2) + FedRate24*(grossPay-MBracket3)
			fmt.Printf("%v", tax)
		} else if grossPay >= MBracket6 {
			tax = FedRate10*MBracket1 + FedRate12 + (grossPay - MBracket1) + FedRate22*(grossPay-MBracket3) + FedRate24*(grossPay-MBracket4)
			fmt.Printf("%v", tax)
		}
	}
}

func main() {
	askGrossPay()
	askMaritalStatus()
	computeTaxSingle()
	computeTaxMarried()
}

/*
public class ComputeTax
{
       double RATE1 = 0.15;      // Tax rates
       double RATE2 = 0.25;
       double RATE3 = 0.31;
       double S1 = 21450.0;      // Tax brackets for single
       double S2 = 51900.0;
       double M1 = 35800.0;      // Tax brackets for married
       double M2 = 86500.0;

      char   status;
      double income = 0;
      double tax = 0;


      if ( status == 'S' )
      {  // ***** single
         if ( income <= S1 )
            tax = RATE1 * income;
         else if ( income <= S2 )
            tax = RATE1 * S1 + RATE2 * (income - S1);
         else
            tax = RATE1 * S1 + RATE2 * (S2 - S1) + RATE3 * (income - S2);
      }
      else
      {  // ***** married
         if ( income <= M1 )
            tax = RATE1 * income;
         else if ( income <= M2 )
            tax = RATE1 * M1 + RATE2 * (income - M1);
         else
            tax = RATE1 * M1 + RATE2 * (M2 - M1) + RATE3 * (income - M2);
      }

      System.out.println("Tax = " + tax);
   }
}
*/
