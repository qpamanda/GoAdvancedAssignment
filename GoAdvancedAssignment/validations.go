package main

import (
	aPizza "GoAdvancedAssignment/adminPizza"
	"errors"
	"fmt"
	"strconv"
)

func validatePizzaNo(pizzaNo string) (int, error) {

	retValue, err := strconv.Atoi(pizzaNo)

	//fmt.Println("retValue = ", retValue)
	//fmt.Println("err = ", retValue)

	if err != nil {
		pizzaNo = ""
		fmt.Print(">> Please enter a valid Pizza No: ")
		fmt.Scanln(&pizzaNo)
		fmt.Println()

		if pizzaNo != "" {
			retValue, _ = validatePizzaNo(pizzaNo)
		}
	} else {

		var pizzaOrder aPizza.Pizza
		pizzaOrder, err = pizzaList.SearchPizza(retValue)

		if err == nil {
			retValue = pizzaOrder.PizzaNo
		} else {
			pizzaNo = ""
			retValue, _ = validatePizzaNo(pizzaNo)
		}
	}

	return retValue, nil
}

func validateOrderQuantity(orderQty string) (int, error) {
	retValue, err := strconv.Atoi(orderQty)

	if err != nil {
		orderQty = ""
		fmt.Printf(">> Please enter a valid quantity (max %d): ", maxOrderQty)
		fmt.Scanln(&orderQty)
		fmt.Println()

		retValue, _ = validateOrderQuantity(orderQty)
	} else {
		if retValue == 0 || retValue > maxOrderQty {
			orderQty = ""
			retValue, _ = validateOrderQuantity(orderQty)
		}
	}
	return retValue, nil
}

func validateOrderNo(orderNo string) (int, error) {
	retValue, err := strconv.Atoi(orderNo)

	if err != nil {
		orderNo = ""
		fmt.Print(">> Please enter a valid order no: ")
		fmt.Scanln(&orderNo)
		fmt.Println()

		if orderNo != "" {
			retValue, _ = validateOrderNo(orderNo)
		} else {
			return retValue, errors.New(">> Please enter a valid order no")
		}

	}
	return retValue, nil
}

func validatePizzaPrice(pizzaPrice string) (float64, error) {
	retValue, err := strconv.ParseFloat(pizzaPrice, 64)

	if err != nil {
		pizzaPrice = ""
		fmt.Print(">> Please enter a valid price: ")
		fmt.Scanln(&pizzaPrice)
		fmt.Println()

		retValue, _ = validatePizzaPrice(pizzaPrice)
	} else {
		if retValue == 0 {
			pizzaPrice = ""
			retValue, _ = validatePizzaPrice(pizzaPrice)
		}
	}
	return retValue, nil
}

func checkPizzaInOrder(pizzaNo int) (bool, error) {

	// If there are no orders made means there are no pizza in any order, thus return false
	if orderQueue.IsEmpty() {
		return false, nil
	} else {
		return orderQueue.SearchPizzaInOrder(pizzaNo)
	}
	//return false, nil
}
