package main

import (
	aOrder "GoAdvancedAssignment/adminOrder"
	aPizza "GoAdvancedAssignment/adminPizza"
	"fmt"
)

func addOrder(orderSlice []aOrder.OrderItem, pizzaNo int, orderQty int) []aOrder.OrderItem {
	orderItem := aOrder.OrderItem{
		PizzaNo:  pizzaNo,
		OrderQty: orderQty,
	}

	orderSlice = append(orderSlice, orderItem)

	return orderSlice
}

func dequeueOrder() {
	inputConfirm := "N"

	fmt.Print("Remove completed order from queue? \n(Enter 'Y' to confirm): ")
	fmt.Scanln(&inputConfirm)
	fmt.Println()

	if inputConfirm == "Y" || inputConfirm == "y" {
		pizzaOrder, err := orderQueue.Dequeue()

		/* Once order is dequeued, it will be moved into the binary
		search tree (BST) data structure in package adminBST to keep
		track of pizza sales */
		updatePizzaSales(pizzaOrder.OrderSlice)

		if err != nil {
			fmt.Println(">> No order to remove")
		} else {
			if pizzaOrder.OrderNo != 0 {
				printOrder(pizzaOrder.OrderNo, pizzaOrder.OrderSlice, pizzaOrder.TotalCost)
				fmt.Println()
				fmt.Printf(">> Order: %d removed from queue and added to pizza sales.\n", pizzaOrder.OrderNo)
			}
		}
	} else {
		fmt.Println(">> No order to remove")
	}
}

func editOrder() {

	if !orderQueue.IsEmpty() {
		inputOrderNo := ""

		fmt.Println()
		fmt.Print("Enter order no: ")
		fmt.Scanln(&inputOrderNo)
		fmt.Println()

		if inputOrderNo != "" {
			orderNo, err := validateOrderNo(inputOrderNo)

			if err != nil {
				fmt.Println(err)
			} else {

				var pizzaOrder aOrder.Order
				pizzaOrder, err = orderQueue.SearchOrder(orderNo)

				if err != nil {
					fmt.Println(err)
				} else {
					if pizzaOrder.OrderNo != 0 {
						fmt.Println("You have selected the following order to edit:")
						fmt.Println()
						printOrder(orderNo, pizzaOrder.OrderSlice, pizzaOrder.TotalCost)

						fmt.Println()

						err := pizzaList.PrintPizzaMenu()
						if err != nil {
							fmt.Println(">> Sorry. No pizza on the menu today.")
						} else {
							fmt.Println()
							orderSlice := make([]aOrder.OrderItem, 0)
							addMore := "Y"

							for addMore == "Y" || addMore == "y" {
								inputPizzaNo := ""
								inputOrderQty := ""
								fmt.Print("Enter Pizza No \n(Or press 'Enter' to save order & go back to main menu)?: ")
								fmt.Scanln(&inputPizzaNo)
								fmt.Println()

								if inputPizzaNo != "" {
									pizzaNo, err := validatePizzaNo(inputPizzaNo)

									if pizzaNo == 0 || err != nil {
										break
									} else {
										fmt.Printf("Enter Order Quantity (max %d): ", maxOrderQty)
										fmt.Scanln(&inputOrderQty)
										fmt.Println()
									}

									orderQty, _ := validateOrderQuantity(inputOrderQty)

									orderSlice = addOrder(orderSlice, pizzaNo, orderQty)

									addMore = "N"

									printDividerLine()
									fmt.Print("Enter 'Y' to add more pizza\n(Or press 'Enter' to save order & go back to main menu): ")
									fmt.Scanln(&addMore)
									printDividerLine()
									fmt.Println()

									if addMore != "Y" && addMore != "y" {
										break
									}
								} else {
									break
								}
							}

							if len(orderSlice) > 0 {
								//orderNo := generateOrderNo(orderSlice)

								totalCost := getTotalCost(orderSlice)

								orderQueue.UpdateOrder(orderNo, orderSlice, totalCost)

								printDividerLine()
								fmt.Println("* RECEIPT *")
								fmt.Println()

								printOrder(orderNo, orderSlice, totalCost)

								fmt.Println()
								fmt.Printf(">> Order: %d updated successfully. Please make the difference in payment accordingly.\n", orderNo)

							} else {
								fmt.Println(">> No orders updated")
							}
						}

					}
				}

			}
		} else {
			fmt.Println(">> Please enter a valid order no")
		}
	} else {
		fmt.Println(">> No orders in the queue")
	}

}

