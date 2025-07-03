/*Inventory Products Tracker CLI Application*/

package main

import (
	"Inventory_Tracker_CLI/internal/inventory"
	"fmt"
)

func main() {
	defer fmt.Println("Thankyou for using Inventory Products Tracker CLI Application")
	fmt.Println("Welcome to Inventory Products Tracker CLI Application")
	var invList []inventory.InventoryStructure // Slice to hold the inventory products

	inventory.DisplayMenu()
	for {
		var choices int

		fmt.Print("Enter your choice : ")
		fmt.Scanln(&choices)
		switch choices {
		case 1:
			inventory.AddingProducts(&invList) // Function to add products to the inventory
		case 2:
			inventory.ListProducts(&invList) // Function to list all products in the inventory
		case 3:
			inventory.UpdateProductQuantity(&invList) // Function to update the quantity of a product
		case 4:
			inventory.DeleteProduct(&invList)
		case 5:
			inventory.TotalInventoryValue(&invList) // Function to calculate the total value of the inventory
		case 6:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}

}
