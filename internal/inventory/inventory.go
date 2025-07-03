package inventory

import (
	"fmt"
	"strings"
)

func DisplayMenu() {
	fmt.Println("Welcome to the Inventory Tracker CLI")
	fmt.Println("1. Add Products")
	fmt.Println("2. List Products")
	fmt.Println("3. Update Quantity of the Product")
	fmt.Println("4. Delete Product")
	fmt.Println("5. Calculate Total Value of Products in a Category")
	fmt.Println("5. Exit")
}

func AddingProducts(inventory *[]InventoryStructure) {

	var name string
	var price float64
	var quantity int
	var category string

	fmt.Println("Enter the Product name : ")
	fmt.Scanln(&name)
	fmt.Println("Enter the price of the Product")
	fmt.Scanln(&price)
	fmt.Println("Enter the Quantity of the Product")
	fmt.Scanln(&quantity)
	fmt.Println("Enter the category of the Product")
	fmt.Scanln(&category)

	ProductStorage := InventoryStructure{
		ProductName:     name,
		ProductPrice:    price,
		ProductQuantity: quantity,
		ProductCategory: category,
	}

	*inventory = append(*inventory, ProductStorage)
	fmt.Println("Product added in the Inventory")

}

func ListProducts(inventory *[]InventoryStructure) {
	if len(*inventory) == 0 {
		fmt.Println("No products in the inventory.")
		return
	}

	fmt.Println("Products in the inventory:")
	for i, product := range *inventory {
		fmt.Printf("Product ID: %d, Name: %s, Price: %.2f, Quantity: %d, Category: %s\n",
			i+1, product.ProductName, product.ProductPrice, product.ProductQuantity, product.ProductCategory)
	}
}

func UpdateProductQuantity(inventory *[]InventoryStructure) {
	var Product_Name string
	fmt.Println("Enter the Product Name")
	fmt.Scanln(&Product_Name)

	for i, storages := range *inventory {
		if strings.EqualFold(storages.ProductName, Product_Name) {
			var NewQauntity int
			fmt.Println("Enter the new Quantity for the product")
			fmt.Scanln(&NewQauntity)
			(*inventory)[i].ProductQuantity = NewQauntity // Assigning only new quantity to the structure value

			fmt.Println("Updated the product quantity")
			return
		}

	}
}

func DeleteProduct(inventory *[]InventoryStructure) {

	var name string
	fmt.Println("Enter the Product Name to delete:")
	fmt.Scanln(&name)
	for i, err := range *inventory {
		if strings.EqualFold(err.ProductName, name) {
			*inventory = append(((*inventory)[:i]), ((*inventory)[i+1:])...)
			fmt.Println("Product deleted from the inventory")
			return
		}
	}
}

func TotalInventoryValue(invetory *[]InventoryStructure) float64 {
	total := 0.0
	var category string
	fmt.Println("Enter the category of the product to calculate total value:")
	fmt.Scanln(&category)

	for _, Products := range *invetory {
		if strings.EqualFold(Products.ProductCategory, category) {
			total = Products.ProductPrice * float64(Products.ProductQuantity)
		}

	}
	fmt.Printf("Total value of products in the category '%s': %.2f\n", category, total)
	return total
}
