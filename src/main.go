package main

import (
	"fmt"
	"strconv"

	"github.com/adelylria/Calculo-NIF/calcdni"
)

func main() {
	var entry string
	for {
		fmt.Print("Introduce los numeros de tu DNI: ")
		_, err := fmt.Scan(&entry)

		if err != nil {
			fmt.Println("Error: Debes introducir solo números.")
			continue
		}

		if !calcdni.IsNumber(entry) {
			fmt.Println("Error: Debes introducir solo números.")
			continue
		}

		number, _ := strconv.Atoi(entry)

		fmt.Printf("La letra del dni es: %c",
			calcdni.LetrasListCalculation(number),
		)
		break
	}
}
