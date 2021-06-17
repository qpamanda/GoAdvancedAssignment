package adminBST

import (
	"fmt"
)

type PizzaSales struct {
	PizzaNo    int
	PizzaName  string
	OrderQty   int
	PizzaPrice float64
}

type BinaryNode struct {
	Item  PizzaSales  // to store the data item
	Left  *BinaryNode // pointer to point to left node
	Right *BinaryNode // pointer to point to right node
}

type BST struct {
	Root *BinaryNode
}

func (bst *BST) InsertNode(t **BinaryNode, pizzaNo int, pizzaName string, orderQty int, pizzaPrice float64) error {
	if *t == nil {
		newOrder := PizzaSales{
			PizzaNo:    pizzaNo,
			PizzaName:  pizzaName,
			OrderQty:   orderQty,
			PizzaPrice: pizzaPrice,
		}

		newNode := &BinaryNode{
			Item:  newOrder,
			Left:  nil,
			Right: nil,
		}
		*t = newNode
		return nil
	}

	if pizzaName < (*t).Item.PizzaName {
		bst.InsertNode(&((*t).Left), pizzaNo, pizzaName, orderQty, pizzaPrice)
	} else {
		bst.InsertNode(&((*t).Right), pizzaNo, pizzaName, orderQty, pizzaPrice)
	}

	return nil
}

func (bst *BST) Insert(pizzaNo int, pizzaName string, orderQty int, pizzaPrice float64) {
	bst.InsertNode(&bst.Root, pizzaNo, pizzaName, orderQty, pizzaPrice)

}

// Print
func (bst *BST) InOrderTraverse(t *BinaryNode, totalSales *float64) {

	if t != nil {
		bst.InOrderTraverse(t.Left, totalSales)

		totalCost := float64(t.Item.OrderQty) * t.Item.PizzaPrice
		*totalSales = *totalSales + totalCost
		fmt.Printf("%s\t\t%d\t\t$%.2f\n", t.Item.PizzaName, t.Item.OrderQty, totalCost)

		bst.InOrderTraverse(t.Right, totalSales)
	}
}

func (bst *BST) InOrder() {
	if bst.Root != nil {
		fmt.Println("PIZZA NAME\t\tORDER QUANTITY\tTOTAL COST")
		fmt.Println()

		totalSales := 0.0
		bst.InOrderTraverse(bst.Root, &totalSales)

		fmt.Println()
		fmt.Println("------------------------------------------------------------")
		fmt.Printf("TOTAL SALES OF THE DAY:\t\t\t$%.2f\n", totalSales)

	} else {
		fmt.Println(">> No sales today")
	}
}

func (bst *BST) SearchNode(t *BinaryNode, pizzaNo int, pizzaName string) *BinaryNode {
	if t == nil {
		return nil

	} else {
		if t.Item.PizzaNo == pizzaNo {
			return t
		} else {
			if pizzaName < t.Item.PizzaName {
				return bst.SearchNode(t.Left, pizzaNo, pizzaName)
			} else {
				return bst.SearchNode(t.Right, pizzaNo, pizzaName)
			}
		}
	}
}

func (bst *BST) Search(pizzaNo int, pizzaName string) *BinaryNode {
	return bst.SearchNode(bst.Root, pizzaNo, pizzaName)
}

func (bst *BST) UpdateNode(t *BinaryNode, pizzaNo int, pizzaName string, orderQty int) *BinaryNode {
	if t == nil {
		return nil

	} else {
		if t.Item.PizzaNo == pizzaNo {
			t.Item.OrderQty = t.Item.OrderQty + orderQty
			return t
		} else {
			if pizzaName < t.Item.PizzaName {
				return bst.UpdateNode(t.Left, pizzaNo, pizzaName, orderQty)
			} else {
				return bst.UpdateNode(t.Right, pizzaNo, pizzaName, orderQty)
			}
		}
	}
}

func (bst *BST) Update(pizzaNo int, pizzaName string, orderQty int) *BinaryNode {
	return bst.UpdateNode(bst.Root, pizzaNo, pizzaName, orderQty)
}
