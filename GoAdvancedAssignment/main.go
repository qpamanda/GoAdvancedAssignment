package main

import (
	aBST "GoAdvancedAssignment/adminBST"
	aOrder "GoAdvancedAssignment/adminOrder"
	aPizza "GoAdvancedAssignment/adminPizza"
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup

	mu sync.Mutex

	standardPizza = []string{"Hawaiian Pizza", "Cheese Pizza", "Pepperoni Pizza", "Chicken Pizza", "Vegan Pizza"}

	pizzaNo = len(standardPizza)

	orderNo = 1000

	pizzaList = &aPizza.Linkedlist{
		Head: nil,
		Size: 0,
	}

	orderQueue = &aOrder.Queue{
		Front: nil,
		Back:  nil,
		Size:  0,
	}

	salesBST = &aBST.BST{
		Root: nil,
	}
)

const (
	standardPrice = 10.90 // standard price of all pizza start at $10.90
	maxOrderQty   = 5     // max order is 5 per selected pizza
)

func printMainMenu() string {
	choice := ""

	fmt.Println()
	fmt.Println("==============================================")
	fmt.Println("*    WELCOME TO AWESOME PIZZA (MAIN MENU)    *")
	fmt.Println("==============================================")

	fmt.Println("1. View Menu & Order a Pizza") // Enqueue
	fmt.Println("2. Edit an Order")
	fmt.Println("3. Remove Order in Queue") // Dequeue (FIFO)
	fmt.Println("4. Search an Order in Queue")
	fmt.Println("5. View All Orders in Queue")
	fmt.Println("6. View Pizza Sales of the Day")
	fmt.Println("7. Manage Pizza (Admin Menu)")
	fmt.Println("8. Exit Program")

	fmt.Println()
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	fmt.Println()

	return choice
}

/* Print a divider line to segregate the sections for easy viewing */
func printDividerLine() {
	fmt.Println("------------------------------------------------------------")
}

func confirmExit() bool {

	inputExit := "N"

	fmt.Print("Are you sure you want to exit the program? \nNote: All data will be lost! \n(Enter 'Y' to confirm): ")
	fmt.Scanln(&inputExit)

	if inputExit == "Y" || inputExit == "y" {
		fmt.Print(">>>>>> Exiting program .................... Goodbye!")
		fmt.Println()
		fmt.Println()

		return true
	}
	return false
}

func manageUserSelection(choice string) {

	defer wg.Done()

	mu.Lock()

	switch choice {
	case "1":

		printDividerLine()
		fmt.Println("# VIEW MENU & ORDER A PIZZA #")
		printDividerLine()

		// Display the pizza menu so that it is easier to make an order
		err := pizzaList.PrintPizzaMenu()

		if err != nil {
			fmt.Println(">> Sorry. No pizza on the menu today.")
		} else {
			fmt.Println()
			/* Go to function in main() to enqueue a pizza order */
			enqueueOrder()
		}
		printDividerLine()

	case "2":
		printDividerLine()
		fmt.Println("# EDIT AN ORDER #")
		printDividerLine()

		/* Go to function in main() to search an order */
		editOrder()

		printDividerLine()

	case "3":

		printDividerLine()
		fmt.Println("# REMOVE AN ORDER (DEQUEUE) #")
		printDividerLine()

		/* Go to function in main() to dequeue an order that came
		in first in the queue */
		dequeueOrder()

		printDividerLine()

	case "4":
		printDividerLine()
		fmt.Println("# SEARCH AN ORDER #")
		printDividerLine()

		/* Go to function in main() to search an order */
		searchOrder()

		printDividerLine()

	case "5":
		printDividerLine()
		fmt.Println("# VIEW ALL ORDERS IN QUEUE #")
		printDividerLine()

		/* Go to function in adminOrder package to print all orders */
		orderQueue.PrintAllOrders(pizzaList)

		printDividerLine()

	case "6":
		printDividerLine()
		fmt.Println("# VIEW PIZZA SALES OF THE DAY #")
		printDividerLine()

		/* Go to function in adminBST to print report of the sales of
		pizza for the day (ordered by pizza name using BST) */
		salesBST.InOrder()

		printDividerLine()

	case "7":
		printAdminMenu()

	default:
		fmt.Println(">> You have input an invalid choice. Please try again.")
	}

	mu.Unlock()
}

func main() {

	pizzaList.CreateStartMenu(standardPizza, standardPrice)

	for {
		choice := printMainMenu()

		// Allow user to exit the program if choice is 8
		if choice == "8" {
			// Prompt user to confirm exit
			if confirmExit() {
				break
			}
		} else {

			wg.Add(1)
			go manageUserSelection(choice)
			//go manageUserSelection("2")
			//go manageUserSelection("3")

			wg.Wait()

			fmt.Println()
		}
	}
}
