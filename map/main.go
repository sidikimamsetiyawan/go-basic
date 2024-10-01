package main

import "fmt"

func main() {
	// Create Map
	var person = map[string]string{
		"name": "Sidik",
		"job":  "Backend Developer",
	}
	// Map[Job:Backend Developer name:Sidik]
	fmt.Println()
	fmt.Println(person)
	fmt.Println(person["job"])

	stock := map[string]int{
		"book": 2,
		"pen":  12,
	}
	// map[book:2 pen:12]

	fmt.Println()
	fmt.Println(stock)
	fmt.Println(stock["book"])

	var product = make(map[string]string)
	product["code"] = "COD001"
	product["merk"] = "Apple"
	product["model"] = "Mackbook Pro"

	quantity := make(map[string]int)
	quantity["cod001"] = 1
	quantity["stock"] = 100
	quantity["price"] = 16000000

	fmt.Println()
	fmt.Println(product)
	fmt.Println(product["merk"])

	// Length of map
	fmt.Println()
	fmt.Println(len(person))
	fmt.Println(len(stock))
	fmt.Println(len(product))

	// Create Empty Map
	var emptyMap1 = make(map[string]string)
	var emptyMap2 map[string]string
	fmt.Println()
	fmt.Println(emptyMap1)
	fmt.Println(emptyMap2)
	fmt.Println(emptyMap1 == nil)
	fmt.Println(emptyMap2 == nil)

	//Add and Update Element
	product["tahun"] = "2021"
	product["model"] = "mackbook air"

	fmt.Println()
	fmt.Println(product)

	// Delete Element
	delete(product, "merk")
	fmt.Println()
	fmt.Println(product)

	// Map as pass by reference - Jika newProduct diubah, maka map product juga ikut berubah
	newProduct := product

	fmt.Println()
	fmt.Println(product)
	fmt.Println(newProduct)

	newProduct["tahun"] = "2024"
	fmt.Println()
	fmt.Println(product)
	fmt.Println(newProduct)

}
