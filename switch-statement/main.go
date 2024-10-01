package main

import "fmt"

func main() {
	favoriteColor := "Blue"

	switch favoriteColor {
	case "Blue":
		fmt.Println()
		fmt.Println("Your Favorite Color Is Blue")
	case "Red":
		fmt.Println()
		fmt.Println("Your Favorite Color Is Red")
	case "Green":
		fmt.Println()
		fmt.Println("Your Favorite Color Is Green")
	default:
		fmt.Println()
		fmt.Println("Your Favorite Color Is Blue, Red, Green")
	}

	// Multi Switch Statement
	day := "minggu"

	switch day {
	case "senin", "selasa", "rabu", "kamis", "jumat":
		fmt.Println()
		fmt.Println("Weekday")
	case "sabtu", "minggu":
		fmt.Println()
		fmt.Println("Weekend")
	default:
		fmt.Println()
		fmt.Println("Public Holiday")
	}

}