func enqueueOrder() {

	orderSlice := make([]aOrder.OrderItem, 0)
	addMore := "Y"

	for addMore == "Y" || addMore == "y" {
		inputPizzaNo := ""
		inputOrderQty := ""
		fmt.Print("Enter Pizza No \n(Or press 'Enter' to save order & go back to main menu)?: ")
		fmt.Scanln(&inputPizzaNo)
		fmt.Println()

		if inputPizzaNo != "" {
			pizzaNo, err := validatePizzaNo(inputPizzaNo)

			if pizzaNo == 0 || err != nil {
				break
			} else {
				fmt.Printf("Enter Order Quantity (max %d): ", maxOrderQty)
				fmt.Scanln(&inputOrderQty)
				fmt.Println()
			}

			orderQty, _ := validateOrderQuantity(inputOrderQty)

			orderSlice = addOrder(orderSlice, pizzaNo, orderQty)

			addMore = "N"

			printDividerLine()
			fmt.Print("Enter 'Y' to add more pizza\n(Or press 'Enter' to save order & go back to main menu): ")
			fmt.Scanln(&addMore)
			printDividerLine()

			fmt.Println()

			if addMore != "Y" && addMore != "y" {
				break
			}
		} else {
			break
		}
	}

	if len(orderSlice) > 0 {
		orderNo := generateOrderNo(orderSlice)

		totalCost := getTotalCost(orderSlice)

		orderQueue.Enqueue(orderNo, orderSlice, totalCost)

		printDividerLine()
		fmt.Println("* RECEIPT *")
		fmt.Println()

		printOrder(orderNo, orderSlice, totalCost)

	} else {
		fmt.Println(">> No orders made")
	}
}

func generateOrderNo(orderSlice []aOrder.OrderItem) int {

	// Increment orderNo global variable by 1 if there are OrderItem in the slice
	if len(orderSlice) > 0 {
		orderNo = orderNo + 1
	}

	return orderNo
}

func printOrder(orderNo int, orderSlice []aOrder.OrderItem, totalCost float64) {

	fmt.Println("Order No: ", orderNo)
	fmt.Println()

	pizzaTotal := 0.0

	for _, v := range orderSlice {
		pizzaOrder, _ := pizzaList.SearchPizza(v.PizzaNo)
		pizzaTotal = float64(v.OrderQty) * pizzaOrder.PizzaPrice

		fmt.Printf("%d x %s\t$%.2f\n", v.OrderQty, pizzaOrder.PizzaName, pizzaTotal)
	}

	fmt.Println("\t\t\t--------")
	fmt.Printf("TOTAL PAYMENT\t\t$%.2f\n", totalCost)
	fmt.Println("\t\t\t--------")
}

func getTotalCost(orderSlice []aOrder.OrderItem) float64 {

	orderTotal := 0.0
	pizzaTotal := 0.0

	var pizzaOrder aPizza.Pizza

	for _, v := range orderSlice {
		pizzaOrder, _ = pizzaList.SearchPizza(v.PizzaNo)
		pizzaTotal = float64(v.OrderQty) * pizzaOrder.PizzaPrice
		orderTotal = orderTotal + pizzaTotal
	}

	return orderTotal
}

func searchOrder() {

	if !orderQueue.IsEmpty() {
		inputOrderNo := ""

		fmt.Println()
		fmt.Print("Enter order no: ")
		fmt.Scanln(&inputOrderNo)
		fmt.Println()

		if inputOrderNo != "" {
			orderNo, err := validateOrderNo(inputOrderNo)

			if err != nil {
				fmt.Println(err)
			} else {
				var pizzaOrder aOrder.Order
				pizzaOrder, err = orderQueue.SearchOrder(orderNo)

				if err != nil {
					fmt.Println(err)
				} else {
					if pizzaOrder.OrderNo != 0 {
						printOrder(orderNo, pizzaOrder.OrderSlice, pizzaOrder.TotalCost)
					}
				}
			}
		} else {
			fmt.Println(">> Please enter a valid order no")
		}
	} else {
		fmt.Println(">> No orders in the queue")
	}

}

func updatePizzaSales(orderSlice []aOrder.OrderItem) {

	for _, v := range orderSlice {
		pizzaOrder, _ := pizzaList.SearchPizza(v.PizzaNo)
		binaryNode := salesBST.Search(v.PizzaNo, pizzaOrder.PizzaName)

		if binaryNode == nil {
			salesBST.Insert(v.PizzaNo, pizzaOrder.PizzaName, v.OrderQty, pizzaOrder.PizzaPrice)
		} else {
			salesBST.Update(v.PizzaNo, pizzaOrder.PizzaName, v.OrderQty)
		}
	}
}
