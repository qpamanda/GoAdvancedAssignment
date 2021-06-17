package main

import (
	"bufio"
	"fmt"
	"os"
)

func printAdminMenu() {
	choice := ""
	for {
		fmt.Println("===================================")
		fmt.Println("*    MANAGE PIZZA (ADMIN MENU)    *")
		fmt.Println("===================================")
		fmt.Println("1. Add New Pizza")
		fmt.Println("2. Edit Pizza")
		fmt.Println("3. Delete Pizza")
		fmt.Println("4. View All Pizza")
		fmt.Println("5. Go Back to Main Menu")
		fmt.Println()
		fmt.Print("\nEnter your choice: ")
		fmt.Scanln(&choice)
		fmt.Println()

		if choice == "5" {
			break
		} else {
			switch choice {
			case "1":
				printDividerLine()
				fmt.Println("# ADD NEW PIZZA #")
				printDividerLine()
				fmt.Println()

				addPizza()

				printDividerLine()

			case "2":
				printDividerLine()
				fmt.Println("# EDIT PIZZA #")
				printDividerLine()
				fmt.Println()

				editPizza()

				printDividerLine()

			case "3":
				printDividerLine()
				fmt.Println("# DELETE PIZZA #")
				printDividerLine()
				fmt.Println()

				deletePizza()

				printDividerLine()

			case "4":
				printDividerLine()
				fmt.Println("# VIEW ALL PIZZA #")
				printDividerLine()

				err := pizzaList.PrintPizzaMenu()

				if err != nil {
					fmt.Println(">>Sorry. No pizza to view.")
				}

				printDividerLine()

			default:
				fmt.Println(">> You have input an invalid choice. Please try again.")
			}
		}
		fmt.Println()
	}
}

func addPizza() {

	inputPizzaPrice := ""
	inputConfirm := "N"

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Pizza Name: ")
	scanner.Scan()
	//fmt.Println(scanner.Text())

	inputPizzaName := scanner.Text()

	fmt.Println()

	if inputPizzaName != "" {
		fmt.Print("Enter Pizza Price: ")
		fmt.Scanln(&inputPizzaPrice)
		fmt.Println()

		pizzaPrice, _ := validatePizzaPrice(inputPizzaPrice)

		fmt.Printf("Are you sure you want to add %s @ $%.2f? (Enter 'Y' to confirm): ", inputPizzaName, pizzaPrice)
		fmt.Scanln(&inputConfirm)
		fmt.Println()

		if inputConfirm == "Y" || inputConfirm == "y" {
			pizzaNo := generatePizzaNo()
			pizzaList.AddPizza(pizzaNo, inputPizzaName, pizzaPrice)

			fmt.Printf(">> %s @ $%.2f added successfully.\n", inputPizzaName, pizzaPrice)
		} else {
			fmt.Println(">> No pizza added")
		}
	} else {
		fmt.Println(">> No pizza added")
	}
}

func editPizza() {

	inputPizzaNo := ""
	inputPizzaPrice := ""
	inputConfirm := "N"

	fmt.Print("Enter Pizza No: ")
	fmt.Scanln(&inputPizzaNo)
	fmt.Println()

	if inputPizzaNo != "" {

		pizzaNo, err := validatePizzaNo(inputPizzaNo)

		if pizzaNo == 0 || err != nil {
			fmt.Println(">> Invalid pizza no")
		} else {

			bPizzaInOrder, _ := checkPizzaInOrder(pizzaNo)

			if !bPizzaInOrder {

				pizzaOrder, _ := pizzaList.SearchPizza(pizzaNo)

				fmt.Printf("You have selected to update: %s @ $%.2f\n", pizzaOrder.PizzaName, pizzaOrder.PizzaPrice)
				fmt.Println()

				scanner := bufio.NewScanner(os.Stdin)
				fmt.Print("Enter New Pizza Name: ")
				scanner.Scan()

				//fmt.Println(scanner.Text())

				inputPizzaName := scanner.Text()

				fmt.Println()

				if inputPizzaName != "" {
					fmt.Print("Enter New Pizza Price: ")
					fmt.Scanln(&inputPizzaPrice)
					fmt.Println()

					pizzaPrice, _ := validatePizzaPrice(inputPizzaPrice)

					fmt.Printf("Are you sure you want to update to %s @ $%.2f? (Enter 'Y' to confirm): ", inputPizzaName, pizzaPrice)
					fmt.Scanln(&inputConfirm)
					fmt.Println()

					if inputConfirm == "Y" || inputConfirm == "y" {
						pizzaList.EditPizza(pizzaNo, inputPizzaName, pizzaPrice)

						fmt.Printf(">> %s @ $%.2f updated successfully.\n", inputPizzaName, pizzaPrice)
					} else {
						fmt.Println(">> No pizza updated")
					}
				} else {
					fmt.Println(">> No pizza updated")
				}
			} else {
				fmt.Println(">> Orders have been made on the selected pizza. Cannot update.")
			}
		}

	} else {
		fmt.Println(">> No pizza updated")
	}
}

func deletePizza() {

	inputPizzaNo := ""
	inputConfirm := "N"

	fmt.Print("Enter Pizza No: ")
	fmt.Scanln(&inputPizzaNo)
	fmt.Println()

	if inputPizzaNo != "" {

		pizzaNo, err := validatePizzaNo(inputPizzaNo)

		if pizzaNo == 0 || err != nil {
			fmt.Println(">> Invalid pizza no")
		} else {

			bPizzaInOrder, _ := checkPizzaInOrder(pizzaNo)

			if !bPizzaInOrder {

				pizzaOrder, _ := pizzaList.SearchPizza(pizzaNo)

				fmt.Printf("Are you sure you want to delete %s @ $%.2f? (Enter 'Y' to confirm): ", pizzaOrder.PizzaName, pizzaOrder.PizzaPrice)
				fmt.Scanln(&inputConfirm)
				fmt.Println()

				if inputConfirm == "Y" || inputConfirm == "y" {
					pizzaList.DeletePizza(pizzaNo)

					fmt.Printf(">> %s @ $%.2f deleted successfully.\n", pizzaOrder.PizzaName, pizzaOrder.PizzaPrice)
				} else {
					fmt.Println(">> No pizza deleted")
				}

			} else {
				fmt.Println(">> Orders have been made on the selected pizza. Cannot delete.")
			}
		}

	} else {
		fmt.Println(">> No pizza updated")
	}
}

func generatePizzaNo() int {

	// Increment PizzaNo global variable by 1
	pizzaNo = pizzaNo + 1

	return pizzaNo
}
