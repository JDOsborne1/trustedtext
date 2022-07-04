package main

import "fmt"



func main() {
	test_ttc := Genesis(
		"Johnny Bravo",
		[]string{"diary", "test"},
	)
	fmt.Println(test_ttc)
	test_ttc = Amend(
		test_ttc,
		"Johnny Bravo",
		"This is the best and worst day ever, and I'm glad I can record it in a trusted way",
	)
	fmt.Println(test_ttc)

	test_ttc = Amend(
		test_ttc, 
		"Johnny Bravo", 
		"Actually, now I've recorded it, it's simply the best day",
	)
	fmt.Println(test_ttc)

}