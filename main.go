package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	// "html"
	"log"
	"net/http"
)

var test_chain trustedtext_chain_s
const test_pub_key = "faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
const test_pri_key = "366c15a87d86f7a6fe6f7509ecaab3d453f0488b414aef12175a870cc5d1b124faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"


func extract_submitted_block(r *http.Request) (trustedtext_s, error) {
	var post_deposit []byte
	var err error
	post_deposit, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return trustedtext_s{}, err
	}
	resultant_block := &trustedtext_s{}
	err = json.Unmarshal(post_deposit, resultant_block)
	if err != nil {
		return trustedtext_s{}, err
	}
	return *resultant_block, nil
}

func Verify_block_is_valid(_input_block trustedtext_s) (bool, error) {
	rehash_of_body, err := return_hash(_input_block)
	if err != nil {
		return false, err
	} 
	if rehash_of_body != _input_block.Hash {
		return false, errors.New("body content doesn't match body hash")
	}
	signature_is_valid, err := Verify_hex_encoded_values(_input_block.Author, _input_block.Hash, _input_block.Hash_signature)
	if err != nil {
		return false, err
	} 
	if !signature_is_valid {
		return false, errors.New("hash signature not verified")
	}
	return true, nil
}
 

func test_handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	resultant_block, err := extract_submitted_block(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}
	
	hash_already_in_chain := Is_hash_in_chain(test_chain, resultant_block.Hash)

	if hash_already_in_chain {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Hash already in chain")
		return
	}

	hash_is_valid, err := Verify_block_is_valid(resultant_block)
	if !hash_is_valid {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Block cannot be verified, with error: ", err)
		return
	}

	
	text_block, err := json.Marshal(resultant_block)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, string(text_block))
	
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
	http.HandleFunc("/test", test_handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}