package main

import "fmt"

func main() {
	pizza := &veggieMania{}

	pizzaWithCheese := &cheeseTopping{
		pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggieMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
