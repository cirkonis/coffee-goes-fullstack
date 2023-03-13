package client

import (
	"bytes"
	"coffee-goes-fullstack/pkg/enums"
	"coffee-goes-fullstack/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func beginCoffeeMachineProcess() {
	fmt.Printf("Write action (buy, fill, take, remaining, exit) \n")
	var input string
	fmt.Scan(&input)
	switch true {
	case input == enums.BuyOption:
		chooseCoffee()
	case input == enums.FillOption:
		fillMachine()
	case input == enums.TakeOption:
		takeMoney()
	case input == enums.RemainingOption:
		remainingSupplies()
	case input == enums.ExitOption:
		exitProgram()
	default:
		fmt.Printf("I think you got confused... \n")
		beginCoffeeMachineProcess()
	}
}

func chooseCoffee() {
	fmt.Printf("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu: \n")
	var optionChosen string
	fmt.Scan(&optionChosen)
	switch true {
	case optionChosen == enums.EspressoOption:
		checkIfEnough(espressoWaterPerCup, espressoMilkPerCup, espressoBeansPerCup)
		updateMachine(-1*espressoWaterPerCup, -1*espressoMilkPerCup, -1*espressoBeansPerCup, -1, espressoCost)
	case optionChosen == enums.LatteOption:
		checkIfEnough(latteWaterPerCup, latteMilkPerCup, latteBeansPerCup)
		updateMachine(-1*latteWaterPerCup, -1*latteMilkPerCup, -1*latteBeansPerCup, -1, latteCost)
	case optionChosen == enums.CappuccinoOption:
		checkIfEnough(cappuccinoWaterPerCup, cappuccinoMilkPerCup, cappuccinoBeansPerCup)
		updateMachine(-1*cappuccinoWaterPerCup, -1*cappuccinoMilkPerCup, -1*cappuccinoBeansPerCup, -1, cappuccinoCost)
	case optionChosen == enums.BackOption:
		chooseCoffee()
	case optionChosen == enums.RemainingOption:
		remainingSupplies()
	default:
		fmt.Printf("I think you got confused... \n")
		beginCoffeeMachineProcess()
	}
	beginCoffeeMachineProcess()
}

func fillMachine() {
	var (
		waterToAdd int
		milkToAdd  int
		beansToAdd int
		cupsToAdd  int
	)

	fmt.Printf("Write how many ml of water you want to add: \n")
	fmt.Scan(&waterToAdd)
	fmt.Printf("Write how many ml of milk you want to add: \n")
	fmt.Scan(&milkToAdd)
	fmt.Printf("Write how many grams of coffee beans you want to add: \n")
	fmt.Scan(&beansToAdd)
	fmt.Printf("Write how many disposable cups you want to add: \n")
	fmt.Scan(&cupsToAdd)

	data := models.CoffeeMachine{Water: waterToAdd, Milk: milkToAdd, Beans: beansToAdd, Cups: cupsToAdd}

	// Convert the struct to JSON
	jsonStr, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create a new request with the JSON body
	req, err := http.NewRequest(http.MethodPut, "https://localhost:8080/machine", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the content type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Handle the response
	fmt.Println("Response status code:", resp.StatusCode)

	beginCoffeeMachineProcess()
}
