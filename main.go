package main

import "fmt"



func main() {
	test_tt := Instantiate(
		"Johnny Bravo",
		[]string{"diary", "test"},
		"This is the best and worst day ever, and I'm glad I can record it in a trusted way",
	)
	fmt.Println(test_tt)
	test_tt = Edit(test_tt, "Johnny Bravo", "Actually, now I've recorded it, it's simply the best day")
	fmt.Println(test_tt)

}