package main

//Format Package
import (
	"fmt"
	"math/rand"
)

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

func getRandomNumber(min int, max int) int {
	//rand generate a number in the range of our min and max
	randomNumber := rand.Intn(max-min+1) + min
	return randomNumber
}

// Try to revise and understand the code, what this funciton are doing
func getSpin(reel []string, rows int, cols int) [][]string {
	result := [][]string{}

	// we are adding and the empty rows to the result
	//for we have those and manipulate at the next step
	for i := 0; i < rows; i++ {
		//We are gonna to append an empty slice
		result = append(result, []string{})
	}

	for col := 0; col < cols; col++ {
		//This is a map because go does not have a set but treatly as a set
		//This gonna get randomly position that we get from reel
		selected := map[int]bool{}
		for row := 0; row < rows; row++ {
			//Here we will gonna randonly pick a position
			for true {
				randomIndex := getRandomNumber(0, len(reel)-1)
				// The first value "_," Will get the value of the selected[random index]
				// and the exist will check if that value exist
				_, exists := selected[randomIndex]

				if !exists {
					selected[randomIndex] = true
					result[row] = append(result[row], reel[randomIndex])
					break
				}
			}
		}
	}
	return result
}

func printSpin(spin [][]string) {
	for _, row := range spin {
		for j, symbol := range row {
			fmt.Printf(symbol)
			if j != len(row)-1 {
				fmt.Print("|")
			}
		}
		fmt.Println()
	}
}

func generatesSymbolArray(symbols map[string]uint) []string {
	symbolArray := []string{}
	// At the end of this we will have a array with all symbols and frequency that we defined
	//  key    ,value
	for symbol, count := range symbols {
		for i := 0; i < int(count); i++ {
			symbolArray = append(symbolArray, symbol)
		}
	}
	return symbolArray
}

func checkWin(spin [][]string, multipliers map[string]uint) []uint {
	lines := []uint{}

	for _, row := range spin {
		win := true
		checkSymbol := row[0]
		for _, symbol := range row[1:] {
			if checkSymbol != symbol {
				win = false
				break
			}
		}
		if win {
			lines = append(lines, multipliers[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}
	return lines
}

// Entry point, this will run automatically when you run
func main() {

	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}

	multipliers := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}

	symbolArr := generatesSymbolArray(symbols)
	balance := uint(200)
	getName()

	for balance > 0 {
		bet := getBet(uint(balance))
		if bet == 0 {
			break
		}

		balance -= bet
		spin := getSpin(symbolArr, 3, 3)
		printSpin(spin)

		winningLines := checkWin(spin, multipliers)
		for i, multi := range winningLines {
			win := multi * bet
			balance += win

			if multi > 0 {
				fmt.Printf("Won $%d, (%d) on Liner  #%d\n", win, multi, i+1)
			}
		}

		fmt.Println(winningLines)
	}

}
