package main

import (
	"log"
	"net/http"
)

var test_chain trustedtext_chain_s
const test_pub_key = "faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
const test_pri_key = "366c15a87d86f7a6fe6f7509ecaab3d453f0488b414aef12175a870cc5d1b124faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"






func test_handle(w http.ResponseWriter, r *http.Request) {
	
	
}


func main() {

	test_chain, _ = Genesis(
		test_pub_key,
		[]string{"test", "not-a-blog"},
		test_pri_key,
	)
	test_block_1, _ := Instantiate(
		test_pub_key,
		tt_body{
			Instruction_type: "publish",
			Instruction: "My First ever message",
		},
		test_pri_key,
	)
	test_block_2, _ := Instantiate(
		test_pub_key,
		tt_body{
			Instruction_type: "publish",
			Instruction: "My Second ever message",
		},
		test_pri_key,
	)
	test_chain, _ = Process_incoming_block(test_chain, test_block_1)
	test_chain, _ = Process_incoming_block(test_chain, test_block_2)

	http.HandleFunc("/block", give_block)
	http.HandleFunc("/known_blocks", give_known_blocks)
	http.HandleFunc("/submit_block", submit_block)
	http.HandleFunc("/test", test_handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}