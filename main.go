package main

import "fmt"

func main() {
	kelvin := 373.15 // Ponto de ebulição da água em Kelvin

	celsius := kelvin - 273.15 // Fórmula para converter Kelvin para Celsius
	fmt.Printf("O ponto de ebulição da água é: %.2f°C\n", celsius)
}
