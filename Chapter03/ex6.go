package main

import "fmt"

func main() {
	fmt.Print("Enter a feet value: ");
	var feet float64;
	fmt.Scanf("%f", &feet);

	meters := feet * 0.3048;

	fmt.Println("Meters: ", meters);
}