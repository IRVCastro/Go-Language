package main

//Format Package
import "fmt"

func getName() string {
	//Most Common Way to define a variablke
	name := ""
	fmt.Println("Welcome to De Castro's Cassino...")
	fmt.Printf("Enter your Name:")

	//This is gonna Scan the entire line, or anny input that i enter, and assign value to the name var
	// When i use & i'm making reference to thei pointer, the memory location where we are storing the var
	_, err := fmt.Scanln(&name)

	//nil = not equal null
	if err != nil {
		return ""
	}

	fmt.Printf("Welcome %s, let's play \n", name)
	return name
}

func getBet(balance uint) uint {
	var bet uint

	for true {
		fmt.Printf("Enter your bet, or 0 to quit (balance = $%d):", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println(" Bet cannot be Larger than balance")
		} else {
			break
		}
	}
	return bet
}

//Entry point, this will run automatically when you run
func main() {
	balance := 200
	getBet(uint(balance))
	getName()

}
