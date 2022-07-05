package main

import "fmt"



func main() {
	test_ttc, _ := Genesis(
		"Johnny Bravo",
		[]string{"diary", "test"},
	)
	fmt.Println(test_ttc)
	test_ttc, _ = Amend(
		test_ttc,
		"Johnny Bravo",
		"This is the best and worst day ever, and I'm glad I can record it in a trusted way",
	)
	fmt.Println(test_ttc)

	test_ttc, _ = Amend(
		test_ttc, 
		"Johnny Bravo", 
		"Actually, now I've recorded it, it's simply the best day",
	)
	fmt.Println(test_ttc)

	// test_tt := trustedtext_s{
	// 	"Jimmy Neutron",
	// 	[]string{"Science"},
	// 	"Science is the best",
	// 	"",
	// 	"",
	// }
	// fmt.Println(test_tt)
	// fmt.Printf("%x\n", sign_tt(test_tt).hash)

}