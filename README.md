# 📦 Go Inventory Tracker CLI

A simple command-line application written in **Golang** to manage product inventory.  
This CLI tool allows users to add, view, update, and delete products, and also calculate the total inventory value.

---

## 🚀 Features

- ✅ Add new products (Name, Price, Quantity, Category)
- 📋 List all products in a readable format
- ✏️ Update product quantity by name
- ❌ Delete products from inventory
- 💰 Calculate total inventory value
- 💻 Clean and interactive CLI interface

---

## 🧱 Project Structure

```
Simple-CLI-app-to-manage-product-inventory-using-Golang/
├── cmd/
│   └── app/
│       └── main.go               # Entry point
├── internal/
│   └── inventory/
│       ├── inventory.go          # Core business logic (Add, Update, Delete)
│       └── model.go              # Product structure definition
├── go.mod                        # Go Modules
└── README.md
```

---

## 🛠️ How to Run

### 1. Clone the Repository

```bash
git clone https://github.com/Sid-sy/Simple-CLI-app-to-manage-product-inventory-using-Golang.git
cd Simple-CLI-app-to-manage-product-inventory-using-Golang
```

### 2. Run the App

```bash
go run ./cmd/app
```

---

## 🧪 Sample Usage

```bash
Welcome to Inventory Products Tracker CLI Application

1. Add Products
2. List Products
3. Update Quantity of the Product
4. Delete Product
5. Calculate Inventory Value
6. Exit

Enter your choice: 1
Enter the Product name : Laptop
Enter the price of the Product: 65000
Enter the Quantity of the Product: 5
Enter the category of the Product: Electronics
Product added in the Inventory.

Enter your choice: 2
Products in the inventory:
Product ID: 1, Name: Laptop, Price: 65000.00, Quantity: 5, Category: Electronics
```

---

## 📈 Future Enhancements

- 🧠 Input validation
- 🗃️ Save/load inventory to/from JSON file
- 🌐 REST API version using Gin or net/http
- ✅ Unit tests and mocks
- 🖥️ Frontend UI with React or Go templates

---

## 🙋‍♂️ Author

**Siddharth Rai**  
Backend Developer | Golang | C | Linux | DSA  
📫 GitHub: [Sid-sy](https://github.com/Sid-sy)

---

## 🪪 License

This project is licensed under the [MIT License](LICENSE)

---

## 🤝 Contributions

Contributions are welcome! To contribute:

```bash
git add README.md
git commit -m "Add clean README"
git push
```
