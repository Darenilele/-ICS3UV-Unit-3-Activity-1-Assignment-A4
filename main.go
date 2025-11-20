/**
 * @author Daren Ileleji
 * @versopn 1.0.0
 * @date 2025-11-20
 * @fileoverview This program will display the average of the following numbers: 56.9, 89.7, 90.2
 */

package main

import "fmt"

func main () {
	// INPUT (Displaying Constants)
	const NUMBERONE float32 = 56.9
	const NUMBERTWO float32 = 89.7
	const NUMBERTHREE float32 = 90.2

	// PROCESSING (Displaying Calculations)
	answer := (NUMBERONE + NUMBERTWO + NUMBERTHREE) / 3

	// OUTPUT (Showcasing Results)
	fmt.Println("The average of ", NUMBERONE, ",", NUMBERTWO, ", and", NUMBERTHREE, " is ", answer)
	fmt.Println("\nDone.")
	
}